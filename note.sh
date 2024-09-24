#!/bin/bash

# Run Application
go run cmd/web/main.go

# Run Worker
go run cmd/worker/main.go