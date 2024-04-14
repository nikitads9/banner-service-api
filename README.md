# banner-service-api

[![build and test](https://github.com/nikitads9/banner-service-api/actions/workflows/build.yml/badge.svg)](https://github.com/nikitads9/banner-service-api/actions/workflows/build.yml)
[![linters](https://github.com/nikitads9/banner-service-api/actions/workflows/linter.yml/badge.svg)](https://github.com/nikitads9/banner-service-api/actions/workflows/linter.yml)
[![test coverage](https://codecov.io/gh/nikitads9/banner-service-api/graph/badge.svg?token=FGLI9UL7CS)](https://codecov.io/gh/nikitads9/banner-service-api) <br />
Avito backend task 2024.

## Сервис баннеров

В Авито есть большое количество неоднородного контента, для которого необходимо иметь единую систему управления. В 
частности, необходимо показывать разный контент пользователям в зависимости от их принадлежности к какой-либо группе. 
Данный контент мы будем предоставлять с помощью баннеров.

## Описание задачи

Необходимо реализовать сервис, который позволяет показывать пользователям баннеры, в зависимости от требуемой фичи и 
тега пользователя, а также управлять баннерами и связанными с ними тегами и фичами.

## Как решалась задача

- Задача выполнена на языке `Go` версии 1.22.1. Проект построен по принципу `model`, `view`, `controller` .
- Для генерирования валидаторов, прототипов хэндлеров был использован генератор [ogen-go](https://github.com/ogen-go/ogen), который создает основные схемы для взаимодействия с клиентом, генерирует сервер, базовые middleware для сбора метрик и трейсов в формате [open telemetry](https://github.com/open-telemetry). Протокол OTLP позволяет централизованно собирать телеметрию, обрабатывать под нужный формат и передавать получателям, таким как LogStash, Fluentd, Prometheus, Jaeger, Elasticsearch. Я выбрал `ogen-go`, поскольку привык к удобной генерации кода для создания бэкенд-интерфейсов на опыте с связкой `grpc` + `protobuf`, где есть удобная утилита `protogen`.
- Логирование было реализовано с помощью slog, формат логов выбирается в зависимости от обозначенной в конфиге или в переменной окружения среды запуска сервера: local, dev, prod. В двух последних случаях логи выводятся в stdout в формате JSON.
- Написаны middleware для логирования и авторизации по JWT-токену.
- Операции `CreateBanner`, `SetBanner` производятся в транзакции с уровнем изоляции ReadCommitted и откате при какой-либо ошибке.
- В качестве основной среды для запуска решения предполагается `Docker` контейнеры. Описана конфигурация docker-compose для контейнеров с резидентной кэш базой данных `Redis`, реляционной СУБД `PostgreSQL`, коллектора телеметрии `otelcol`, платформы для трейсинга `Jaeger`, NoSQL time series DB `Prometheus`. Для сборки образа серверного приложения написан multi-stage `Dockerfile`, в котором сборка производится на основе образа `golang:1.22-alpine`, а запускаться сервер будет в `alpine:latest`.
- Драйвер для подключения и запросов в PostgreSQL: [pgx](https://github.com/jackc/pgx) и [pgxpool](https://github.com/jackc/pgx/tree/master/pgxpool).
- Для хранения кэша используется `Redis` с SDK [go-redis](https://github.com/redis/go-redis).
- Для удобного развертывания были написаны скрипты `Makefile`, позволяющие накатывать и откатывать миграции, собирать бинарники, производить тестирование и генерировать моки (`mockgen`) и серверный код (`ogen-go`).
- Линтер-проверки, проверка на прохождение тестов и проверка на build выполнены в `Github Actions` с помощью пайплайнов build и linter.
<br /> <br /> <br />
  ![project design](assets/project-design-dark.png#gh-dark-mode-only)
  ![project design](assets/project-design-light.png#gh-light-mode-only)

## Пункты задания

1. ✔ Используйте этот [API](https://github.com/avito-tech/backend-trainee-assignment-2024/blob/main/api.yaml).
   * Данная документация была использована при генерации `ogen-go`. В нее были внесены изменения (для того, чтобы сгенерированный сервер функционировал так как нужно + некоторые изменения были нужны, чтобы удовлетворить требования стандарта Open API 3.0). Окончательная версия документации лежит в директории `docs` и проверена в Insomnia (синтаксических ошибок open api не выявлено). <br /> 

      Изменения: <br /> 
   - Во всех запросах и ответах, содержащих поле content, тип object заменен на free-form json `{}`, чтобы ogen сгенерировал структуру с полем jx.Raw, которая является аналогом json.RawMessage.
   - В запросах и некоторых ответах обозначил required поля, чтобы `ogen-go` генерировал ненуллабельные поля в структурах.
   - Добавил формат ***int64*** для того, чтобы API не использовал неопределенный (зависящий от архитектуры процессора клиента) тип ***int***.
   - Добавил поля, необходимые для удовлетворения требований к синтаксису OAS. Это такие поля как: описание операций, теги операций, глобальные теги, контактные данные.
   - Добавил securitySchema конфигурацию, для того, чтобы четко сформулировать механизм получения токенов.
2. ✔ Тегов и фичей небольшое количество (до 1000), RPS — 1k, SLI времени ответа — 50 мс, SLI успешности ответа — 99.99%
   * В качестве инфраструктуры observability использована связка Prometheus + Jaeger. Первый служит хранилищщем метрик, второй нужен для трейсинга запросов и мониторинга (к сожалению не в риалтайм, приходится прожимать F5). За выполнением SLI можно следить в Jaeger.
3. ✔ Для авторизации доступов должны использоваться 2 вида токенов: пользовательский и админский.
   Получение баннера может происходить с помощью пользовательского или админского токена, а все остальные 
   действия могут выполняться только с помощью админского токена. 
   * Написан middleware для проверки токенов, токены админов начинаются с префикса AdminToken, а у юзеров - UserToken.
4. ✘ Реализуйте интеграционный или E2E-тест на сценарий получения баннера.
   * В процессе.
5. ✔ Если при получении баннера передан флаг use_last_revision, необходимо отдавать самую актуальную информацию.
   В ином случае допускается передача информации, которая была актуальна 5 минут назад.
   * Для пользовательского метода получения содержимого баннера реализовано кэширование по классической схеме. При наличии флага ***use_last_revision*** в запросе, кэш не используется (однако полученное значение все равно запишется в кэш), запрос отправляется в базу данных. Если этого флага нет или он обозначен как ***false***, то сначала происходит попытка получить значение из кэша.
6. ✔ Баннеры могут быть временно выключены. Если баннер выключен, то обычные пользователи не должны его получать, 
   при этом админы должны иметь к нему доступ.
    * При попытке получения баннера пользователем SELECT запрос в БД проверяет баннер на значение флага ***is_active***.

## Пункты дополнительного задания

1. ✔ Адаптировать систему для значительного увеличения количества тегов и фичей, при котором допускается 
   увеличение времени исполнения по редко запрашиваемым тегам и фичам. 
   * Помимо автоматически назначаемых postgres-ом индексов на первичные ключи, были созданы индексы btree на внешние ключи ***feature_id*** и ***tag_id***.
2. ✘ Провести нагрузочное тестирование полученного решения и приложить результаты тестирования к решению.
   * В процессе
3. ✘ Иногда получается так, что необходимо вернуться к одной из трех предыдущих версий баннера в связи с 
   найденной ошибкой в логике, тексте и т.д.  Измените API таким образом, чтобы можно было просмотреть существующие 
   версии баннера и выбрать подходящую версию.
   * Для данного решения представляется оптимальным использовать триггеры на insert/update с удалением третьей с конца строки выбранной из таблицы версий баннеров. Не успел реализовать, сижу над дипломом :(
4. ✘ Добавить метод удаления баннеров по фиче или тегу, время ответа которого не должно превышать 100 мс, 
   независимо от количества баннеров.  В связи с небольшим временем ответа метода, рекомендуется ознакомиться 
   с механизмом выполнения отложенных действий.
5. ✘ Реализовать интеграционное или E2E-тестирование для остальных сценариев
   * В процессе
6. ✔ Описать конфигурацию линтера
   * Настроен пайплайн golangci-lint в файле [linter.yml](.github/workflows/linter.yml).

## Как запустить

### Исполняемый файл сервиса баннеров

Бинарник поддерживает два флага (не считая `-help` или `-h`).

***Флаги:***
```bash
Usage of ./banners:
  -config string
        path to config file (default "./configs/banners_config.yml")
  -configtype string
        type of configuration: environment variables (env) or env/yaml file (file) (default "file")
```

### Конфигурация

Конфиг серверного приложения находится в директории `configs`, а конфиги остальных серверов (Prometheus, otelcol) в папке `deploy`.
Остальные настройки (Jaeger, PostgreSQL) обеспечиваются через параметры, указанные в файле `docker-compose`.
Конфигурационный файл имеет следующие поля:

```yaml
env: "dev" # Определение текущего окружения, возможные значения local, dev, prod
server: # Конфигурация сервера
  host: "0.0.0.0" # Хост сервера
  port: "3000" # Порт, на котором будет поднят сервер
  timeout: 6s # Таймаут на ответ сервера
  idle_timeout: 30s # Таймаут соединения TCP без событий
database: # Конфигурация БД
  database: "banners_db" # Имя БД
  host: "db" # Имя хоста в докер-сети, на котором поднята база данных. В случае если сервер запускается не в контейнере, указываем localhost
  port: "5433" # Порт, на котором будет развернута база данных
  user: "postgres" # Имя пользователя БД
  password: "banners_pass" # Пароль. В идеале пароль необходимо задавать через секреты.
  ssl: "disable" # Обмен сертификатами и ключами шифрования
  max_opened_connections: 10 # Максимальное количество одновременных соединений с БД, разрешенных драйверу PGX.
jwt: # Конфигурация JWT
  secret: "verysecretivejwt" # Ключ подписи токена. Должен совпадать с тем, который предоставите при проверке.
  expiration: 2160h # TTL генерируемого токена
redis: # Конфигурация Редиса
  user: "" # Временно не используется
  password: "" # Временно не используется
  host: "redis" # Имя хоста в докер-сети, на котором поднят Редис.
  port: "6379" # Порт, на котором поднят Редис.
tracer: # Конфигурация трейсинга
  endpoint_url: "http://otelcol:4318" # URL+ порт, на котором поднят коллектор телеметрии
  sampling_rate: 1.0 # Частота сэмплирования трейсов
```

### Сборка и запуск

Для полноценного запуска всего проекта необходимо переименовать `.env.example` в `.env` и запустить команду:
```
make run
```
Если настройки не были изменены, сервер будет поднят на http://localhost:3000/
<br />
Для остановки контейнеров:
```
make stop
```
Для удаления контейнеров используйте команду ниже. Она не удалит образы и тома (в отличие от `docker system prune -a`), только контейнеры.
```
make down
```

Если необходимо отдельно собрать образ приложения, то сделать это можно командой:
```bash
make docker-build-banner
```
А вот эта команда соберет и образ приложения и образ мигратора, обспечивающего изменение структуры БД через механизм миграций (goose).
```bash
make docker-build
```
### Работа с API

После запуска Используя Insomnia или Postman, импортировать ` docs/api.yml ` и настроить параметры окружения (хост сервера, базовый путь, схему http). Ниже пример окружения для запросов пользователей.
```json
{
	"scheme": "http",
	"base_path": "",
	"host": "localhost:3000",
	"token": "UserToken eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk5ODYyMzUsInNjb3BlIjoidXNlciJ9.vmt-FrTKksPPLAnzvXzj3R7lLcVe06xAEi5s_2NLRVI"
}
```
Для админа окружение аналогичное, отличие лишь в токене.
```json
{
	"scheme": "http",
	"base_path": "",
	"host": "localhost:3000",
	"token": "AdminToken eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk5ODYyMzUsInNjb3BlIjoiYWRtaW4ifQ.cev1h-ivEbwx3UJDYOoWIAid-gSRuPh5RObOkkuOY2g"
}
```
Если по каким-то причинам токены, приведенные выше, устарели, [здесь](https://jwt.io/) можно создать новые с следующим содержанием (токен админа отличается тем, что в поле `scope` указано значение ***admin***):
```json
// HEADER:ALGORITHM & TOKEN TYPE
{
  "alg": "HS256",
  "typ": "JWT"
}
// PAYLOAD:DATA
{
  "exp": 1719986235,
  "scope": "user"
}

// VERIFY SIGNATURE
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  
verysecretivejwt

)

```

### Внесение изменений в код и структуру БД

Подробнее про текущий механизм миграций [goose](https://github.com/pressly/goose).
Если вы хотите изменить структуру БД: добавить новые таблицы или поменять текущие, то необходимо создать новый файл формата **ГГГГММДДЧЧММСС_описание_коммита.sql** в директории ` deploy/migrations `. Если у вас установлен `goose`, то создать файл миграции можно командой 
```bash
goose create <NAME> sql
```
Перед перечислением изменений, которые вы хотите внести, необходимо прописать ` -- +goose Up `, если вы не создавали файл с помощью `goose`. <br />
Также необходимо прописать и откат этих изменений (обратные операции), предварив их ` -- +goose Down `. <br />

Чтобы применить изменения, просто пропишите, хотите ли вы накатить миграцию (up) или откатить миграцию (down) в скрипте ` deploy/migrations/migration.sh ` и запустите ранее созданный контейнер ***migrator***. В логах контейнера вы увидите статус исполнения миграции.
Запустить ранее созданный контейнер ***migrator*** можно как командой `docker start migrator`, так и:
```bash
make migrate
```
Как только ваша миграция успешно накатилась, проверьте, сработает ли откат - вернется ли база данных в изначальный статус.

## Нагрузочное тестирование

Было проведено нагрузочное тестирование с использованием утилиты [wrk](https://github.com/wg/wrk).
Параметры устройства, на котором было проведено нагрузочное тестирование:
- OS: Ubuntu 22.04 (WSL)
- CPU: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
- RAM: Ограничено 5GB DDR4
- SSD: Samsung PM981 Polaris 1TB M.2

## Структура проекта

<details>

   ```
   📦 banner-service-api
   ├─ .env.example
   ├─ .github
   │  └─ workflows
   │     ├─ build.yml
   │     └─ linter.yml
   ├─ .gitignore
   ├─ Makefile
   ├─ README.md
   ├─ assets
   │  ├─ project-design-dark.png
   │  └─ project-design-light.png
   ├─ cmd
   │  ├─ client
   │  │  └─ main.go
   │  └─ server
   │     └─ main.go
   ├─ configs
   │  └─ banners_config.yml
   ├─ deploy
   │  ├─ banner
   │  │  └─ Dockerfile
   │  ├─ database
   │  │  └─ init.sql
   │  ├─ migrations
   │  │  ├─ 20240408140625_add_table.sql
   │  │  ├─ 20240413195657_fill_db.sql
   │  │  ├─ Dockerfile
   │  │  └─ migration.sh
   │  ├─ otelcollector
   │  │  └─ otelcol-config.yml
   │  └─ prometheus
   │     └─ prometheus.yml
   ├─ docker-compose.yml
   ├─ docs
   │  └─ api.yml
   ├─ go.mod
   ├─ go.sum
   ├─ internal
   │  ├─ app
   │  │  ├─ api
   │  │  │  ├─ banner.go
   │  │  │  ├─ create.go
   │  │  │  ├─ delete_banner.go
   │  │  │  ├─ get_banner.go
   │  │  │  ├─ get_banners.go
   │  │  │  └─ set_banner.go
   │  │  ├─ convert
   │  │  │  └─ convert.go
   │  │  ├─ model
   │  │  │  └─ banner.go
   │  │  ├─ repository
   │  │  │  └─ banner
   │  │  │     ├─ cache.go
   │  │  │     ├─ cache
   │  │  │     │  └─ banner.go
   │  │  │     ├─ db.go
   │  │  │     ├─ mocks
   │  │  │     │  ├─ banner_service_cache.go
   │  │  │     │  └─ banner_service_repository.go
   │  │  │     ├─ postgres
   │  │  │     │  ├─ banner.go
   │  │  │     │  ├─ create.go
   │  │  │     │  ├─ delete.go
   │  │  │     │  ├─ get_banner.go
   │  │  │     │  ├─ get_banners.go
   │  │  │     │  └─ set_banner.go
   │  │  │     └─ table
   │  │  │        └─ table.go
   │  │  └─ service
   │  │     ├─ banner
   │  │     │  ├─ banner.go
   │  │     │  ├─ create.go
   │  │     │  ├─ delete_banner.go
   │  │     │  ├─ get_banner.go
   │  │     │  ├─ get_banners.go
   │  │     │  └─ set_banner.go
   │  │     └─ jwt
   │  │        └─ jwt.go
   │  ├─ config
   │  │  └─ banner_config.go
   │  ├─ logger
   │  │  ├─ handlers
   │  │  │  └─ slogdiscard
   │  │  │     └─ slogdiscard.go
   │  │  └─ sl
   │  │     └─ sl.go
   │  ├─ middleware
   │  │  ├─ auth
   │  │  │  └─ auth.go
   │  │  └─ logger
   │  │     └─ logger.go
   │  └─ pkg
   │     ├─ banner
   │     │  ├─ app.go
   │     │  └─ service-provider.go
   │     ├─ db
   │     │  ├─ interface.go
   │     │  ├─ mocks_db
   │     │  │  └─ mock_db.go
   │     │  ├─ mocks_tx
   │     │  │  └─ mock_tx.go
   │     │  ├─ pg
   │     │  │  ├─ client.go
   │     │  │  ├─ pg.go
   │     │  │  └─ transaction.go
   │     │  └─ redis
   │     │     └─ client.go
   │     └─ observability
   │        └─ tracer.go
   └─ pkg
      ├─ banner-api
      │  ├─ oas_cfg_gen.go
      │  ├─ oas_client_gen.go
      │  ├─ oas_handlers_gen.go
      │  ├─ oas_interfaces_gen.go
      │  ├─ oas_json_gen.go
      │  ├─ oas_middleware_gen.go
      │  ├─ oas_parameters_gen.go
      │  ├─ oas_request_decoders_gen.go
      │  ├─ oas_request_encoders_gen.go
      │  ├─ oas_response_decoders_gen.go
      │  ├─ oas_response_encoders_gen.go
      │  ├─ oas_router_gen.go
      │  ├─ oas_schemas_gen.go
      │  ├─ oas_security_gen.go
      │  ├─ oas_server_gen.go
      │  ├─ oas_unimplemented_gen.go
      │  └─ oas_validators_gen.go
      └─ generate.go
   ```

</details>