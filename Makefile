.PHONY: install
install:
	go get ./...
	go mod tidy

.PHONY: run
run:
	go run -race ./cmd/render/main.go -width=$(width) -height=$(height) -mode=$(mode) -parallel=$(parallel)

.PHONY: build
build:
	go build -o raytracer cmd/render/main.go

.PHONY: render-low-res
render-low-res:
	go run ./cmd/render/main.go -width=160 -height=90

.PHONY: render-high-res
render-high-res:
	go run ./cmd/render/main.go -width=1280 -height=720

.PHONY: help
help:
	@echo "Please use 'make <target>' where <target> is one of"
	@echo "  install                 get all dependencies"
	@echo "  build                   build executable"
	@echo "  run                     run the console ray-tracer"
	@echo "  render-low-res          run the console ray-tracer in low res mode"
	@echo "  render-high-res         run the console ray-tracer in high res mode"
