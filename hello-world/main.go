package main

import "fmt"

func main() {
	// line comment
	fmt.Println("Hello World!")
	fmt.Println("My", "name", "is", "Ali"+"!")
	/*
		Multiline Comment
	*/
	fmt.Printf("My name is %s! I'am %d years old. \n", "Ali", 20)

	// Penulisan variabel terdapat dua cara,

	// Penulisan dengan tipe data:
	var firstName string = "Muhammad Ali"
	// Penulisan tanpa tipe data:
	lastName := "Mustaqim"
	fmt.Println("My", "name", "is", firstName, lastName+"!")

	// Oneline Variable Declaration
	var firstName2, lastName2 string = "Muhammad Ali", "Mustaqim"
	firstName3, lastName3 := "Muhammad Ali", "Mustaqim"
	fmt.Println(firstName2, lastName2)
	fmt.Println(firstName3, lastName3)

	// Pointer
	name := new(string)
	*name = "Muhammad Ali Mustaqim"

	fmt.Println("My name is: ", name)  // akan muncul alamat memori variabel
	fmt.Println("My name is: ", *name) // akan muncul isi variabel

}
