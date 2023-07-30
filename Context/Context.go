package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(res http.ResponseWriter, req *http.Request) {
	// Recovers request context from req.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	// Awaits 10 secconds, an then returns response to the client...
	case <-time.After(10 * time.Second):
		fmt.Fprintf(res, "hello\n")

	// If the context channel is closed (client cancels the request)...
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(res, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)

	// Listen to the 8090 TCP port.
	http.ListenAndServe(":8090", nil)
}
