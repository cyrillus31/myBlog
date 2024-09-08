package web

import (
	"cyrillus31/myBlog/api"
	"strconv"
	"net/http"
)

type Server struct {
	Port string
	Host string
}

func NewServer(port int, host string) *Server {
  portString := strconv.Itoa(port)
	return &Server{
		Port: portString,
		Host: host,
	}
}

func (s *Server) Run() {
	mux := http.NewServeMux()
  mux.HandleFunc("GET /", api.HandleBozo)

  apiMux := http.NewServeMux()
  apiMux.Handle("/a/", http.StripPrefix("/a", mux))

  // http.ListenAndServe(s.Host + ":" + s.Port, apiMux)
  http.ListenAndServe(s.Host + ":" + s.Port, mux)
}
