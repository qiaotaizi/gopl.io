package main

import (
	"flag"
	"fmt"
	"strings"
)

//指针是实现标准库中flag包的关键技术

//声明-n参数,用于忽略最后的换行符
var n = flag.Bool("n", false, "omit trailing newline")
//声明-s参数,用于指定输出参数的分隔符,默认是空格
var sep = flag.String("s", " ", "separator")

//flag.Bool和flag.String都返回了指针类型的变量

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep)) //这里对sep指针做一下寻址
	if !*n { //寻址
		fmt.Println()
	}
}
