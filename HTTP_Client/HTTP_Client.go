package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// Executing a HTTP get and recovering html into resp.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Printing response status.
	fmt.Println("Response status:", resp.Status)

	// Scanning each line of response body content and printing it.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
