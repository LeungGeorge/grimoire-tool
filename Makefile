default: fmt
	@echo "build success..."

fmt:
	go mod tidy
	go mod download
	go fmt ./...
	go build -o bin/

sync:
	go build -o bin/
	grimoire-tool git sync
	npm version patch
	grimoire-tool git sync
	npm publish
