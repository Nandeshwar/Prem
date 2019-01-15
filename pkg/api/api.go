package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/intro", intro)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/build/static/"))))
	router.PathPrefix("/favicon.ico").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/build/favicon.ico")
	})
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/build/index.html")
	})
	return *router
}

func intro(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Coming soon"))
}
