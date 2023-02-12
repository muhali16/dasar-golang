package main

import (
	"fmt"
)

func main() {
	var data map[string]string
	data = map[string]string{}

	data["nama"] = "Muhammad Ali Mustaqim"
	data["npm"] = "202110715138"

	fmt.Println(data)

	fmt.Println("--------------------------------------")
	data["fakultas"] = "Ilmu Komputer"

	for key, dt := range data {
		fmt.Println(key, ":", dt)
	}

	fmt.Println("--------------------------------------")
	// slice-map yaitu campuran antara slice dan map
	var mhs = []map[string]string{
		{
			"nama":     "Muhammad Ali Mustaqim",
			"npm":      "2002",
			"fakultas": "Informatika",
		},
		{
			"nama":     "Laurence Nicholas Saputra",
			"npm":      "2003",
			"fakultas": "Sistem Informasi",
		},
	}
	for key, m := range mhs {
		fmt.Println(key+1, "\t Nama:", m["nama"])
		fmt.Println("\t NPM:", m["npm"])
		fmt.Println("\t Fakultas:", m["fakultas"])
	}
}
