package main

import (
    "errors"
    "net/rpc"
    "net"
    "net/http"
    "fmt"
)

type Args struct{
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

type Arith int

func (t *Arith) Multiply ( args *Args, reply *int ) error{
    fmt.Println("Multiply called...", args, reply)
    *reply = args.A * args.B
    return nil
}


func (t *Arith) Divide( args *Args, quo *Quotient ) error {
    fmt.Println("Divide called...", args, quo)
    if args.B == 0 {
        return errors.New("divide by zero")
    }

    quo.Quo = args.A / args.B
    quo.Rem = args.A % args.B
    return nil
}

func main(){
    arith := new (Arith)
    rpc.Register( arith )
    rpc.HandleHTTP()

    // 创建一个监听端
    l, e := net.Listen( "tcp", ":1234" )
    if e != nil {
        fmt.Println("Listen failed!")
        return
    }

    fmt.Println("Listening...")
    http.Serve(l, nil )
}
