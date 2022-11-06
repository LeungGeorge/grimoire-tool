default: fmt
	@echo "build success..."

fmt:
	go mod tidy
	go mod download
	go fmt ./...
	go build -o bin/
	npm version patch

