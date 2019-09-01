package main

import (
	"fmt"
	"github.com/qiaotaizi/gopl.io/ch3/printints"
)

func main() {
	s:=printints.IntsToString([]int{1,2,3,4,5})
	fmt.Println(s)
}
