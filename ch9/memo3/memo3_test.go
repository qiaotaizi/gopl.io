package memo3

import (
	"fmt"
	"gopl.io/ch9/httpgetbody"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestMemo3(t *testing.T){
	m:=New(httpgetbody.HttpGetBody)
	incomingUrls:=[]string{
		"https://www.baidu.com",
		"https://m.ttpai.cn",
		"https://www.ttpai.cn",
		"https://www.baidu.com",
		"https://www.baidu.com",
	}
	for _,url:=range incomingUrls{
		wg.Add(1)
		//go test -run=TestMemo3 -race -v ch9/memo3/memo3_test.go ch9/memo3/memo3.go
		//来检查goroutine运行的情况
		go visit(m,url)
	}
	wg.Wait()
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
	wg.Done()
}
/*
visiting https://www.baidu.com spent 101.9427ms seconds
visiting https://www.baidu.com spent 103.9469ms seconds
visiting https://www.baidu.com spent 105.9194ms seconds
visiting https://m.ttpai.cn spent 193.9583ms seconds
visiting https://www.ttpai.cn spent 431.1873ms seconds
*/

