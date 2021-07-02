package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/people", people)
	r.HandleFunc("/phone", phone)
	http.ListenAndServe(":9090", r)
}

func people(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "this is version AhmadPage ")
}

func phone(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "we will write the phone numbers here ! ")
}
