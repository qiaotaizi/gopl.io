package main

import (
	"fmt"
	"os"
)

func main() {
	//unicode总共可以表现多少字符?
	//1比特:0|0000000~0|1111111
	//2比特:110|00001_10|000000~110|11111_10|111111
	//3比特:1110|0001_10|000000_10|000000~1110|1111_10|111111_10|111111
	//4比特:11110|001_10|000000_10|000000_10|000000~11110|111_10|111111_10|111111_10|111111
	bit1 :=btod2("1111111")-btod2("0")
	fmt.Println(bit1)
	bit2 := btod2("11111111111")-btod2("1000000")
	fmt.Println(bit2)
	bit3 := btod2("1111111111111111")-btod2("1000000000000")
	fmt.Println(bit3)
	bit4:=btod2("111111111111111111111")-btod2("1000000000000000000")
	fmt.Println(bit4)
	fmt.Println(bit1+bit2+bit3+bit4)

}

//2进制转10进制
//如:输入1110 返回14
//必须输入形如10101的整数
//否则直接返回0
func btod(b int) (d int) {

	for i:=0;b>0;i++{
		//除数
		//取余
		mod:=b%10
		if mod>1{
			fmt.Fprintf(os.Stderr,"btod: wrong binary number %d\n",b)
			os.Exit(1)
		}
		d+=f(2,i)*mod
		b/=10
	}
	return
}

func f( m,n int) (r int) {
	//返回m的n次方
	r=1
	for i:=0;i<n;i++{
		r*=m
	}
	return
}

func btod2(b string) (d int){
	//0=>48
	//1=>49
	for i,s:=range b{
		if s-48>1 || s-48<0{
			fmt.Fprintf(os.Stderr,"btod: wrong binary number %s\n",b)
			os.Exit(1)
		}
		d+=f(2,len(b)-i-1)*(int(s)-48)
	}
	return
}
