package main


import (
    "fmt"
    "net/rpc"
)

func main(){
    fmt.Println("hello, world!")

    fmt.Println( "Dialing http...")
    client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
    if err != nil {
        fmt.Println("DialHTTP failed:", err )
        return
    }

}
