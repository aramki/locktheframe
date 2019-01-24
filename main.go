package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to LocktheFrame</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe("localhost:8000", nil)
}
