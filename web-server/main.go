package main

import (
	. "fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	Fprintln(w, "Apakabar?")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Fprintln(w, "Halo!")
	})

	http.HandleFunc("/index", index)

	http.HandleFunc("/mhs", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]interface{}{
			"Nama": "Muhammad Ali Mustaqim",
			"Npm":  202110715138,
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			Println(err)
			return
		}

		t.Execute(w, data)
	})

	Println("Starting web server at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
