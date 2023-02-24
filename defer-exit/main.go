package main

import . "fmt"

func sayHello(name string) {
	defer func() {
		Println("Your majesty!")
	}()
	Println("Hello,", name)
}

func calcNum(number int) {
	if number == 3 {
		Println("hello 1")
		defer Println("hello 3")
	}
}

func main() {
	Println("---------------------------")
	defer sayHello("Ali") // akan dijalankan terakhir
	sayHello("Loren")

	Println("---------------------------")
	calcNum(3)
	Println("hello 2")
	Println("---------------------------")

}
