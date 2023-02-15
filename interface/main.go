package main

import (
	"fmt"
	"math"
)

type hitung interface {
	keliling() float64
	luas() float64
}

type persegipanjang struct {
	panjang float64
	lebar   float64
}

type segitigasiku struct {
	alas   float64
	tinggi float64
}

func (p persegipanjang) keliling() float64 {
	var result float64 = 2 * (p.panjang + p.lebar)
	return result
}

func (p persegipanjang) luas() float64 {
	var result float64 = (p.panjang * p.lebar)
	return result
}

func (s segitigasiku) getLuasSegitiga() float64 {
	var result float64 = 0.5 * float64(s.alas*s.tinggi)
	return result
}

func (s segitigasiku) keliling() float64 {
	var miring float64 = float64(math.Sqrt(math.Pow(float64(s.alas), 2) + math.Pow(float64(s.tinggi), 2)))
	var result float64 = miring + s.alas + s.tinggi
	return result
}

func (s segitigasiku) luas() float64 {
	var result float64 = s.getLuasSegitiga() / 2
	return result
}

func main() {
	var bangunDatar hitung

	// Menghitung persegi panjang
	bangunDatar = persegipanjang{2.0, 4.0}
	fmt.Println("Persegi Panjang ---------------")
	fmt.Println("Panjang\t:", bangunDatar.(persegipanjang).panjang, "cm")
	fmt.Println("Lebar\t:", bangunDatar.(persegipanjang).lebar, "cm")
	fmt.Println("Keliling:", bangunDatar.keliling(), "cm")
	fmt.Println("Luas\t:", bangunDatar.luas(), "cm")

	// Menghitung segitiga siku-siku
	bangunDatar = segitigasiku{10.0, 12.0}
	fmt.Println("Segitiga Siku-Siku ---------------")
	fmt.Println("Alas\t:", bangunDatar.(segitigasiku).alas, "cm")
	fmt.Println("Tinggi\t:", bangunDatar.(segitigasiku).tinggi, "cm")
	fmt.Println("LuasAsli:", bangunDatar.(segitigasiku).getLuasSegitiga(), "cm")
	fmt.Println("Keliling:", bangunDatar.keliling(), "cm")
	fmt.Println("Luas\t:", bangunDatar.luas(), "cm")
}
