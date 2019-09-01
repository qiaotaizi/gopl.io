package main

import (
	"fmt"
	"golang.org/x/net/html"
	"github.com/qiaotaizi/gopl.io/ch5/visit"
	"os"
)

//函数的递归调用

//递归获取HTML的节点,收集其中的a标签的href

func main() {
	//这里使用标准输入流作为数据来源
	//例子中使用fetch程序输出html文档
	//使用|操作符链接fetch程序与本程序

	//运行方式:
	//编译fetch和findlinks
	//$ .../gopl.io
	//go build ch1/fetch.go
	//go build ch5/findlinks1.go
	//./fetch.exe http://www.baidu.com | ./findlinks1.exe

	doc,err:=html.Parse(os.Stdin)
	if err!=nil{
		fmt.Fprintf(os.Stderr,"findlinks1: %v\n",err)
		os.Exit(1)
	}
	for _,link:=range visit.Visit(nil,doc){
		fmt.Println(link)
	}
}
