package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// Router strict by slash
	router := mux.NewRouter().StrictSlash(true)

	// Default route
	router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello go-rest")
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}