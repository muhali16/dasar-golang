package main

import (
	. "fmt"
)

func main() {
	// interface kosong akan menjadi tipedata apapun
	var nama interface{}

	// variabel nama dengan tipedata interface akan menjadi string
	nama = "Muhammad Ali Mustaqim"
	Println(nama)

	// variabel nama dengan tipedata interface akan menjadi int
	nama = 32
	Println(nama)

	// variabel nama dengan tipedata interface akan menjadi slice dengan isi tipedata apapun ada di dalamnya
	var buah = []interface{}{"ali", 202}
	Println(buah...)

	var mhs = map[string]interface{}{
		"nama": "Muhammad Ali Mustaqim",
		"nim":  2004,
	}
	Println(mhs)

	var data = []map[string]interface{}{
		{
			"nama": "Muhammad Ali Mustaqim",
			"nim":  2004,
		},
		{
			"nama": "Laurence Nicholas Saputra",
			"nim":  2008,
		},
	}
	Println(data)

	var fakultas = []interface{}{
		[]map[string]interface{}{
			{
				"nama": "Muhammad Ali Mustaqim",
				"nim":  2004,
			},
			{
				"nama": "Laurence Nicholas Saputra",
				"nim":  2008,
			},
		}, []interface{}{
			"Ilmu Komputer", "Teknik", "Ilmu Ekonomi",
		},
	}
	Println(fakultas...)
}
