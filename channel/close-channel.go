package main

import (
	. "fmt"
	"runtime"
	"time"
)

func sendMessage(ch chan<- string) {
	Println("Sending data...")
	for i := 0; i < 10; i++ {
		ch <- Sprintf("data %d\n", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	for message := range ch {
		Print(message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	var ch = make(chan string)
	go sendMessage(ch)
	time.Sleep(time.Second)
	printMessage(ch)
}
