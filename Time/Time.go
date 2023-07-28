package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Sampling sections of time variable.
	p(then.Year())
	p(then.Month())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	// Time comparison.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Time subtraction (difference)
	diff := now.Sub(then)
	p(diff)

	// Total of time difference in different scales.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Adding a duration interval to a time.
	p(then.Add(diff))
	p(then.Add(-diff))
}
