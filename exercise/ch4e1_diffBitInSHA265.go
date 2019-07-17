package main

import (
	"crypto/sha256"
	"fmt"
	"gopl.io/ch2/popcount"
)

//练习 4.1： 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的
//PopCount函数。)

func main() {
	c1:=sha256.Sum256([]byte{'x'})
	c2:=sha256.Sum256([]byte{'X'})
	fmt.Println(diffBitCountInSHA256Hashes(c1,c2))
}

func diffBitCountInSHA256Hashes(b1,b2 [32]byte) (count int){
	//同时遍历两个数组
	//同索引元素按位异或
	//调用popcount函数计算1的数量
	//累加

	for i:=0;i< len(b1);i++{
		e1:=b1[i]
		e2:=b2[i]
		r:=e1 ^ e2
		pop:=popcount.PopCount(uint64(r))
		count+=pop
		fmt.Printf("%b ^ %b = %b (%d)\n",e1,e2,r,pop)
	}
	return
}


