package main

import (
	. "fmt"
	"runtime"
	"time"
)

func printText(max int, message string) {
	for i := 1; i <= max; i++ {
		time.Sleep(100 * time.Millisecond)
		Println(i, message)
	}
}

func main() {
	runtime.GOMAXPROCS(4)

	go printText(4, "Hello")
	printText(4, "Ali")

	// var nama1, nama2, nama3 string
	// Scanln(&nama1, &nama2, &nama3)

	// Println("Hi,", nama1, nama2, nama3)
}
