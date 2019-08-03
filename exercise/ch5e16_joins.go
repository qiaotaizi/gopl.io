package main

import (
	"bytes"
	"fmt"
)

//练习5.16：编写多参数版本的strings.Join。

func main() {

	s:=joins(",","我","是","谁","?")
	fmt.Println(s)

}

func joins(sep string, a ...string) string {
	var buf bytes.Buffer
	l:=len(a)
	for i,s:=range a{
		buf.WriteString(s)
		if i<l-1{
			buf.WriteString(sep)
		}
	}
	return buf.String()
}
