package main

import (
	"fmt"
	"gopl.io/ch6/geometry"
)

func main() {
	o:=geometry.Point{1,1}
	fmt.Println(o)
	o.ScaleBy(5)
	fmt.Println(o)
	p:=geometry.Point{1,1}
	fmt.Println(p)
	p.ScaleBy2(5)
	fmt.Println(p)
}
