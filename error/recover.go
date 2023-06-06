package main

import . "fmt"

func main() {
	var peoples = []string{"loren", "budhi", "ali", "adhit", "pariji"}

	for _, people := range peoples {

		func() {
			defer func() {
				if r := recover(); r != nil {
					Println("** Error occured:", r)
				} else {
					Println("Lancar boss!")
				}
			}()
			if people == "ali" {
				Println("nama", people, "tidak aman.")
				panic("ada nama ali")
			}
		}()
		Println("nama", people, "aman.")
	}
}
