package main


import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Ploio. Safe, Reliable, and fast production deployments will come to you.")
}

func main() {
	http.HandleFunc("/hi", hello)
	http.ListenAndServe(":8000", nil)
}
