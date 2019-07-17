package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//练习 2.2： 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的
//话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对
//应英尺和米，重量单位可以对应磅和公斤等。

func main() {
	if len(os.Args) == 1 {
		//无参数输入,根据标准输入获得输入值
		fmt.Print("请输入长度:")
		input := bufio.NewScanner(os.Stdin)
		input.Scan();
		lstr := input.Text()
		convAndPrint(lstr)
	} else {
		//有参数输入,循环转换
		for _, arg := range os.Args[1:] {
			convAndPrint(arg)
		}
	}
}

//字符转浮点数并打印单位转换结果
func convAndPrint(lstr string) {
	l, err := strconv.ParseFloat(lstr, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch2e2_lengthConv: %v\n", err)
		os.Exit(1)
	}
	m := Meter(l)
	f := Foot(l)
	fmt.Printf("%s=%s, %s=%s\n", m, mtof(m), f, ftom(f))
}

//长度转换
type Meter float64 //米
type Foot float64  //英尺

//转换率
const convRate = 3.2808399

func (m Meter) String() string {
	return fmt.Sprintf("%g米", m)
}

func (f Foot) String() string {
	return fmt.Sprintf("%g英尺", f)
}

func mtof(m Meter) Foot {
	return Foot(m * convRate)
}

func ftom(f Foot) Meter {
	return Meter(f / convRate)
}
