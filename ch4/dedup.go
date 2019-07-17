package main

import (
	"bufio"
	"fmt"
	"os"
)

//运行程序后,使用ctrl+z向标准输入流发送EOF信号,终结输入

func main() {
	seen:=make(map[string]bool)
	input:=bufio.NewScanner(os.Stdin)
	for input.Scan(){
		line:=input.Text()
		//键line不存在时,seen[line]将返回bool的零值:false
		if !seen[line]{
			seen[line]=true
			fmt.Println(line)
		}
	}

	if err:=input.Err();err!=nil{
		fmt.Fprintf(os.Stderr,"dedup: %v\n",err)
		os.Exit(1)
	}
}
