package main

import (
	"fmt"
)

func main() {
	slice:=[]string{
		"寻","寻","觅","觅",",",
		"冷","冷","清","清",",",
		"凄","凄","惨","惨","戚","戚",",",
		"乍","暖","还","寒","时","候",",",
		"最","难","将","息","."}
	slice=nonRepeat(slice)
	fmt.Println(slice)
}

//练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。

func nonRepeat(slice []string)[]string{
	var temp string
	i:=0
	for _,s:=range slice{
		if temp!=s{
			slice[i]=s
			i++
		}
		temp=s
	}
	return slice[:i]
}
