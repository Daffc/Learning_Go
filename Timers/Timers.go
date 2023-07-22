package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	// Blocks execution util the timer1 fire (similar to sleep).
	<-timer1.C
	fmt.Println("Timer 1 fired!")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired!")
	}()

	// Stops timer2 before it fires and, if it is successfully stopped, return true to the 'stop2' variable.
	// Difference between timer and sleet = timers can be stopped.
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped!!!")
	}

	time.Sleep(2 * time.Second)
}
