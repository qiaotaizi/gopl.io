package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _,url:=range os.Args[1:]{
		doc,err:=findDoc(url)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"%v\n",err)
			continue
		}
		forEachNode(doc,startElement,endElement)
	}
}

func findDoc(url string) (*html.Node,error)  {
	resp,err:=http.Get(url)
	if err!=nil{
		return nil,err
	}
	if resp.StatusCode!=http.StatusOK{
		resp.Body.Close()
		return nil,fmt.Errorf("getting %s: %s",url,resp.Status)
	}
	doc,err:=html.Parse(resp.Body)
	if err!=nil{
		return nil,fmt.Errorf("parsing %s as HTML: %v",url,err)
	}
	return doc,err
}

//使用函数值实现aop
func forEachNode(n *html.Node,pre,post func(n *html.Node)){
	if pre!=nil{
		pre(n)
	}
	for c:=n.FirstChild;c!=nil;c=c.NextSibling{
		forEachNode(c,pre,post)
	}
	if post!=nil{
		post(n)
	}
}

var depth int

func startElement(n *html.Node){
	if n.Type==html.ElementNode{
		fmt.Printf("%*s<%s>\n",depth*2,"",n.Data)//打印缩进和标签头
		depth++
	}
}
func endElement(n *html.Node){
	if n.Type==html.ElementNode{
		depth--
		fmt.Printf("%*s</%s>\n",depth*2,"",n.Data)
		//格式化打印解释:
		//%*s填充depth*2数量的空格,再输出""
	}
}
