.DEFAULT_GOAL := build
fmt:
	go fmt ./...
.PHONY:fmt
#lint: fmt
#	golint ./pkg... ./cmd/...
#.PHONY:lint
#vet: fmt
#	go vet ./pkg... ./cmd/...
#.PHONY:vet
build: fmt
	go build -o bin/todoApp cmd/main/main.go

run:
	./bin/todoApp
