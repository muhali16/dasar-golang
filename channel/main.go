package main

import (
	. "fmt"
	"runtime"
)

// contoh channel dalam parameter
func printHello(person chan string) {
	Println("Hello,", <-person)
}

func main() {
	runtime.GOMAXPROCS(2)

	Println("------------------------------------")
	var messages = make(chan string)

	var SayHelloTo = func(person string) {
		var msg = Sprintf("Hello, %s", person)
		messages <- msg
	}

	go SayHelloTo("Ali")
	go SayHelloTo("Loren")
	go SayHelloTo("Budhi")

	var printer1 = <-messages
	Println(printer1)
	var printer2 = <-messages
	Println(printer2)
	var printer3 = <-messages
	Println(printer3)

	Println("------------------------------------")
	var messagess = make(chan string)
	var person = []string{"Aziz", "Jessar", "Krisna"}

	for _, psn := range person {
		go func(who string) {
			var res = Sprintf(who)
			messagess <- res
		}(psn)
	}

	// printing the result
	for i := 0; i < len(person); i++ {
		printHello(messagess)
	}
}
