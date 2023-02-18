package main

import (
	. "fmt"
	"reflect"
)

type student struct {
	Name   string
	Number int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	// cek apakah value s berupa pointer
	if reflectValue.Kind() == reflect.Pointer {
		reflectValue = reflectValue.Elem() // mengambil data dari pointer
	}

	var reflectType = reflectValue.Type() // mengambil data tipe

	for i := 0; i < reflectValue.NumField(); i++ {
		Println("Nama\t:", reflectType.Field(i).Name)        // berisi nama property yang menampung data
		Println("Tipe\t:", reflectType.Field(i).Type)        // berisi tipe data yang menampung data
		Println("Isi\t:", reflectValue.Field(i).Interface()) // isi dari propery yang dipanggil
		Println("")
	}
}

func main() {
	var nama string = "Muhammad Ali Mustaqim"
	var reflectValue = reflect.ValueOf(nama)
	Println("Isi \t\t:", reflectValue.String())
	Println("Tipe Data\t:", reflectValue.Type())
	Println("Tipe Data\t:", reflectValue.Kind())

	Println("------------------------------------")
	var nim interface{} = 75
	Println("Isi \t\t:", reflect.ValueOf(nim))
	Println("Tipe Data\t:", reflect.TypeOf(nim))

	Println("------------------------------------")
	var number interface{} = 75
	reflectValue = reflect.ValueOf(number)
	Println("Isi \t\t:", reflectValue.Interface().(int))
	Println("Isi \t\t:", reflectValue.Int())
	Println("Tipe Data\t:", reflectValue.Type())

	Println("------------------------------------")
	var student1 = &student{Name: "Dina", Number: 26}
	student1.getPropertyInfo()

	Println("------------------------------------")
	Println("Sebelum:", student1.Name)

	// memanggil method melalui reflect
	var reflectStudent = reflect.ValueOf(student1)
	var methodStudent = reflectStudent.MethodByName("SetName")
	name := []reflect.Value{
		reflect.ValueOf("Muhammad Ali"),
	}
	methodStudent.Call(name)

	Println("Sesudah:", student1.Name)
}
