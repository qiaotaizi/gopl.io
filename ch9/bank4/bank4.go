package main

import "sync"

//正常情况下
//读操作远远多余写操作
//如果大量读线程都要等待锁的释放,效率是比较低下的
//如果出现写操作,可能需要等待很多读线程结束才会轮到自己

//这种情况下,可以使用读写锁,降低只读线程之间的竞争

var (
	mu      sync.RWMutex //读写锁
	balance int
)

func Balance()int{
	mu.RLock()  //读锁
	defer mu.RUnlock()
	return balance
}

func Deposit(amount int)  {
	mu.Lock()  //写锁
	defer mu.Unlock()
	balance+=amount
}

func main() {

}
