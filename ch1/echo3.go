package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//如果存在大连拼接,strings.Join函数更加高效
	//用byte数组操作替代字符串的直接拼接操作
	fmt.Println(strings.Join(os.Args[1:]," "))
}
