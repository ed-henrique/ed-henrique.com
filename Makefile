watch:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."

build:
	@templ generate
	@mkdir -p bin
	@go build -o bin/main main.go

prod:
	@templ generate
	@mkdir -p bin
	@go build -trimpath -ldflags "-s -w" -o bin/main main.go

.PHONY: watch
