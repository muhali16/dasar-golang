package main

import "fmt"

func main() {
	// contoh perulangan for ke 1
	fmt.Println("Contoh ke 1")
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke", i+1, ":", i)
	}

	fmt.Println("-------------------------------------")
	// contoh perulangan for ke 2
	fmt.Println("Contoh ke 2")
	j := 0
	for j < 3 {
		fmt.Println("Perulangan ke", j+1, ":", j)
		j++
	}

	fmt.Println("-------------------------------------")
	// contoh perulangan for ke 3
	fmt.Println("Contoh ke 3")
	k := 0
	for {
		fmt.Println("Perulangan ke", k+1, ":", k)
		k++
		if k == 3 {
			break
		}
	}

	fmt.Println("-------------------------------------")
	// contoh perulangan for ke 3
	fmt.Println("Angka ganjil dan genap")
	for l := 1; l <= 10; l++ {
		result := new(string)
		if l%2 == 0 {
			*result = "Genap"
		} else {
			*result = "Ganjil"
		}
		fmt.Println("Angka", l, "adalah", *result)
	}
}
