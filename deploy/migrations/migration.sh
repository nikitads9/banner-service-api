#!/bin/bash
export GOOSE_DBSTRING="host=${DB_HOST} port=${DB_PORT} dbname=${DB_NAME} user=${DB_USER} password=${DB_PASSWORD} sslmode=${DB_SSL}"
export GOOSE_MIGRATION_DIR=.
sleep 2 && goose up -v