package main

import "fmt"

var v int64

const count = 100000000

func decre(c chan int) {
	for n := 0; n < count; n++ {
		v--
	}
	c <- 0
}

func main() {
	c := make(chan int)

	go decre(c)

	for n := 0; n < count; n++ {
		v++
	}

	_ = <-c

	fmt.Println("v=", v)
}
