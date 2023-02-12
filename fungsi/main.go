package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Random string:", getRandomString(10))

	fmt.Println("--------------------------------")
	var a, c float64 = circle(14.0)
	fmt.Println("Cirlce with 14.00 cm diameter")
	fmt.Printf("Area\t\t: %.1f cm\n", a)
	fmt.Printf("Circumference\t: %.1f cm\n", c)

	fmt.Println("--------------------------------")
	var value = []int{10, 20, 30, 50, 40}
	var avg float64 = getAverage(value...)
	text := fmt.Sprintln("Rata-rata nilai: ", avg)
	fmt.Println(text)
}

// penulisan fungsi dasar
func getRandomString(max int) string {
	var strings string = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	var randString string = ""

	for i := 0; i < max; i++ {
		var randNumber int = rand.Int()%((len(strings)-1)-0+1) + 0
		randString += string(strings[randNumber])
	}

	return randString
}

// penulisan fungsi dengan multiple return
func circle(diameter float64) (area, circumference float64) {
	area = math.Pi * math.Pow((diameter/2), 2)
	circumference = math.Pi * diameter
	return
}

func getAverage(values ...int) float64 {
	total := 0
	for _, value := range values {
		total += value
	}

	return float64(total) / float64(len(values))
}
