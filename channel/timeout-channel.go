package main

import (
	. "fmt"
	"math/rand"
	"runtime"
	"time"
)

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func recieveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			Println("receive data", data)
		case <-time.After(time.Second * 5):
			Println("Timeout, no activities at 5 sec...")
			break loop
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	var ch = make(chan int)
	go sendData(ch)
	recieveData(ch)
	close(ch)
}
