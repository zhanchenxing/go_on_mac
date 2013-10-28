package main

import (
	"net/http"
	"sync"
	"fmt"
)

func main(){
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _, url := range(urls){
		wg.Add(1)
		go func (url string){
			text, err := http.Get(url)
			if err != nil {
				fmt.Println( url, "err=",err)
			} else {
				fmt.Println(text)
			}
			wg.Done()
		}( url )

	}

	wg.Wait()
}

