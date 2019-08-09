package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var callNum int

func main() {
	if l:=len(os.Args);l!=2{
		fmt.Fprint(os.Stderr,"请输入正确的整数类型参数\n")
		os.Exit(1)
	}

	ns:=os.Args[1]
	n,err:=strconv.Atoi(ns)
	if err!=nil{
		fmt.Fprintf(os.Stderr,"%v\n",err)
		os.Exit(1)
	}
	go spinner(100 * time.Millisecond)
	fibN := fib2(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	fmt.Printf("fib函数被调用%d次", callNum)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

//求斐波那契数列第x位的值
//这是一个很低效的算法
//目的是为了在命令行运行时显示协程的效果
func fib(x int) int {
	callNum++
	if x < 2 {
		return x
	}
	return fib(x-2) + fib(x-1)
}

func fib2(x int) int{
	callNum++
	if x<2{
		return x
	}
	a,b,t:=0,1,0
	for i:=0;i<x-1;i++{
		t=a
		a=b
		b=t+a
	}
	return b
}
