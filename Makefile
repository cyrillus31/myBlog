build:
	go build -o bin/myBlog ./cmd/myBlog.go

run: build
	bin/myBlog

