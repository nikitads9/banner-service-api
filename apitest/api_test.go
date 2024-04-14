package apitest

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/nikitads9/banner-service-api/internal/app/api"
	"github.com/nikitads9/banner-service-api/internal/middleware/auth"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/nikitads9/banner-service-api/internal/app/repository/banner"
	"github.com/nikitads9/banner-service-api/internal/app/repository/banner/cache"
	"github.com/nikitads9/banner-service-api/internal/app/repository/banner/postgres"
	bannerService "github.com/nikitads9/banner-service-api/internal/app/service/banner"
	"github.com/nikitads9/banner-service-api/internal/app/service/jwt"
	"github.com/nikitads9/banner-service-api/internal/pkg/db"
	"github.com/nikitads9/banner-service-api/internal/pkg/db/pg"
	"github.com/nikitads9/banner-service-api/internal/pkg/observability"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.opentelemetry.io/otel/trace"
)

const (
	TagIDParam           = "tag_id"
	FeatureIDParam       = "feature_id"
	UseLastRevisionParam = "use_last_revision"
	VersionParam         = "version"
	LimitParam           = "limit"
	OffsetParam          = "offset"
	TokenHeader          = "Token"
)

type Database struct {
	Host                 string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port                 string `yaml:"port" env:"DB_PORT" env-default:"5433"`
	Name                 string `yaml:"database" env:"DB_NAME" env-default:"banners_db"`
	User                 string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Password             string `yaml:"password" env:"DB_PASSWORD" env-default:"banners_pass"`
	Ssl                  string `yaml:"ssl" env:"DB_SSL" env-default:"disable"`
	MaxOpenedConnections int32  `yaml:"max_opened_connections" env:"DB_MAX_CONN" env-default:"10"`
}

// JWT конфигурация выдаваемых и проверяемыхтокенов
type JWT struct {
	Secret     string        `yaml:"secret" env:"JWT_SIGNING_KEY" env-default:"verysecretivejwt"`
	Expiration time.Duration `yaml:"expiration" env:"JWT_EXPIRATION" env-default:"2160h"`
}

// Redis конфигурация
type Redis struct {
	Host string `yaml:"host" env:"REDIS_HOST" env-default:"localhost"`
	Port string `yaml:"port" env:"REDIS_PORT" env-default:"5679"`
}

// TestBannerConfig тестовая конфигурация сервиса
type TestBannerConfig struct {
	Database Database
	JWT      JWT
	Redis    Redis
}

// ReadTestBannerConfigFile
func ReadTestBannerConfigFile(path string) (*TestBannerConfig, error) {
	config := &TestBannerConfig{}

	err := cleanenv.ReadConfig(path, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// GetDBConfig ...
func (b *TestBannerConfig) GetDBConfig() (*pgxpool.Config, error) {
	dbDsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s", b.Database.User, b.Database.Name, b.Database.Password, b.Database.Host, b.Database.Port, b.Database.Ssl)

	poolConfig, err := pgxpool.ParseConfig(dbDsn)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	poolConfig.MaxConns = b.Database.MaxOpenedConnections

	return poolConfig, nil
}

type ApiSuite struct {
	suite.Suite
	handler          *api.Implementation
	server           *desc.Server
	pgClient         pg.Client
	txManager        db.TxManager
	rdsClient        *redis.Client
	bannerRepository banner.Repository
	bannerService    *bannerService.Service
	jwtService       jwt.Service
	bannerCache      banner.Cache
	tracer           trace.Tracer
	logger           *slog.Logger
	adminToken       string
	userToken        string
}

func (as *ApiSuite) SetupTest() {
	t := as.Suite.T()
	t.Log("Загрузка конфигурации окружения")
	cfg, err := ReadTestBannerConfigFile("../configs/banners_test_config.yml")
	if err != nil {
		t.Fatalf("error read test config: %v", err)
	}

	ctx := context.Background()

	as.tracer, err = observability.NewMockTracer(ctx, "banners-api-test")
	if err != nil {
		t.Fatalf("[App] Init - could not create mock tracer. Error: %s", err)
	}

	as.logger = slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{Level: slog.LevelError}))

	fmt.Println(cfg.Database)

	t.Log("Проверка работы базы данных окружения")
	cfx, err := cfg.GetDBConfig()
	if err != nil {
		t.Fatalf("[App] Init - cannot get database config. Error: %s", err)
	}

	as.pgClient, err = pg.NewClient(ctx, as.logger, cfx)
	if err != nil {
		t.Fatalf("[App] Init - cannot create connection to database. Error: %s", err)
	}

	as.txManager = pg.NewTransactionManager(as.pgClient.DB())

	if err = as.pgClient.DB().Ping(ctx); err != nil {
		t.Fatalf("[App] Init - cannot check connection to database. Error: %s", err)
	}

	t.Log("Проверка работы хранилища кэша окружения")
	as.rdsClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: "",
		DB:       0,
	})

	if err := as.rdsClient.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("[App] Init - canot check redis connection. Error: %s", err)
	}

	t.Log("Инициализация хранилищ")

	// Repositories
	as.bannerRepository = postgres.NewBannerRepository(as.pgClient, as.tracer, as.logger)
	as.bannerCache = cache.NewBannerCache(as.rdsClient, as.tracer, as.logger)

	t.Log("Инициализация сервисного слоя")

	// Service layer
	as.bannerService = bannerService.NewBannerService(as.bannerRepository, as.bannerCache, as.tracer, as.logger, as.txManager)
	as.jwtService = jwt.NewJWTService(cfg.JWT.Secret, cfg.JWT.Expiration, as.logger)

	t.Log("Инициализация хэндлеров")
	// Handlers
	as.handler = api.NewImplementation(as.bannerService, as.tracer)

	t.Log("Запуск сервера")
	// routes
	as.server, err = desc.NewServer(as.handler, auth.NewSecurity(as.logger, as.jwtService))
	if err != nil {
		t.Fatalf("init router error: %s", err)
	}

	t.Log("Генерация токенов")
	token, err := as.jwtService.GenerateToken(ctx, "user")
	if err != nil {
		t.Fatalf("[App] Init - cannot generate user token. Error: %s", err)
	}
	as.userToken = "UserToken " + token

	token, err = as.jwtService.GenerateToken(ctx, "admin")
	if err != nil {
		t.Fatalf("[App] Init - cannot generate admin token. Error: %s", err)
	}
	as.adminToken = "AdminToken " + token
}

func (as *ApiSuite) TearDownTest() {
	t := as.Suite.T()
	t.Run("truncate", func(t *testing.T) {
		_, err := as.pgClient.DB().ExecContext(context.Background(), db.Query{QueryRaw: `TRUNCATE banners CASCADE`})
		require.Nil(t, err)
	})

	t.Run("flush cache", func(t *testing.T) {
		err := as.rdsClient.FlushAll(context.Background()).Err()
		require.Nil(t, err)
	})

	as.pgClient.Close()
}

func (as *ApiSuite) checkDeleted(bannerID int64) error {
	id := 0
	return as.pgClient.
		DB().
		QueryRowContext(context.Background(), db.Query{QueryRaw: "SELECT banner_id FROM banners_tags WHERE banner_id = $1"}, bannerID).
		Scan(&id)
}

func TestRunApiTest(t *testing.T) {
	suite.Run(t, &ApiSuite{Suite: suite.Suite{}})
}
