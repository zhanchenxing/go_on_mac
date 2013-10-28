package main

import (
    "net/rpc"
    "log"
    "rpctest/server"
)

func main(){

    log.Println( "Trying to connect rpc server..." )
    client, err := rpc.DialHTTP( "tcp", "127.0.0.1:1234" )
    if err != nil {
        log.Fatal( "connect rpc server failed:", err )
    }

    args := &server.Args{7, 8}
    var reply int

    err = client.Call("Arith.Multiply", args, &reply )
    if err != nil {
        log.Fatal("arith error:", err )
    }

    log.Println( "Arith:", args.A, "*", args.B, "=", reply )

}
