package main

import "fmt"

func main() {
	//声明一个数组
	a:=[...]int{1,2,3,4,5}
	reverse(a[:])//将a切片并作为reverse的入参
	fmt.Println(a)

	//将a中的元素整体向左移动3位
	rotate(a[:],3)
	fmt.Println(a)
}

//切片反转
func reverse(s []int){
	for i,j:=0, len(s)-1;i<j;i,j=i+1,j-1{
		s[i],s[j]=s[j],s[i]
	}
}

//将切片元素向左移动n位
func rotate(s []int,n int){
	//反转前n位
	reverse(s[:n])
	//反转剩余几位
	reverse(s[n:])
	//整体反转
	reverse(s)
}
