package main

import (
	"fmt"
	"image/color"
	"os"
	"encoding/binary"
	"io"
)

type Gopher struct {
	Name     string
	Age      int32
	FurColor color.Color
}

func (g *Gopher) DumpBinary(w io.Writer) error {
	err := binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	_, err = w.Write([]byte(g.Name))
	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	err = binary.Write(w, binary.LittleEndian, g.Age)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	return binary.Write(w, binary.LittleEndian, g.FurColor)
}

func main(){

	g := Gopher{ "zhanchenxing", 30, color.RGBA{255,0,0,0}}
	g.DumpBinary( os.Stdout )

	binary.Write( os.Stdout, binary.LittleEndian, "hello")

}
