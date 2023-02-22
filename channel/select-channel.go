package main

import (
	. "fmt"
	"runtime"
	"time"
)

func getAverage(numbers []int, ch chan float64) {
	var sum int
	for _, j := range numbers {
		sum += j
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max int = numbers[0]
	for _, k := range numbers {
		if max < k {
			max = k
		}
	}
	ch <- max
}

func main() {
	runtime.GOMAXPROCS(2)

	var numbers = []int{23, 12, 46, 23}
	Println("Numbers:", numbers)

	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)

	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	time.Sleep(time.Second)
	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			Printf("Avg: %.2f \n", avg)
		case max := <-ch2:
			Printf("Max: %d \n", max)
		}
	}
}
