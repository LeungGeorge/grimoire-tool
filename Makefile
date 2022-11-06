default: fmt
	@echo "build success..."

fmt:
	go mod tidy
	go mod download
	go fmt ./...
	go build -o bin/

sync:
	go build -o bin/
	git add .
	go commit -m "audo commit by robot"
	git pull
	npm version patch
	git add .
	go commit -m "audo commit by robot"
	git push


publish:
	go mod tidy
	go mod download
	go fmt ./...
	go build -o bin/
	npm version patch
