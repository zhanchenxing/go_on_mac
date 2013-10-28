package main

import (
	"encoding/xml"
	"fmt"
	"os"
)


type MyConfig struct {
//	XMLName xml.Name `xml:"Data"`
	XMLName2 xml.Name `xml:"Data2"`
	Myname string `xml:"myname"`
	Myage int `xml:"myage"`

	Mysubs []Sub `xml:"sub"`
}
func (c MyConfig) String() string {
	return fmt.Sprintf( "XMLName: v, name: %v, age: %v, Mysubs:%v", /*c.XMLName,*/ c.Myname, c.Myage, c.Mysubs )
}

type Sub struct {
	Where string `xml:"where,attr"`
	Sub1 string `xml:"sub1"`
	Sub2 string `xml:"sub2"`
}

func (s Sub) String() string {
	return fmt.Sprintf("Sub1=%v, Sub2=%v, Where=%v", s.Sub1, s.Sub2, s.Where )
}


func main() {
	xmlFile, err := os.Open("simple.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	buff := make([]byte, 16*1024)
	readed, err := xmlFile.Read(buff)
	if err != nil {
		fmt.Println("read error:", err)
	}
	fmt.Println("readed=", readed)

	var q MyConfig
	err = xml.Unmarshal(buff[0:readed], &q)
	if err != nil {
		fmt.Println("Unmarshal failed:", err)
	}

	fmt.Println( q )
}
