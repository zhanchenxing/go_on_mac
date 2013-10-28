package main

import (
	"log"
	"os"
	"io"
)

func main() {
	logFile, err := os.OpenFile("log.txt", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666 )
	if err != nil {
		log.Fatal( err )
	}
	logFile2, err := os.OpenFile("log2.txt", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666 )
	if err != nil {
		log.Fatal( err )
	}

	multiLog := io.MultiWriter( logFile2, logFile, os.Stdout )

	log.SetOutput( multiLog)
	log.Print("This is log.Print")
	log.Println("This is log.Println")
	log.Println("This is log.Println")
	log.Println("This is log.Println")
	log.Fatal("This is log.Fatal")
	log.Fatal("This is log.Fatal")
	log.Fatal("This is log.Fatal")
	log.Fatal("This is log.Fatal")
	log.Fatal("This is log.Fatal")
	log.Fatal("This is log.Fatal")
	log.Panic("This is log.Panic");
	log.Panic("This is log.Panic");
	log.Panic("This is log.Panic");
	log.Panic("This is log.Panic");

}
