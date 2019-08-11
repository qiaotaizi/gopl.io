package main

import "fmt"

var ch =make(chan int)

func main() {
	//go spinner()
	syncSend()
	go syncReceive()
	//flag=true

	//ch:=make(chan int)
	//
	//go func() {
	//	ch <- 1
	//	fmt.Println("send",1)  // 1
	//}()
	//fmt.Println("start receive routine" ,2) //2
	//x := <-ch
	//fmt.Println(x,3) //3



	//按照书上的描述:
	// 当通过一个无缓存Channels发送数据
	//时，接收者收到数据发生在唤醒发送者goroutine之前
	//
	// 期望结果:
	//打印顺序:
	//2 1 3
}

var flag bool

func spinner(){
	for {
		if flag{
			break
		}
	}
}

func syncSend() {
	ch <- 1
	fmt.Println("send",1)  // 1
}

func syncReceive() {
	//time.Sleep(3*time.Second)
	fmt.Println("start receive routine" ,2) //2
	x := <-ch
	fmt.Println(x,3) //3
}
