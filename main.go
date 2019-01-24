package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to LocktheFrame</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "Please get in touch with us @ <a href=\"mailto:support@locktheframe.com\">support@locktheframe.com</a>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe("localhost:8000", nil)
}
