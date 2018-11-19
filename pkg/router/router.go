package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RunRouter(port string) {
	log.Printf("http server started at port=" + port)
	router := mux.NewRouter()
	router.HandleFunc("/", intro).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}

func intro(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Coming soon"))
}
