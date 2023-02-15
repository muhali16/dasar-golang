package main

import (
	. "fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

// interface embed dengan interface hitung2d dan hitung3d
type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (q *kubus) keliling() float64 {
	return q.sisi * 4
}

func (q *kubus) luas() float64 {
	return q.sisi * q.sisi
}

func (q *kubus) volume() float64 {
	return math.Pow(q.sisi, 3)
}

func main() {
	Println("Hitung Kubus ------------------")
	// var qubes = &kubus{4.0}
	var bangunRuang hitung = &kubus{4.0}

	Println("Sisi\t:", bangunRuang.(*kubus).sisi, "cm")
	Println("Keliling:", bangunRuang.keliling(), "cm^3")
	Println("Luas\t:", bangunRuang.luas(), "cm^3")
	Println("Volume\t:", bangunRuang.volume(), "cm^3")

}
