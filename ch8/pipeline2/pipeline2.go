package main

import "fmt"

//单方向的channel
//无法向只读channel写入
//无法从只写channel读取
//关闭只读channel时会出现编译错误

func main() {
	naturals:=make(chan int)  //无缓存channel的容量和长度都是0
	square:=make(chan int)
	fmt.Println(cap(naturals),len(naturals))
	fmt.Println(cap(square),len(square))
	go counter(naturals)
	go squarer(naturals,square)
	printer(square)

}

//传入的channel out隐式转换为只写的channel
func counter(out chan<- int){
	for i:=0;i<10;i++{
		out<-i
	}
	close(out)
}

//传入的channel out隐式转换为只写的channel
//传入的channel in隐式转换为只读的channel
func squarer(in <-chan int,out chan<- int){
	for x:=range in{
		out<-x*x
	}
	close(out)
}

func printer(out <-chan int){
	for x:=range out{
		fmt.Println(x)
	}
}
