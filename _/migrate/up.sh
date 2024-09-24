#!/bin/bash
# chmod +x _/migrate/up.sh

DB_USER="postgres"
DB_PASSWORD="root"
DB_HOST="localhost"
DB_PORT="5432"
DB_NAME="go-clearn-architecture"
DB_SSL_MODE="disable"
MIGRATION_PATH="db/migrations"

# Build the database URL
DATABASE_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}"

# Run the migrate command
migrate -database "$DATABASE_URL" -path "$MIGRATION_PATH" up
