package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parsing URL and checking for errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)

	// Retrieving user and password from URL.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// Retrieving host and port.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Retrieving path and fragment.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)            // Retrieving query params.
	m, _ := url.ParseQuery(u.RawQuery) // Mapping query params into a map.
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
