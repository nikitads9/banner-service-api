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

.PHONY: run-tests
run-tests: test-unit run-test-environment test-integration coverage down-test-environment

.PHONY: run-test-environment 
run-test-environment: env docker-compose-test
docker-compose-test:	
	docker compose -f ./docker-compose-test.yml up -d

.PHONY: test-unit
test-unit:
	go test ./... --tags=unit

.PHONY: test-integration
test-integration:
	go test ./... --tags=integration

.PHONY: down-test-environment
down-test-environment:
	docker compose -f ./docker-compose-test.yml down

.PHONY: coverage
coverage: unit-tests integration-tests cover
unit-tests:
	export PWD=$(pwd)
	go test -cover ./... -tags=unit -args -test.gocoverdir="${PWD}/coverage/unit/"
integration-tests:
	go test -cover ./... -tags=integration -args -test.gocoverdir="${PWD}/coverage/integration/"
cover:
	go tool covdata percent -i=./coverage/unit,./coverage/integration
	go tool covdata textfmt -i=./coverage/unit,./coverage/integration -o coverage/profile
	go tool cover -func=coverage/profile
	go tool cover -html=coverage/profile
	rm ./coverage/unit/* && rm ./coverage/integration/*
