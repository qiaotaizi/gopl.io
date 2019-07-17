package main

import (
	"bytes"
	"fmt"
	"strings"
)

//练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。
//练习 3.11： 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。

func comma2(num string)string{
	var prefixLen int
	var buf bytes.Buffer
	if strings.HasPrefix(num,"-"){
		prefixLen=1
		buf.WriteByte('-')
	}
	l:=len(num)
	if l-prefixLen<=3{
		return num
	}

	num=num[prefixLen:]
	dot:=strings.LastIndex(num,".")
	var numSuf string
	if dot!=-1{
		//小数点存在
		numSuf=num[dot:]
		num=num[:dot]
	}
	lOfNum:= len(num)
	mod:=lOfNum%3
	if mod!=0 {
		buf.WriteString(num[:mod])
//		buf.WriteByte(',')
	}
	for i:=mod;i<lOfNum;i++{
		if i>0 && (i-mod)%3==0{
			buf.WriteByte(',')
		}
		buf.WriteByte(num[i])
	}
	buf.WriteString(numSuf)

	return buf.String()
}

func main() {
	str:="12342342.34544566"
	str=comma2(str)
	fmt.Println(str)


}


