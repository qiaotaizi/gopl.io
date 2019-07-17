package main

import (
	"bufio"
	"fmt"
	"os"
)

//练习 1.4： 修改 dup2 ，出现重复的行时打印文件名称。
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		//没有输入文件,则还是调用标准输入来读取
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				//注意: go语言官方的异常打印方式
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			//不建议在循环中调用defer,会造成资源不足
			//可以把for循环内部逻辑抽取成一个函数,在函数中使用defer
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	//读取文件
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if counts[text] > 1 {
			fmt.Println(f.Name())
		}
	}
}
