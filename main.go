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

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, "Will be available shortly. Please contact @ <a href=\"mailto:support@locktheframe.com\">support@locktheframe.com</a> in case of any urgent queries")
}

func handle404NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprint(w, "Sorry. This is embarassing. If you continue to face this issue, please contact us @ <a href=\"mailto:support@locktheframe.com\">support@locktheframe.com</a>")
}

func main() {
	//Initial Implementation
	//http.HandleFunc("/", handlerFunc)
	//http.ListenAndServe("localhost:8000", nil)

	//Using Gorilla Mux
	var notFoundHandler http.Handler = http.HandlerFunc(handle404NotFound)
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/faq", faq)
	router.NotFoundHandler = notFoundHandler
	http.ListenAndServe(":8000", router)
}
