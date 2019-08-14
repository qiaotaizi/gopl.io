package main

import (
	"log"
	"net"
)

//聊天程序示例

func main() {
	listener,err:=net.Listen("tcp","localhost:8000")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("服务器启动")
	go broadcaster()
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
