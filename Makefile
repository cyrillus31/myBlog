include .env

build:
	@go build -o bin/myBlog ./cmd;

run: build
	@echo "Server is running on http://${APP_HOST}:${APP_PORT}";
	@bin/myBlog;

