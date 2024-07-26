.PHONY: all build run clean

all: build run

build: build-backend build-frontend

build-backend:
	@echo "Building backend..."
	@	

build-frontend:
	@echo "Building frontend..."
	@cd frontend && GOOS=js GOARCH=wasm go build -o ../static/main.wasm
	@cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" static/

run:
	@echo "Running server..."
	@./bin/server

clean:
	@echo "Cleaning up..."
	@rm -rf bin static/main.wasm static/wasm_exec.js