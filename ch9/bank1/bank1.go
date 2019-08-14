package main

//本文件的代码展示了一种规避并发问题的方式:
//balance(银行余额)是多个线程共享的数据
//但是该变量被封装在了teller协程中,作为一个后台协程运行
//这种协程被称为监视器协程
//基于select的作用,同一时间只有一个case会被select选中
//即balance同一时间只能和一个协程通信,避免了并发情况下的不安全问题

//GO语言口头禅:
//不要使用共享数据来通信
//使用通信来共享数据

func main() {

}

var deposits = make(chan int) //存款写入通道
var balances = make(chan int) //余额查询通道

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			//do nothing
		}
	}
}

func init(){
	go teller()
}
