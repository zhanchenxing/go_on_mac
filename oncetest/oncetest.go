package main

import (
	"fmt"
	"sync"
)

func main() {

	var once sync.Once

	onceBody := func() {
		fmt.Println("This line should be printed once!")
	}

	done := make(chan bool)
	for n := 0; n < 10; n++ {
		go func() {
			once.Do(onceBody)

			done <- true
		}()
	}

	for n := 0; n < 10; n++ {
		<-done
	}

}
