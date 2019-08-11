package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn,err:=net.Dial("tcp","localhost:8000")
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout,conn)//读取服务端输出,使用标准输出流进行打印
	mustCopy(conn,os.Stdin)//读取客户端输入,发送给服务端
}


func mustCopy(dst io.Writer,src io.Reader){
	if _,err:=io.Copy(dst,src);err!=nil{
		log.Fatal(err)
	}
}