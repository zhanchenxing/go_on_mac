package main

import (
    "log"
    "rpctest/server"
    "net/rpc"
    "net"
    "net/http"
)

func main(){

    arith := new (server.Arith)
    rpc.Register( arith )
    rpc.HandleHTTP()

    l, e := net.Listen( "tcp", ":1234" )
    if e != nil {
        log.Fatal( "Listen failed:", e )
    }


    log.Println( "Listening...")
    http.Serve( l, nil )
}
