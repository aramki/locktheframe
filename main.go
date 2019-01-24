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
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Sorry. We cannot find the page you are looking for. </h1> If you keep getting this error repeatedly, please email us @ <a href=\"mailto:support@locktheframe.com\">support@locktheframe.com</a>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe("localhost:8000", nil)
}
