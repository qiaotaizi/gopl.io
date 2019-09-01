package main

import (
	"fmt"
	"github.com/qiaotaizi/gopl.io/ch3/basename"
)

func main() {
	fileName:="/c/sourceCode/ttpai/ttpai_api_new/pom.xml"
	s1:=basename.BaseName1(fileName)
	fmt.Println(s1)
	s2:=basename.BaseName2(fileName)
	fmt.Println(s2)
}
