package main

import "fmt"

func main() {
	/*
		tidak mendukung pengkondisian ternary

		misal:
		var isTrue bool = (isTrue ? "True" : "False")
	*/

	// Contoh pengkondisian if
	var nilai int = 50
	if nilai > 50 {
		fmt.Printf("Nilai anda %d dan anda dinyatakan LULUS\n", nilai)
	} else if nilai > 30 {
		fmt.Printf("Nilai anda %d dan anda dinyatakan LULUS BERSYARAT\n", nilai)
	} else {
		fmt.Printf("Nilai anda %d dan anda dinyatakan TIDAK LULUS\n", nilai)
	}

	// if temporary variable
	point := 300.0

	if avg := point / 4; avg > 80 {
		fmt.Printf("%.3f%s perfect!\n", avg, "%")
	} else if avg > 50 {
		fmt.Printf("%.3f%s nod bad!\n", avg, "%")
	} else {
		fmt.Printf("%.3f%s bad!\n", avg, "%")
	}

	// switch case

	var kelas int = 5
	switch kelas {
	case 1:
		fmt.Println("Anda memilih nomor 1")
	// case dengan banyak konsisi
	case 2, 4, 5:
		fmt.Println("Anda memilih nomor 2 atau 4 atau 5")
	case 3:
		{
			fmt.Println("Anda memilih nomor 3")
			fmt.Println("Pilihan nomor 3")
		}
	default:
		// kurung kurawal digunakan untuk membatasi perintah yang akan dilakukan dalam case jika terdapat lebih dari 2 baris atau statement
		{
			fmt.Println("Pilihan tidak ada")
			fmt.Println("Silakan pilih nomor 1-3")
		}
	}

	// switch case dengan gaya penulisan if else
	index := "c"
	switch {
	case index == "a":
		fmt.Println("Nilai anda antara 80-100")
	case index == "b" || index == "b+":
		fmt.Println("Nilai anda antara 70-79")
	default:
		fmt.Println("Nilai anda dibawah 70")
	}
}
