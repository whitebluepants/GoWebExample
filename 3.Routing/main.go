package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("%s", r.URL)
		log.Printf("\n%+v", vars)

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", vars["title"], vars["page"])

	})

	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You've got the all books!")
	})
	bookrouter.HandleFunc("/{title}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You've got the %s", mux.Vars(r)["title"])
	})

	http.ListenAndServe(":80", r)
}
