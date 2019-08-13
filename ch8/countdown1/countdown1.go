package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick:=time.Tick(1*time.Second)
	for countdown:=10;countdown>0;countdown--{
		fmt.Println(countdown)
		<-tick//Tick函数返回一个只读channel(缓冲区为1),读取channel时,发生阻塞,每个周期往channel中传递一个值,读取goroutine将以构造方法定义的周期长度读取这个值
	}
	launch()
}

func launch(){
	fmt.Println("Lift off!")
}
