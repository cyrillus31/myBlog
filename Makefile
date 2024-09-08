build:
	@go build -o bin/myBlog ./cmd/myBlog.go

run: build
	@bin/myBlog
	echo "Server is running..."

