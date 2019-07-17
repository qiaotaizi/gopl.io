package main

import "fmt"

func main() {
	slice:=[]int{1,2,3,4,5,6,7,8,9}
	slice=rotate(slice,5)
	fmt.Println(slice)
}

//练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。

//将slice中的元素依次向左移动i位
//缺陷:需要创建另外一个数组对象
func rotate(slice []int,pos int)[] int{
	length:=len(slice)
	if pos>=length{
		pos=pos%length
	}
	if pos==0{
		return slice
	}
	slice2:=make([]int, length)
	for i:=length-1;i>=0;i--{
		slice2Pos:=i-pos
		if slice2Pos<0{
			slice2Pos=length+slice2Pos
		}
		fmt.Println(slice2Pos)
		slice2[slice2Pos]=slice[i]
	}
	return slice2
}
