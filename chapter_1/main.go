package main

import (
	"encoding/xml"
	"fmt"
)


type Foo struct {
	Bar string `xml:"id,attr"`
	Baz string `xml:"parent>child"`
}

func strlen(s string, c chan int) {
	c <- len(s)
}

func main() {
	c := make(chan int)
	go strlen("Salutations", c)
	go strlen("World", c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
	f := Foo{"Joe Junior", "Hello Shabado"}
	b, _ := xml.Marshal(f)
	fmt.Println(string(b))
	xml.Unmarshal(b, &f)
}
