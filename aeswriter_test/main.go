package main

import (
    "aesrw"
    "fmt"
)

type TestReaderWriter struct{
    buff []byte
}

func MakeTestReaderWriter () (*TestReaderWriter){
    t := TestReaderWriter{}
    t.buff = make( []byte, 0, 2048 )
    return &t
}

func (t *TestReaderWriter) Write(b []byte) (int, error) {
    fmt.Println("TestReaderWriter writting", b)
    t.buff = append( t.buff, b...)
    return len(b), nil
}

func (t *TestReaderWriter) Read(b []byte) (n int, err error) {
    todo := aesrw.Min( len(b), len(t.buff) )
    fmt.Println("len(b)=", len(b), "len(t.buff)=", len(t.buff), "todo=", todo)
    copy( b[:todo], t.buff[:todo] )
    t.buff = t.buff[todo:]
    return todo, nil
}

func main(){
    t := TestReaderWriter{}
    w,e :=aesrw.MakeAESWriter(&t )
    if e != nil {
        fmt.Println("MakeAESWriter failed", e)
        return
    }

    w.Write( []byte("1234567890123456") )
    w.Write( []byte("1234567890123456") )
    w.Write( []byte("1234567890123456") )


    r,e := aesrw.MakeAESReader( &t )
    if e != nil {
        fmt.Println("MakeAESReader failed!", e)
        return
    }

    one := make([]byte, 16)
    fmt.Println("len(one)=", len(one))
    for i := 0; i < 3; i++ {
        l, e := r.Read( one[:16] )
        fmt.Println("l=", l, "e=", e, "one=", one, string(one))
    }
    
    fmt.Println("hello, world")
}
