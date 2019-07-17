package main

import (
	"fmt"
	"os"
	"strings"
)

//练习 1.1： 修改 echo 程序，使其能够打印 os.Args[0] ，即被执行命令本身的名字。

func main() {
	sep:=" "
	arg0:=os.Args[0]
	//截取一下
	idx:=strings.LastIndex(arg0,string(os.PathSeparator));
	arg0=arg0[idx+1:len(arg0)-4]
	fmt.Println(arg0+sep+strings.Join(os.Args[1:],sep))
}
