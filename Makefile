.PHONY: test fmt vet run build docker coverage

APP=go-simple-ci-cd
PKG=./...

test:
	go test $(PKG) -race -count=1

coverage:
	go test $(PKG) -coverprofile=coverage.out
	go tool cover -func=coverage.out | tail -n 1

fmt:
	gofmt -w .

vet:
	go vet $(PKG)

run:
	go run ./cmd/server

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-s -w" -o bin/$(APP) ./cmd/server

docker:
	docker build -t $(APP):local .
