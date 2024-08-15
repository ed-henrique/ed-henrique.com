live/templ:
	templ generate --watch --proxy="http://localhost:8080" --open-browser=false -v

live/server:
	go run github.com/air-verse/air@v1.52.3 \
	--build.cmd "go build -o bin/main" --build.bin "bin/main" --build.delay "100" \
	--build.exclude_dir "bin,node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

live/tailwind:
	npx tailwindcss -i ./public/input.css -o ./public/output.css --minify --watch

live/sync_assets:
	go run github.com/air-verse/air@v1.52.3 \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "bin" \
	--build.include_dir "public" \
	--build.include_ext "js,css"

live: 
	make -j4 live/templ live/server live/tailwind live/sync_assets

build:
	@templ generate
	@npx tailwindcss -i ./public/input.css -o ./public/output.css
	@mkdir -p bin
	@go build -o bin/main main.go

prod:
	@templ generate
	@npx tailwindcss -i ./public/input.css -o ./public/output.css --minify
	@mkdir -p bin
	@go build -trimpath -ldflags "-s -w" -o bin/main main.go

.PHONY: live/templ live/server live/tailwind live/sync_assets live prod
