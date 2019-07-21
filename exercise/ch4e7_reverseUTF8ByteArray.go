package main

import "fmt"

func main() {

	s:="你好,世界!"
	s=string(reverseUTF8ByteArr([]byte(s)))
	fmt.Println(s)

}

//练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内
//存？

// TODO 不知道如何界定一个字符的边界,这个题目先搁置
func reverseUTF8ByteArr(slice []byte)[]byte{

	for i,j:=0, len(slice)-1;i<j;i,j=i+1,j-1{
		slice[i],slice[j]=slice[j],slice[i]
	}
	return slice
}
