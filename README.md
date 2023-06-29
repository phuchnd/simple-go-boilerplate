# simple-go-boilerplate

### Setting up local env
```shell
go mod download
```

### Run docker compose
```shell
docker-compose up -d
```
### Migrate up database
```shell
go run ./cmd migrate -d up
```
### Migrate down database
```shell
go run ./cmd migrate -d down
```
### Start HTTP service with default config
```shell
go run ./cmd run -t http
```
### Start GRPC service with default config
```shell
go run ./cmd run -t grpc
```

### Start service with customize config file
```shell
APP_CONFIG_DIR=<path-to-config-file> go run ./cmd run
```