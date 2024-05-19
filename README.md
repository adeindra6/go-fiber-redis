# Boilerplate for Golang Fiber with Redis Caching

## This repository is to create Golang Fiber backend with Redis for caching
## System Requirements
1. Golang version 1.18 or higher
2. redis-server and redis-cli

## How to run
1. Copy .env.example and rename it to .env
2. Fill .env file with running port and redis server address
3. Run `go mod tidy`
4. Fill redis password on main.go file
5. Run `go get` or maybe you can just run `go run main.go`
