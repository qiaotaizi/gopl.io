package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(string2unicode(os.Args[1]))
	} else {
		fmt.Println("请输入要转unicode的文字")
	}
}

func string2unicode(s string) string {
	r := []rune(s)
	var buf bytes.Buffer
	for _, ru := range r {
		if ru<128 {
			buf.WriteRune(ru)
		} else {
			buf.WriteByte('\\')
			buf.WriteByte('u')
			buf.WriteString(fmt.Sprintf("%X", ru))
		}
	}
	return buf.String()
}
