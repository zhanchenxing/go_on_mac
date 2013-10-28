package main

import (
    "fmt"
    "net"
    "net/rpc"
    "net/http"
    "rpc_lib"
)

func main(){

    wather := new( rpc_lib.Watcher )
    rpc.Register( wather )

    rpc.HandleHTTP()

    l, err := net.Listen( "tcp", ":1234")
    if err != nil {
        fmt.Println("Listen failed!")
        return
    }

    fmt.Println("Listening on 1234...")

    http.Serve( l, nil )
    fmt.Println("Server ended!")
}
