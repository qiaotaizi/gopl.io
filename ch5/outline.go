package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	//结合fetch使用
	//使用方法见findlinks1.go
	doc,err:=html.Parse(os.Stdin)
	if err!=nil{
		fmt.Fprintf(os.Stderr,"outline: %v\n",err)
		os.Exit(1)
	}
	outline(nil,doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type==html.ElementNode{
		stack=append(stack, n.Data)
		fmt.Println(stack)
	}
	for c:=n.FirstChild;c!=nil;c=c.NextSibling{
		outline(stack,c)
		//这里传入stack时
		//在下一级递归的outline函数中使用的是stack的拷贝,
		// 当outline运行结束并返回上一级递归时,拷贝的stack将被释放
		//上一级递归的stack还是老样子

		//这也就是为什么append函数增加切片元素时,一定要重新赋值出去
	}
}
