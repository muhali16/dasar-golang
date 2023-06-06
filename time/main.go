package main

import (
	. "fmt"
	"time"
)

func main() {
	var now = time.Now()

	Println("year:", now.Year(), "month:", now.Month())
}
