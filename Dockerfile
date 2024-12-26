FROM golang:1.23-alpine3.21 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o migrate ./cmd/migrator/
RUN go build -o main ./cmd/s3-mini-storage/

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main /app/main
COPY --from=builder /app/migrate /app/migrate
COPY --from=builder /app/internal/adapters/repository/db/migrations /app/internal/adapters/repository/db/migrations

EXPOSE 50051

COPY .env.example .env

COPY docker_entrypoint.sh /app/docker_entrypoint.sh

RUN chmod +x /app/docker_entrypoint.sh

ENTRYPOINT ["/app/docker_entrypoint.sh"]