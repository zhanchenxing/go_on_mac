package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"io"
	"time"
	"net/http"
)

func main(){
	fmt.Println("Hello, world!")
	conn, err := net.Dial("tcp", "localhost:8080" )
	checkError(err)

	reader := bufio.NewReader( conn )
	start_time := time.Now()
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF{
			fmt.Println("all readed!")
			end_time := time.Now()

			last_time := end_time.Sub(start_time)
			fmt.Println("time used:", last_time)

			os.Exit(0)
		}

		checkError(err)

		fmt.Println("line:", string(line) )
		fmt.Println("isPrefix:", isPrefix)

		_, err = http.Get(string(line))
		fmt.Println("error", err, ", url:", string(line))

		conn.Write([]byte("ok"))
	}
}

func checkError( err error ){
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
