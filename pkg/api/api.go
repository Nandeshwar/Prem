package api

import (
	"net/http"

	"strconv"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer(httpPort int) *Server {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         ":" + strconv.Itoa(httpPort),
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	s := &Server{server: server}
	mux.HandleFunc("/", intro)
	return s
}

func (s *Server) Run() error {
	err := s.server.ListenAndServe()
	if err == http.ErrServerClosed {
		err = nil
	}
	return err
}

func (s *Server) Close() {
	s.server.Close()
}

func intro(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Coming soon"))
}
