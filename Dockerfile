FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd


FROM alpine:latest

WORKDIR /app

COPY --from=0 /app ./

COPY ./app-config.yaml ./config/app-config.yaml

ENV APP_CONFIG_DIR=/app/config/

ENTRYPOINT ["/app/app"]

CMD ["run"]
