# simple-go-boilerplate

### Setting up local env
```shell
go mod download
```

### Run docker compose
```shell
docker-compose up -d
```

### Start service with default config
```shell
go run ./cmd run
```

### Start service with customer config file
```shell
APP_CONFIG_DIR=<path-to-config-file> go run ./cmd server
```