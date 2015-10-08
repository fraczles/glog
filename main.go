package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func About(w http.ResponseWriter, r *http.Request) {
	//TODO: write JSONAPI to response
	w.Write([]byte("Hey!\n"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/about", About)

	http.ListenAndServe(":8001", r)
}
