package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains   ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSufix:  ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repead:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1)) // Replaces all occurrences
	p("Replace:   ", s.Replace("foo", "o", "0", 1))  // Replaces the first 'n' occurrences
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))

}
