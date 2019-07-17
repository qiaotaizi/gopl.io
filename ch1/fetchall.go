package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//并发获取多个url的资源

func main() {

	start :=time.Now()
	ch := make(chan string)//无缓冲通道
	for _,url:=range os.Args[1:]{
		go fetch(url,ch);
	}

	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}

//ch是仅能向其发送消息的通道
func fetch(url string,ch chan<- string){
	start := time.Now()
	resp,err:=http.Get(url)
	//这句如果出现err,resp.Body是空的,直接在这里写defer Close会报空指针
	if err!=nil{
		//错误也发送给通道
		//保持对称
		ch<-fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()
	resp.Body.Close()
	//舍弃消息
	nbytes,err:=io.Copy(ioutil.Discard,resp.Body)
	if err!=nil{
		ch<-fmt.Sprint(err)
		return
	}
	secs:=time.Since(start).Seconds();
	ch<-fmt.Sprintf("%.2fs %7d %s",secs,nbytes,url)
}
