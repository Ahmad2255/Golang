package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "text/html")
// 	fmt.Fprint(w, "<h1>Welcome to Ahmad Site!</h>")
// 	if r.URL.Path == "/" {

// 	} else if r.URL.Path == ":/contact" || r.URL.Path == "/contact/" {
// 		fmt.Fprint(w, "to get in touch , please send email to <a href=\"mailto:ahmadian225@gmail.com\">ahmadian225@gmail.com</a>")
// 	} else {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Println(w, "<h1>We Could Not Find What You were looking for :( <h1><p>please email us if you keep being directed to an invalid page<p>")

// 	}
// }
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@lenslocked.com\">"+
		"support@lenslocked.com</a>.")
}
