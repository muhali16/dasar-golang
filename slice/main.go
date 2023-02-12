package main

import "fmt"

func main() {
	// pendeklarasian slice
	var animals = []string{"Bird", "Horse", "Cat", "Insect"}

	animals1 := animals[0:2] // will take index 0 - 1
	fmt.Println(animals1)

	animals2 := animals[:3] // will take index 0 - 2
	fmt.Println(animals2)

	animals3 := animals[1:] // will take index 1 - last index
	fmt.Println(animals3)

	animals4 := append(animals1, "Cat")
	fmt.Println(animals4)

	animals5 := make([]string, 3)
	copy(animals5, animals)
	fmt.Println(animals5)

	animals6 := animals[0:2:2]
	fmt.Println(animals6)
}
