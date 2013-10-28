package main

import (
    "fmt"
    "net/rpc"
)

func main(){
    client, err := rpc.DialHTTP( "tcp", "127.0.0.1:1234" )

    if err != nil {
        fmt.Println("connect rpc server failed!")
        return
    }

    var reply int
    err = client.Call( "Watcher.GetInfo", 1, &reply )

    if err != nil {
        fmt.Println("client.Call failed!")
        return
    }

    fmt.Println("client.Call returned:", reply )

}

