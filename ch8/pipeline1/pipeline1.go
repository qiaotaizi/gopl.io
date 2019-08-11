package main

import (
	"fmt"
)

//串联的channels
//pipeline

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	go func() { //计数协程
		for i := 0; i < 10; i++ {
			println("counter 协程输出一个值之前")
			naturals <- i
			println("counter 协程输出一个值")
		}
		close(naturals) //关闭channel,可以在读取时进行通知
	}()

	go func() { //求平方协程
		for x:=range naturals{//使用简洁的range语法代替死循环,当natural中没有值且被关闭,循环自动跳出
			//x,ok := <-naturals
			//if !ok{ //非ok说明信道已经关闭且其中没有值可以接收,需要进一步关闭squares
			//	close(squares)
			//	break //跳出死循环
			//}
			println("square 协程输入一个值")
			squares <- x * x//仅在ok的情况下需要输出x
			println("square 协程输出一个值")
		}
		close(squares)
	}()

	for x:=range squares{ //主协程
		//x,ok := <-squares
		//if !ok{
		//	break//信道关闭后跳出死循环
		//}
		println("主协程输出一个值")
		println(x)
	}

}
