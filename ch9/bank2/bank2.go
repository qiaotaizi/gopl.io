package main

func main() {
	
}

//使用二元信号量实现一个互斥锁
//调用存款或者查询函数时
//首先会向sema中传入一个信号
//由于sema有一位buffer,当前协程不会挂起
//如果这时协程切换,再向sema中传入信号时,会导致协程挂起
//直到传入信号地协程释放信号
//其他协程才可以继续进行操作
var (
	sema=make(chan struct{},1) //二元信号量
	balance int //余额变量是被共享的,而不是如bank1那样被监视协程控制着
)

func Deposite(amount int){
	sema<- struct{}{}
	balance+=amount
	<-sema
}

func Balance()int{
	sema<- struct{}{}
	b:=balance
	<-sema
	return b
}
