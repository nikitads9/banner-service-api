package config

import (
	"fmt"
	"time"

	"github.com/exaring/otelpgx"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// BannerServer конфигурация сервера
type BannerServer struct {
	Host        string        `yaml:"host" env:"BANNERS_HOST" env-default:"0.0.0.0"`
	Port        string        `yaml:"port" env:"BANNERS_PORT" env-default:"3000"`
	Timeout     time.Duration `yaml:"timeout" env:"BANNERS_TIMEOUT" env-default:"6s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"BANNERS_IDLE_TIMEOUT" env-default:"30s"`
}

// Database конфигурация базы данных
type Database struct {
	Host                 string `yaml:"host" env:"DB_HOST" env-default:"db"`
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
	User     string `yaml:"user" env:"REDIS_USER" env-default:"banners"`
	Password string `yaml:"password" env:"REDIS_PASS" env-default:"banners_pass"`
	Host     string `yaml:"host" env:"REDIS_HOST" env-default:"redis"`
	Port     string `yaml:"port" env:"REDIS_PORT" env-default:"5679"`
}

// Tracer конфигурация коллектора телеметрии
type Tracer struct {
	EndpointURL  string  `yaml:"endpoint_url" env:"TRACER_URL" env-default:"http://otelcol:4318"`
	SamplingRate float64 `yaml:"sampling_rate" env:"TRACER_SAMPLING_RATE" env-default:"1.0"`
}

// BannerConfig общий конфиг
type BannerConfig struct {
	Env      string       `yaml:"env" env:"ENV" env-default:"dev"`
	Server   BannerServer `yaml:"server"`
	Database Database     `yaml:"database"`
	Jwt      JWT          `yaml:"jwt"`
	Redis    Redis        `yaml:"redis"`
	Tracer   Tracer       `yaml:"tracer"`
}

// ReadBannerConfigFile ...
func ReadBannerConfigFile(path string) (*BannerConfig, error) {
	config := &BannerConfig{}

	err := cleanenv.ReadConfig(path, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// ReadBannerConfigEnv ...
func ReadBannerConfigEnv() (*BannerConfig, error) {
	config := &BannerConfig{}

	err := cleanenv.ReadEnv(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// GetServerConfig ...
func (b *BannerConfig) GetServerConfig() *BannerServer {
	return &b.Server
}

// GetJWTConfig ...
func (b *BannerConfig) GetJWTConfig() *JWT {
	return &b.Jwt
}

// GetEnv ...
func (b *BannerConfig) GetEnv() string {
	return b.Env
}

// GetTracerConfig ...
func (b *BannerConfig) GetTracerConfig() *Tracer {
	return &b.Tracer
}

// GetRedisConfig ...
func (b *BannerConfig) GetRedisConfig() *Redis {
	return &b.Redis
}

// GetDBConfig ...
func (b *BannerConfig) GetDBConfig() (*pgxpool.Config, error) {
	dbDsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s", b.Database.User, b.Database.Name, b.Database.Password, b.Database.Host, b.Database.Port, b.Database.Ssl)

	poolConfig, err := pgxpool.ParseConfig(dbDsn)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.Tracer = otelpgx.NewTracer(otelpgx.WithTrimSQLInSpanName())
	poolConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	poolConfig.MaxConns = b.Database.MaxOpenedConnections

	return poolConfig, nil
}

// GetAddress ...
func (b *BannerConfig) GetAddress(host string, port string) string {
	address := host + ":" + port
	//TODO: regex check
	return address
}
