package main

import (
	"fmt"
	"unicode"
)

func main() {
	s:="abc    defg  hijklm nopqr  stu  vwxyz   ";
	bs:=nonRepeatSpace([]byte(s))
	fmt.Println(string(bs)+";")
}

//练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考
//unicode.IsSpace） 替换成一个空格返回

func nonRepeatSpace(slice []byte)[]byte{
	lastByteIsSpace:=false
	i:=0
	for _,b:=range slice{
		if unicode.IsSpace(rune(b)){
			if !lastByteIsSpace{
				slice[i]=b
				i++
				lastByteIsSpace=true
			}
		}else{
			slice[i]=b

			i++
			lastByteIsSpace=false
		}
	}
	return slice[:i]
}
