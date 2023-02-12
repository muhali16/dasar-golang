package main

import "fmt"

func main() {
	var animals = [3]string{"Horse", "Bird", "Cat"} // jika jumlah elemen terdefinisi
	var cats = [...]string{"Tiger", "Lion"}         // jika jumlah elemen tidak terdefinisi
	var nums = make([]int, 3)                       // membuat array menggunakan make dengan tipe data int sebanyak 3 element
	nums[0] = 100
	nums[1] = 200
	nums[2] = 300

	fmt.Println("Jika variabel i digunakan:")
	for i, animal := range animals {
		fmt.Println("Index ke", i, "adalah", animal)
	}

	fmt.Println("-------------------------------------")
	fmt.Println("Jika variabel i tidak digunakan:")
	// menggunakan variabel underscore, agar variabel i bisa tidak digunakan
	for _, cat := range cats {
		fmt.Println("Cat:", cat)
	}

	fmt.Println("-------------------------------------")
	fmt.Println("For menggunakan fungsi len():")
	// menggunakan variabel underscore, agar variabel i bisa tidak digunakan
	for i := 0; i < len(nums); i++ {
		fmt.Println("Index", i, ":", nums[i])
	}
}
