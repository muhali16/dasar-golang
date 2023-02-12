package main

import "fmt"

type mhs struct {
	nama string
	nim  int
}

type grade struct {
	mathematic int
	science    int
	mhs
}

func main() {
	{
		/*
			Block ini adalah cara penulisan struct dasar
		*/

		var mhs1 = mhs{}
		mhs1.nama = "Muhammad Ali Mustaqim"
		mhs1.nim = 200022
		fmt.Println("1.", "\tNama:", mhs1.nama)
		fmt.Println("\tNIM:", mhs1.nim)

		var mhs2 = mhs{"Laurence NS", 400056}
		fmt.Println("2.", "\tNama:", mhs2.nama)
		fmt.Println("\tNIM:", mhs2.nim)

		var mhs3 = mhs{
			nim:  500023,
			nama: "Raditya Dika",
		}
		fmt.Println("3.", "\tNama:", mhs3.nama)
		fmt.Println("\tNIM:", mhs3.nim)

		var mhs4 = mhs{
			nim:  600024,
			nama: "Raditya Dika",
		}
		fmt.Println("4.", "\tNama:", mhs4.nama)
		fmt.Println("\tNIM:", mhs4.nim)

		// menggunakan pointer untuk menginisiasi struct
		var mhs5 *mhs = &mhs4
		mhs5.nama = "Ali Dongan Harahap"
		fmt.Println("5.", "\tNama:", mhs5.nama)
		fmt.Println("\tNIM:", mhs5.nim)
	}
	fmt.Println("---------------------------------------------------------")
	{
		/*
			Block ini menjabarkan cara penulisan embedded struct
		*/
		var grade1 = grade{}
		grade1.nama = "Muhammad Ali Mustaqim"
		grade1.nim = 202200
		grade1.mathematic = 80
		grade1.science = 60
		fmt.Println("1.", "\tNama:", grade1.mhs.nama) // called from its mhs struct
		fmt.Println("\tNIM:", grade1.mhs.nim)         // called from its mhs struct
		fmt.Println("\tMath:", grade1.mathematic)
		fmt.Println("\tScience:", grade1.science)

		var grade2 = grade{}
		grade2.nama = "Laurence NS"
		grade2.nim = 202223
		grade2.mathematic = 70
		grade2.science = 90
		fmt.Println("2.", "\tNama:", grade2.nama) // called from its grade struct
		fmt.Println("\tNIM:", grade2.nim)         // called from its grade struct
		fmt.Println("\tMath:", grade2.mathematic)
		fmt.Println("\tScience:", grade2.science)
	}
	fmt.Println("---------------------------------------------------------")
	{
		/*
			Cara mengisi dengan sub-struct
		*/
		var mhs1 = mhs{"Raditya Dika", 1013}
		var grade1 = grade{mhs: mhs1, mathematic: 70, science: 90}
		fmt.Println("1.", "\tNama:", grade1.nama) // called from its grade struct
		fmt.Println("\tNIM:", grade1.nim)         // called from its grade struct
		fmt.Println("\tMath:", grade1.mathematic)
		fmt.Println("\tScience:", grade1.science)
	}
	fmt.Println("---------------------------------------------------------")
	{
		/*
			Anonymous Struct, yaitu pembuatan struct tanpa pendelkarasian di awal
		*/

		var mhs1 = mhs{"Laurence NS", 1013}
		var grade = struct {
			mathematic int
			science    int
			mhs
		}{}
		grade.mhs = mhs1
		grade.mathematic = 70
		grade.science = 90

		fmt.Println("1.", "\tNama:", grade.nama) // called from its grade struct
		fmt.Println("\tNIM:", grade.nim)         // called from its grade struct
		fmt.Println("\tMath:", grade.mathematic)
		fmt.Println("\tScience:", grade.science)
	}
}
