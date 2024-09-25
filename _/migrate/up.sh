#!/bin/bash
# Make sure the script is executable with: chmod +x _/migrate/up.sh

# Define variables
DB_USER="postgres"
DB_PASSWORD="root"
DB_HOST="localhost"
DB_PORT="5432"
DB_NAME="godesa"
DB_SSL_MODE="disable"
MIGRATION_PATH="db/migrations"

# Build the database URL
DATABASE_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}"

# Display the database URL for debugging purposes
echo "Running migrations using the following database URL:"
echo "$DATABASE_URL"

# Run the migrate command
migrate -database "$DATABASE_URL" -path "$MIGRATION_PATH" up

# Check for errors
if [ $? -ne 0 ]; then
    echo "Migration failed."
else
    echo "Migration completed successfully."
fi
