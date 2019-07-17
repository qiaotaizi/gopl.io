package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//不使用流的形式读文件
//一次性把文件的所有内容读入内存

func main() {
	counts := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		//读取文件全部内容到一个byte数组
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
