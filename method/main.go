package main

import (
	"fmt"
	"strings"
)

type mhs struct {
	nama string
	nim  int
}

func (m *mhs) updateData(nama string, nim int) {
	m.nama = nama
	m.nim = nim

	fmt.Println("Data has been updated!")
}

func (m mhs) getData() {
	fmt.Println("Nama\t:", m.nama)
	fmt.Println("NIM\t:", m.nim)
}

func (m mhs) getNickname(i int) {
	fmt.Println("Nickname:", strings.Split(m.nama, " ")[i-1])
}

func main() {
	fmt.Println("--------------------------")
	var mhs1 = mhs{"Muhammad Ali Mustaqim", 2003}
	mhs1.getData()
	mhs1.getNickname(2)

	fmt.Println("--------------------------")
	mhs1.updateData("Laurence NS", 2006)
	mhs1.getData()
	mhs1.getNickname(1)
}
