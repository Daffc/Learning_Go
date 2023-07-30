package main

import (
	"fmt"
	"net/http"
)

// Handing '/hello' route GET request.
func hello(res http.ResponseWriter, req *http.Request) {
	// Printing content to the response buffer.
	fmt.Fprintf(res, "hello\n")
}

// Handing '/headers' route GET request.
func headers(res http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(res, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// Defining routes and caling function to handle the request.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
