package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown. Press enter to abort")
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //从标准输入流中读取一个字节
		abort <- struct{}{}
	}()
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
