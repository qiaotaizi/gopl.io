package memo1

import (
	"fmt"
	"gopl.io/ch9/httpgetbody"
	"testing"
	"time"
)

func TestMemo1(t *testing.T){
	m:=New(httpgetbody.HttpGetBody)
	incomingUrls:=[]string{
		"https://www.baidu.com",
		"https://m.ttpai.cn",
		"https://www.ttpai.cn",
		"https://www.baidu.com",
		"https://www.baidu.com",
	}
	for _,url:=range incomingUrls{
		//并发形式访问url,会发现出现了缓存未命中的情况
		//这个时候可以使用命令
		//go test -run=TestMemo1 -race -v ch9/memo1/memo1_test.go ch9/memo1/memo1.go
		//来检查goroutine运行的情况
		go visit(m,url)
		//输出结果如下:
		//==================
		//WARNING: DATA RACE
		//Write at 0x00c0001c6708 by goroutine 7:
		//  command-line-arguments.(*Memo).Get()
		//      G:/goweb/src/gopl.io/ch9/memo1/memo1.go:28 +0x1f6
		//  command-line-arguments.visit()
		//      G:/goweb/src/gopl.io/ch9/memo1/memo1_test.go:33 +0x89
		//
		//Previous write at 0x00c0001c6708 by goroutine 10:
		//  command-line-arguments.(*Memo).Get()
		//      G:/goweb/src/gopl.io/ch9/memo1/memo1.go:28 +0x1f6
		//  command-line-arguments.visit()
		//      G:/goweb/src/gopl.io/ch9/memo1/memo1_test.go:33 +0x89
		//
		//Goroutine 7 (running) created at:
		//  command-line-arguments.TestMemo1()
		//      G:/goweb/src/gopl.io/ch9/memo1/memo1_test.go:20 +0x144
		//  testing.tRunner()
		//      C:/Go/src/testing/testing.go:827 +0x169
		//
		//Goroutine 10 (finished) created at:
		//  command-line-arguments.TestMemo1()
		//      G:/goweb/src/gopl.io/ch9/memo1/memo1_test.go:20 +0x144
		//  testing.tRunner()
		//      C:/Go/src/testing/testing.go:827 +0x169
		//==================
		//可以看到
	}
}


func visit(m *Memo,url string) {
	defer func() func(){
		start:=time.Now()
		return func() {
			d:=time.Since(start)
			fmt.Println("visiting",url,"spent",d,"seconds")
		}
	}()()
	m.Get(url)
}

/*
visiting https://www.baidu.com spent 150.7692ms seconds
visiting https://m.ttpai.cn spent 135.7814ms seconds
visiting https://www.ttpai.cn spent 329.619ms seconds
visiting https://www.baidu.com spent 0s seconds
visiting https://www.baidu.com spent 0s seconds
 */
