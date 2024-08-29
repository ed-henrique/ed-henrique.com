FROM golang:1.22.5-alpine3.20 AS builder

WORKDIR /app

COPY package.json .
COPY package-lock.json .

RUN apk update && \
    apk add --no-cache make npm sqlite && \
    npm i && \
    go install github.com/a-h/templ/cmd/templ@latest && \
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY . .

RUN go generate ./... && make prod

FROM alpine:3.20 AS exec

WORKDIR /app

RUN mkdir -p /app/public/images

COPY --from=builder /app/bin/main .
COPY --from=builder /app/db       ./db
COPY --from=builder /app/public   ./public

EXPOSE 8080

ENTRYPOINT ["/app/main", "-disable-dev"]
