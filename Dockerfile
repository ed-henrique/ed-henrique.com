FROM golang:1.22.5-alpine3.20 AS builder

WORKDIR /app

COPY . .

RUN apk update && \
    apk add --no-cache make npm && \
    npm i && \
    go install github.com/a-h/templ/cmd/templ@latest

RUN make prod

FROM alpine:3.20 AS exec

WORKDIR /app

RUN mkdir -p /app/public/images

COPY --from=builder /app/bin/main               .
COPY --from=builder /app/public/init.js         ./public/init.js
COPY --from=builder /app/public/toggle-theme.js ./public/toggle-theme.js
COPY --from=builder /app/public/output.css      ./public/output.css
COPY --from=builder /app/public/images/         ./public/images/

EXPOSE 8080

ENTRYPOINT ["/app/main", "-disable-dev"]
