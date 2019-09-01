package memo2

import (
	"fmt"
	"github.com/qiaotaizi/gopl.io/ch9/httpgetbody"
	"testing"
	"time"
)

func TestMemo2(t *testing.T){
	m:=New(httpgetbody.HttpGetBody)
	incomingUrls:=[]string{
		"https://www.baidu.com",
		"https://m.ttpai.cn",
		"https://www.ttpai.cn",
		"https://www.baidu.com",
		"https://www.baidu.com",
	}
	for _,url:=range incomingUrls{
		//go test -run=TestMemo2 -race -v ch9/memo2/memo2_test.go ch9/memo2/memo2.go
		//来检查goroutine运行的情况
		go visit(m,url)
		//执行后没有出现race信息
		//但是发现最后三个请求都耗时648.1797ms
		//这是为什么呢?
		//因为前三个请求没有走缓存
		//任一时刻只能有一个请求拿到锁
		//等于是将这些请求串行化了
		//5个请求(也就是5个goroutine)几乎同时开始
		//第一个请求耗时181.0503ms
		//第二个请求耗时(227.0044-181.0503)ms
		//第三个请求耗时(648.1797-227.0044)ms
		//第四和第五个请求因为走了缓存,均耗时0ms
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
visiting https://m.ttpai.cn spent 181.0503ms seconds
visiting https://www.baidu.com spent 227.0044ms seconds
visiting https://www.ttpai.cn spent 648.1797ms seconds
visiting https://www.baidu.com spent 648.1797ms seconds
visiting https://www.baidu.com spent 648.1797ms seconds
 */
