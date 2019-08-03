package main

import "fmt"

func main() {

	//变长参数函数调用
	fmt.Println(sum(1,2,3,4))

	vals:=[]int{1,2,3,4}

	//如何将切片作为变长参数传入函数中
	fmt.Println(sum(vals...))

	//但是注意
	//可变长参数函数和直接以切片为参数的函数是不等价的
}

//变长参数
//对多个int类型的值求和
func sum(vals ...int) int{
	var result int
	for _,i:=range vals{
		result+=i
	}
	return result
}
