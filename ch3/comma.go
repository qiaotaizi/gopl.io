package main

import (
	"fmt"
	"github.com/qiaotaizi/gopl.io/ch3/comma"
)

func main(){
	num:="12345678"
	num=comma.Comma(num)
	fmt.Println(num)
}
