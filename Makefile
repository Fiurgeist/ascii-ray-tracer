.PHONY: install
install:
	go get ./...
	go mod tidy

.PHONY: run
run:
	go run -race ./cmd/render/main.go

.PHONY: help
help:
	@echo "Please use 'make <target>' where <target> is one of"
	@echo "  install                 get all dependencies"
	@echo "  run                     run the console ray-tracer"
