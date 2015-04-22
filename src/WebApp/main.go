package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/func1", func1)
	http.HandleFunc("/func2", func2)
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the root!")
}

func func1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is func1")
}

func func2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is func2")
}
