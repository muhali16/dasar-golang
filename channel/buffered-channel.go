package main

import (
	. "fmt"
	"reflect"
	"time"
)

func main() {
	// runtime.GOMAXPROCS(4)

	// var channell = make(chan int, 4) // terdapat 6 buffered data, yaitu 0 sampai 5

	// go func() {
	// 	for {
	// 		i := <-channell
	// 		Println("receive data", i)
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }()

	// for i := 0; i < 10; i++ {
	// 	channell <- i
	// 	Println("send data", i)
	// }

	// var num int
	// Scanln(&num)

	var messages = make(chan string, 2) // terdapat 4 buffered data, yaitu 0 sampai 3
	var data = []string{"kuda", "kucing", "harimau", "anjing", "kura-kura", "ayam", "bebek"}
	var reflectChannel = reflect.ValueOf(messages).Type()

	go func() {
		var j int = 0
		for {
			time.Sleep(time.Second)
			i := <-messages
			Println("receive data =", i, "to", j)
			j++
		}
	}()

	for i := 0; i < len(data); i++ {
		Println("send data", data[i], "to", i, reflectChannel)
		messages <- data[i]
	}

	// var num string
	// Scanln(&num)
}
