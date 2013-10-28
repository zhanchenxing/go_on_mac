package main

import (
	"fmt"
	"time"
)

type Server struct {
	Quit chan bool
}

func NewServer() *Server{
	s := &Server{ make( chan bool ) }

	go s.run()
	return s
}

func (s *Server) run(){
	for{
		select{
		case <- s.Quit:
			fmt.Println("quit...........")
			return
		case <- time.After( time.Second ):
			fmt.Println("runing task")
		}
	}
}

func (s *Server) Stop(){
	fmt.Println("Server stopping")
	s.Quit <- true
	<-s.Quit
	fmt.Println("server stopped")
}


func main(){
	s := NewServer()
	time.Sleep( 10*time.Second )
	fmt.Println("will quit...")
	s.Quit <- true
	fmt.Println("should quit...")
	fmt.Println("sleeping another 10 seconds")
	time.Sleep(10 * time.Second )
}

