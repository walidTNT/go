package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(404)
	w.Write([]byte("<h1>Hello</h1>"))
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response: %s, URL: %s", reflect.TypeOf(w), r.URL.Path)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":3000", nil)
}
