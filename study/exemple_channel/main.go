package main

import (
	"fmt"
	"time"
)

// chan is the channel, basically the variable that we want to transport from a thread to another
func worker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("worker", workerId, "received", res)
		time.Sleep(time.Second)
	}
}

// make statement makes a channel
func main() {
	channel := make(chan int)

	// More workers can help process tasks concurrently, but only up to the limits of the CPU, I/O, and scheduling overhead.
	go worker(1, channel)
	go worker(2, channel)

	// I send values to my channel, my work go to the channel, get the value and print, after that he will clean the channel and receive more info.
	for i := 0; i < 10; i++ {
		channel <- i
	}
}
