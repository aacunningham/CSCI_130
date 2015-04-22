package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")
	if value[1] == "" {
		value[1] = "Enter some text in the url!"
	}
	fmt.Fprint(w, value[1])
}
