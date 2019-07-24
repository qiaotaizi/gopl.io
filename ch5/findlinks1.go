package main

import (
	"fmt"
	"golang.org/x/net/html"
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
	for _,link:=range visit(nil,doc){
		fmt.Println(link)
	}

}

func visit(links []string, n *html.Node) []string {
	//n是dom节点且为a标签
	if n.Type==html.ElementNode && n.Data=="a"{
		//遍历标签的属性,查找href属性
		for _,attr:=range n.Attr{
			if attr.Key=="href"{
				links=append(links,attr.Val)
				break
			}
		}
	}

	//声明c变量,赋值为n节点的第一个子节点
	//循环每进行一次,将c赋值为c的下一个兄弟节点
	//直到没有下一个兄弟节点
	for c:=n.FirstChild;c!=nil;c=c.NextSibling{
		links=visit(links,c)
	}
	return links
}
