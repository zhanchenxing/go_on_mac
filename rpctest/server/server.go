package server

import (
    "errors"
    "fmt"
)

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

type Arith int

func (t *Arith) Multiply( args *Args, reply *int) error {
    fmt.Println("Multiply called with args=", args.A, args.B )
    *reply = args.A * args.B
    return nil
}

func (t *Arith) Divide( args *Args, quo *Quotient ) error {
    fmt.Println("Divide called with args=", args.A, args.B )
    if args.B == 0 {
        return errors.New("divide by zero!")
    }

    quo.Quo = args.A/args.B
    quo.Rem = args.A % args.B
    return nil
}


