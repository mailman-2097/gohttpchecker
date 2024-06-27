package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "HTTP Header")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}

	fmt.Fprintf(w, "X-Forwarded-For Header (if Set):"+r.Header.Get("X-FORWARDED-FOR"))
}

func main() {
	http.HandleFunc("/", defaultHandler)
	fmt.Println("Starting Server...")
	http.ListenAndServe(":8888", nil)
}
