include .env
BIN_BANNERS := "./bin/bookings"

#GIT_HASH := $(shell git log --format="%h" -n 1)
#LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

env:
	set -o allexport && source ./.env && set +o allexport