package main

import (
	"fmt"
	"os"
)

func main() {
	s,sep:="",""
	for _, arg:= range os.Args[1:]{//首个被舍弃的值是索引
		s+=arg+sep
		sep=" "
	}
	fmt.Println(s)
}
