package main

import (
	"fmt"
	"time"
	"errors"
	"math/rand"
)

func do(job string) error {
	seconds := rand.Int()%3
	time.Sleep( time.Duration(seconds) * time.Second )
	fmt.Println("Doing job", job, "sleep", seconds, "seconds")
	return errors.New( "Something went wrong!" )
}

func main(){
	//jobs := []string{ "one", "two", "three" }
	jobs := make( []string, 1000 )
	for n := 0; n < 1000; n++ {
		jobs[n] = fmt.Sprintf("%d-jobs", n+1)
	}


	errc := make( chan error )
	for _, jog := range( jobs ){
		go func( job string ) {
			errc <- do(job)
		}( jog )
	}

	for _ = range(jobs) {
		if err := <- errc; err!=nil {
			fmt.Println( err )
		}
	}
}
