package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s:="我是姜志恒";
	//RuneCountInString函数获取字符串中的字符数量
	fmt.Println(utf8.RuneCountInString(s))

	//DecodeRuneInString函数返回字符串中第一个字符的rune值和字节数
	r,i:=utf8.DecodeRuneInString(s)
	fmt.Println(r,i)
}
