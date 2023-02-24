package main

import (
	. "fmt"
	"math/rand"
	"time"
)

var letters = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890")

func randStringV1(length int) string {
	var result string
	for i := 0; i < length; i++ {
		result += Sprintf("%c", letters[rand.Intn(len(letters))])
	}
	return result
}

func randStringV2(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] += letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	Println("--------------------------------------------")
	Println("Random number:", rand.Int())
	Println("Random number in certain range 1-6:", rand.Intn(6))

	Println("--------------------------------------------")
	var text = []rune("qwertyuiopasdfghjklzxcvbnm")
	Printf("text in index number 3 is: %c \n", text[3])
	Printf("text in index number 23 is: %c \n", text[23])
	Printf("text in index number 12 is: %c \n", text[12])

	Println("--------------------------------------------")
	Println("Random string v1: ", randStringV1(20))
	Println("Random string v2: ", randStringV2(20))
}
