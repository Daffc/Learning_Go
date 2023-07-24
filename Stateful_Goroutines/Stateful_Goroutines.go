package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

// Used by a goroutine to manage data access to ther goroutines.
func orchestrate(reads chan readOp, writes chan writeOp) {

	// Map used to controll data exchange between reader and writers goroutines.
	var state = make(map[int]int)

	for {
		select {
		case read := <-reads:
			read.resp <- state[read.key]
		case write := <-writes:
			state[write.key] = write.val
			write.resp <- true
		}
	}
}

func writer(writes chan writeOp, writeOps *uint64) {
	for {
		write := writeOp{
			key:  rand.Intn(5),
			val:  rand.Intn(100),
			resp: make(chan bool),
		}
		writes <- write // Sends a value (writeOp) to 'writes' channel
		<-write.resp    // Waits orcherstrator goroutine read previsou setted value.
		atomic.AddUint64(writeOps, 1)
		time.Sleep(time.Millisecond)
	}
}

func reader(reads chan readOp, readOps *uint64) {
	for {
		read := readOp{
			key:  rand.Intn(5),
			resp: make(chan int),
		}
		reads <- read // Sends a value (readOp) to 'reads' channel
		<-read.resp   // Waits orcherstrator goroutine read previsou setted value.
		atomic.AddUint64(readOps, 1)
		time.Sleep(time.Millisecond)
	}
}
func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go orchestrate(reads, writes)

	// Initiates 100 reader goroutines
	for r := 0; r < 100; r++ {
		go reader(reads, &readOps)
	}

	// Initiates 10 writer goroutines
	for w := 0; w < 10; w++ {
		go writer(writes, &writeOps)
	}

	// Gives some time to goroutines to execute.
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
