package main

import (
	"fmt"
	"time"
)

func calculate(kind string) {
	for i := 0; i < 10; i++ {
		fmt.Println(kind, i)
		time.Sleep(time.Second)
	}
}

func main() {
	go calculate("a")
	go calculate("b")
	go calculate("c")
	time.Sleep(time.Second * 60)
}
