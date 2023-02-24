package main

import (
	. "fmt"
	"strconv"
	"strings"
)

func validation(name string) (bool, error) {
	if strings.TrimSpace(name) == "" {
		return false, Errorf("Cannot be empty")
	}

	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		Println("Error occured:", r)
	} else {
		Println("Application running perfectly")
	}
}

func main() {
	defer catch()
	Println("-----------------------------------------")
	var data string
	Print("Input: ")
	Scanln(&data)

	var number int
	var err error
	number, err = strconv.Atoi(data)

	if err == nil {
		Println(number, "is a number")
	} else {
		Println(data, "is not a number")
		Println(err.Error())
	}

	Println("-----------------------------------------")
	var name string
	Print("Nama: ")
	Scanln(&name)

	if valid, err := validation(name); valid {
		Println("Hello,", name)
	} else {
		panic(err.Error())
	}

}
