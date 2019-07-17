package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

//练习 4.8： 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等
//Unicode中不同的字符类别。

func main() {
	counts:=make(map[rune]int)
	var utflen [utf8.UTFMax+1]int
	invalid:=0
	letterCount:=0
	numCount:=0


	in:=bufio.NewReader(os.Stdin)

	for {
		r,n,err:=in.ReadRune()
		if err==io.EOF{
			break
		}
		if err!=nil{
			fmt.Fprintf(os.Stderr,"charcount: %v\n",err)
			os.Exit(1)
		}
		if r==unicode.ReplacementChar && n==1{
			invalid++
			continue
		}

		//统计数字的量
		if unicode.IsNumber(r){
			numCount++
		}

		//统计字母的量
		if unicode.IsLetter(r){
			letterCount++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c,n:=range counts{
		fmt.Printf("%q\t%d\n",c,n)
	}
	fmt.Print("\nlen\tcount\n")
	for i,n:=range utflen{
		if i>0{
			fmt.Printf("%d\t%d\n",i,n)
		}
	}
	if letterCount>0{
		fmt.Printf("\n%d letters\n", letterCount)
	}
	if numCount>0{
		fmt.Printf("\n%d numbers\n", numCount)
	}
	if invalid>0{
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

