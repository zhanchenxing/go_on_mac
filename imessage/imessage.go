package main

import (
	"fmt"
	"net/http"
	"text/template"
	"log"
)

var IndexTemplate *template.Template

func Index( w http.ResponseWriter, r *http.Request ){
	fmt.Println("Index Path =", r.URL.Path)
	path := r.URL.Path[1:]
	if path == "favicon.ico" {
		http.NotFound(w, r)
		return
	}

	//fmt.Fprintf( w, "%s", "hello, world" )
	IndexTemplate.Execute(w, nil)
}

func Send( w http.ResponseWriter, r *http.Request ){
	content := r.FormValue("content")
	to := r.FormValue("to")
	fmt.Println("content=", content)
	fmt.Println("to=", to)

	fmt.Fprintf( w, "Hello, %s", to )
}

func Terminate( w http.ResponseWriter, r *http.Request ){
	fmt.Fprintf(w, "Terminating" )
	log.Fatal("Terminating")
}

func main(){
	fmt.Println("hello, world!")
	var err error
	IndexTemplate, err = template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("index.html format error!")
		return
	}

	http.HandleFunc("/", Index )
	http.HandleFunc("/send", Send )
	http.HandleFunc("/exit", Terminate )

	err = http.ListenAndServe( ":8888", nil )
	if err != nil {
		log.Fatal(err)
	}
}

