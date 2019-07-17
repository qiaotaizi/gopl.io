package main

import (
	"crypto/sha256"
	"fmt"
)

//这段程序隐藏的知识点:
//sha256.Sum256返回的是[32]byte{}
//数组可比较
//%x格式化打印可以打印数组
//%t打印bool
//%T打印类型

func main() {
	c1:=sha256.Sum256([]byte{'x'})
	c2:=sha256.Sum256([]byte{'X'})
	fmt.Printf("%x\n%x\n%t\n%T\n",c1,c2,c1==c2,c1)
}
