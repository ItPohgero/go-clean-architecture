#!/bin/bash

# Run Application
go run cmd/web/main.go

# Run Worker
go run cmd/worker/main.go

# New Migration
migrate create -ext sql -dir db/migrations create_table_xxx

# Run Migrations
# Windows
_/migrate/up.bat
# Linux or MacOS
_/migrate/up.sh