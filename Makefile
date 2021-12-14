.PHONY: build
build:
	go build -o ./build/rrogue -v ./cmd/rrogue

.PHONY: buildwa
buildwa:
	GOOS=js GOARCH=wasm go build -o ./build/rrogue.wasm  ./cmd/rrogue

.PHONY: run
run:
	go run ./cmd/rrogue/main.go

.PHONY: test
test:
	go test -v -timeout 10s ./..

.DEFAULT_GOAL := build
