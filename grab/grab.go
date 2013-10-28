package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"os"
	"log"
)

func main(){
	grab_to_file( "http://store.apple.com/cn/browse/home/shop_mac/family/macbook_air", "air.html")

	return

	fmt.Println("hello, grab")
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("http.Get failed:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("resp=", resp )

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Reading body failed:", err)
		return
	}
	textBody := string(body)
	fmt.Println("Reading body success:", textBody )

	if strings.Contains( textBody, "加入营销计划" ){
		fmt.Println("It is google.com.hk")
	} else {
		fmt.Println("It is not google.com.hk")
	}

	f, err := os.OpenFile( "google.html", os.O_WRONLY|os.O_CREATE, 0666 )
	if err != nil {
		fmt.Println("Open file google.html failed!")
		return
	}

	f.Write( body)
	f.Close()

}

func test(){
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)
	ioutil.WriteFile("robots.txt", robots, 0666 )

}

func grab_to_file( from string, to string ){
	res, err := http.Get(from)
	if err != nil {
		log.Fatal( err )
	}

	body, err := ioutil.ReadAll( res.Body )
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile( to, body, 0666 )
}
