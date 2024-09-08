package main

import (
	"cyrillus31/myBlog/web"
	"fmt"
)

func main() {
  apiServer := web.NewServer(8000, "localhost")
  fmt.Print("Server is running on ", apiServer.Host, ":", apiServer.Port)
  apiServer.Run()
}
