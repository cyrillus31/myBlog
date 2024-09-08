package main

import (
	"fmt"
	"net/http"
)

func main() {
  mux := NewServer()

  mux.HandleFunc("GET /bozo/{number}", 
    func(w http.ResponseWriter, r *http.Request) {
    number := r.PathValue("number")
    fmt.Fprintf(w, "What a bozo. %s", number)
  })

  mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1> %s </h1>", "This is the title of the page")
  })


  http.ListenAndServe(":8000", mux)

}
