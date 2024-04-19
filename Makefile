include .env
BIN_SERVER := "./bin/banners"
BIN_CLIENT := "./bin/client"



.PHONY: env
env:
	set -o allexport && source ./.env && set +o allexport

.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
			go install -v golang.org/x/tools/gopls@latest
			go get -d github.com/ogen-go/ogen
		go mod tidy

.PHONY: generate
generate: generate-server generate-mocks
generate-server:
	go generate ./pkg
generate-mocks:
	go generate ./internal/...


.PHONY: run
run: env docker-compose-up
docker-compose-up:
	docker-compose up -d

.PHONY: stop
stop:
	docker compose stop

.PHONY: down
down:
	docker compose down

.PHONY: docker-build
docker-build: docker-build-banner migrator
docker-build-banner: 
	docker build --no-cache -f ./deploy/banner/Dockerfile . --tag nikitads9/banner-service-api:app
docker-build-migrator: 
	docker build --no-cache -f ./deploy/migrations/Dockerfile  ./deploy/migrations --tag nikitads9/banner-service-api:migrator


.PHONY: build
build: build-server build-client

build-server:
	go build -v -ldflags "-w -s" -o $(BIN_SERVER) ./cmd/server/main.go
build-client:
	go build -v -ldflags "-w -s" -o $(BIN_CLIENT) ./cmd/client/main.go

# Контейнер с БД должен быть запущен
.PHONY: migrate
migrate:
	docker start banner-migrator

.PHONY: run-integration-tests
run-integration-tests: run-test-environment integration-tests down-test-environment

.PHONY: run-test-environment 
run-test-environment: env docker-compose-test
docker-compose-test:	
	docker compose -f ./docker-compose-test.yml up -d

.PHONY: integration-tests
integration-tests:
	go test ./... --tags=integration -tags=integration

.PHONY: down-test-environment
down-test-environment:
	docker compose -f ./docker-compose-test.yml down

.PHONY: coverage
coverage:
	go test ./... --tags=integration  -coverprofile="coverage.out" -covermode=atomic && go tool cover -html=coverage.out