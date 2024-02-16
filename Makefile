SRC = `go list -f {{.Dir}} ./... | grep -v /vendor/`

.PHONY: run
run: templ-generate css-build
	@go run cmd/web/main.go

.PHONY: templ-generate
templ-generate:
	@templ generate

.PHONY: css-build
css-build:
	@./tailwindcss -i ./internal/web/view/assets/css/input.css -o ./internal/web/view/assets/css/output.css --minify

.PHONY: css-watch
css-watch:
	@./tailwindcss -i ./internal/web/view/assets/css/input.css -o ./internal/web/view/assets/css/output.css --minify --watch

.PHONY: dev
dev:
	@air -c .air.toml

.PHONY: get-tailwindcss-cli
get-tailwindcss-cli:
	@curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
	@chmod +x tailwindcss-macos-arm64
	@mv tailwindcss-macos-arm64 tailwindcss

.PHONY: build
build: templ-generate css-build
	@go build -o bin/app cmd/main.go

.PHONY: fmt
fmt:
	@echo "==> Formatting source code..."
	@go fmt $(SRC)

.PHONY: lint
lint:
	@echo "==> Running lint check..."
	@golangci-lint --config .golangci.yml run
	@go vet `go list ./...`
	@echo "==> Finished!"

.PHONY: test
test:
	@echo "==> Running tests..."
	@go clean -testcache
	@go test ./internal/... -race --cover -coverprofile=output/coverage.out
