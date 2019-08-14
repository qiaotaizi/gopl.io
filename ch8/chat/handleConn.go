package main

import (
	"bufio"
	"fmt"
	"net"
)

//处理客户端连接
//本身管理一个读线程
//并开启一个写线程
func handleConn(conn net.Conn){
	ch:=make(chan string)
	go clientWriter(conn,ch)

	who:=conn.RemoteAddr().String()
	ch<-"You are "+who
	messages<-who+" has arrived"
	entering<-ch

	input :=bufio.NewScanner(conn)
	for input.Scan(){
		messages<-who+": "+input.Text()
	}

	leaving<-ch
	messages<-who+" has left"
	conn.Close()
}

func clientWriter(conn net.Conn,ch <-chan string){
	for msg:=range messages{
		fmt.Fprintln(conn,msg)//不断从messages中读取消息并发送给客户端
	}
}
