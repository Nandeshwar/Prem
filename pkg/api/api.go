package api

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func intro(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Coming soon"))
}

type Server struct {
	server *http.Server
	wg     sync.WaitGroup
}

func NewServer(httpPort int) *Server {
	router := mux.NewRouter()
	router.HandleFunc("/intro", intro)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/build/static/"))))
	router.PathPrefix("/favicon.ico").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/build/favicon.ico")
	})
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/build/index.html")
	})

	server := &http.Server{
		Addr:           ":" + strconv.Itoa(httpPort),
		Handler:        router,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1000000,
	}
	s := &Server{server: server}
	return s
}

func (s *Server) Run() error {
	s.wg.Add(1)
	var err error
	go func() {
		err = s.server.ListenAndServe()
		if err == http.ErrServerClosed {
			err = nil
		}
		s.wg.Done()
	}()
	s.wg.Wait()
	return err
}

func (s *Server) Close() {
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		// Looks like we timed out on the graceful shutdown. Force close.
		if err := s.server.Close(); err != nil {
			logrus.Info("\nHttpServer : Service stopping : Error=%v\n", err)
		}
	}
	logrus.Info("\nHttpServer : Stopped\n")
}
