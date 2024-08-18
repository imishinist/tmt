.PHONY: build
build:
	@echo "Building..."
	@go build -o bin/ -ldflags "-s -w" -trimpath .
