package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to LocktheFrame</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, "Please get in touch with us @ <a href=\"mailto:support@locktheframe.com\">support@locktheframe.com</a>")
}

func main() {
	//Initial Implementation
	//http.HandleFunc("/", handlerFunc)
	//http.ListenAndServe("localhost:8000", nil)

	//Using Gorilla Mux
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	http.ListenAndServe(":8000", router)
}
