package main

import (
	"fmt"
	"time"
)

type MyFile struct {
	fd      int
	name    string
	dirinfo string
	nepipe  int
}

func createMyFile() *MyFile {
	r := &MyFile{}
	fmt.Println("returning:", &r)
	return r
}

func Announce( message string, delay time.Duration ){
	go func(){
		time.Sleep(delay)
		fmt.Println("message=", message)
	}()
}

func main() {
	Announce( "System is going to shutdown.", time.Second * 3 )

	time.Sleep( time.Second*5)

	fmt.Println("This project is to study effective go.")
	var a = new(int)
	fmt.Println("a=", *a)

	// testing new 
	m1 := MyFile{}
	fmt.Println("m1=", m1)

	m2 := MyFile{1, "2222", "33333", 4}
	fmt.Println("m2=", m2)

	m3 := MyFile{fd: 1, nepipe: 6}
	fmt.Println("m3=", m3)

	// testing new
	n1 := new(MyFile)
	fmt.Println("n1=", n1)

	mf := createMyFile()
	fmt.Println("mf=", &mf)

	// 
	s := make( []int, 10, 100 )
	fmt.Println("s=", s)

}
