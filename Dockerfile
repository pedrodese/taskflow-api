# Etapa 1: Construção da aplicação
FROM golang:1.24-alpine AS builder


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go build -o app ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/app /app

EXPOSE 8080

CMD ["/app"]
