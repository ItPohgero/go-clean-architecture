@echo off

:: Define variables
set DB_USER=postgres
set DB_PASSWORD=root
set DB_HOST=localhost
set DB_PORT=5432
set DB_NAME=go-clearn-architecture
set DB_SSL_MODE=disable
set MIGRATION_PATH=db/migrations

:: Build the database URL
set DATABASE_URL=postgres://%DB_USER%:%DB_PASSWORD%@%DB_HOST%:%DB_PORT%/%DB_NAME%?sslmode=%DB_SSL_MODE%

:: Run the migrate command
migrate -database "%DATABASE_URL%" -path "%MIGRATION_PATH%" up

pause
