#!/bin/bash

# Nama binary yang akan dihasilkan
BINARY_NAME="godesa"

# Direktori di mana file main.go berada
SOURCE_DIR="./cmd/web"

set -e
echo "Rebuild for $BINARY_NAME..."
# List pm2 processes
sudo pm2 list
# Prompt the user for the KEY value
read -p "Enter the KEY for pm2 restart: " KEY

echo "Pull repository"
git pull

# Build aplikasi
echo "Building the application for $GOOS/$GOARCH..."
go build -o $BINARY_NAME $SOURCE_DIR/main.go

# Cek apakah build berhasil
if [ $? -eq 0 ]; then
    echo "Build successful! Binary: $BINARY_NAME"
    # Restart the pm2 process with the provided KEY
    sudo pm2 restart $KEY
    echo "Build and restart process completed successfully."
else
    echo "Build failed!"
    exit 1
fi
