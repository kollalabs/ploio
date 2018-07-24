package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "analytics suck")
}

func main() {
	http.HandleFunc("/hi", hello)
	http.ListenAndServe(":8000", nil)
}
