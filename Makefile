include .env
BIN_SERVER := "./bin/banners"
BIN_CLIENT := "./bin/client"

#GIT_HASH := $(shell git log --format="%h" -n 1)
#LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)


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
generate: generate-server generate mocks
generate-server:
	go generate ./pkg
generate-mocks:
	go generate ./internal/...


.PHONY: docker
docker: env docker-compose
docker-compose:
	docker compose up -d

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

.PHONY: coverage
coverage:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"

PHONY: test-coverage
test-coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out