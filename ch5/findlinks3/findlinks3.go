package main

import (
	"fmt"
	"github.com/qiaotaizi/gopl.io/ch5/links"
	"log"
	"os"
)

func main() {
	breadthFirst(crawl,os.Args[1:])
}


//广度优先遍历url中出现的a标签href属性
//如果一个html中引用了另一个html链接
//也会递归地访问那个html中地所有a标签
//由于不需要重发访问相同地html链接,会在内存中维护一个已经访问的href值的map
//运行过程中消耗的内存会不短增加
//当所有a标签被访问或者内存耗尽时,程序结束

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0{
		items:=worklist
		worklist=nil
		for _,item:=range items{
			if !seen[item]{
				seen[item]=true
				worklist=append(worklist,f(item)...)//将f(item)的结果视为变长元素设置入append函数
			}
		}
	}
}

func crawl(url string)[]string{
	fmt.Println(url)
	list,err:=links.Extract(url)
	if err!=nil{
		log.Print(err)
	}
	return list
}