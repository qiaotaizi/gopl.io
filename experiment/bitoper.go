package main

import "fmt"

//位运算实验

func main() {
	//fmt.Println(isOdd(1234))
	n:=807666
	key:=807666
	en:=encrypt(n,key)
	fmt.Println(en)
	fmt.Println(encrypt(en,key))
}

//判断一个整数是否是奇数
//使用按位与取末尾,如果是1,就是奇数,反之为偶数
func isOdd(n int)bool{
	return n & 1==1
}

//使用异或运算加密和解密
//调用第一次加密
//调用第二次解密
func encrypt(n,key int)int{
	return n ^ key
}
