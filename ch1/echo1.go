package main

import (
	"fmt"
	"os"
)

func main() {
	var s,sep string
	//os.Args索引0是命令本身
	for i:=1;i<len(os.Args);i++{
		s+=sep+os.Args[i]
		sep=" "
	}
	fmt.Println(s)
}
