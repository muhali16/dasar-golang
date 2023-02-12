package main

import "fmt"

func main() {
	var number1 int = 10
	var number2 *int = &number1

	fmt.Println(number1)
	fmt.Println(&number1)

	fmt.Println(number2)
	fmt.Println(*number2)

	fmt.Println("--------------------")
	*number2 = 20

	fmt.Println(number1)
	fmt.Println(&number1)

	fmt.Println(number2)
	fmt.Println(*number2)

	fmt.Println("--------------------")
	number1 = 40

	fmt.Println(number1)
	fmt.Println(&number1)

	fmt.Println(number2)
	fmt.Println(*number2)
}
