package main

import "fmt"

func main() {
	// Pendeklarasiaan konstanta
	const animal = "Horse"
	fmt.Println("Animal Name:", animal)

	// multi konstanta
	const (
		bird string = "Muray"
		legs        = 2
	)
	fmt.Println("Name:", bird, "Legs:", legs)
}
