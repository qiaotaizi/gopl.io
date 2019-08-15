package main

import (
	"sync"
)

//使用sync包的工具来进行上锁

//再优化一下,
//将导出函数与实际执行逻辑的函数进行分离
//保证内部逻辑可以复用,保证一致性

var (
	mu      sync.Mutex
	balance int
)

//存款函数
func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func deposit(amount int) {
	balance += amount
}

//取款函数
//复用存款函数
//将一个负值传入
//如果余额消减后小于零
//进行回滚
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false;
	}
	return true
}

//余额查询函数
func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func main() {

}
