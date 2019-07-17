package main

import "fmt"

//练习 3.12： 编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的
//字符，但是对应不同的顺序。

func c(s1, s2 string) bool {

	//1.长度相等
	//2.字符在两个字符串中出现的次数相等
	//3.不是所有位置上的字符都相同

	if len(s1) != len(s2) {
		return false
	}

	allTheSame:=true;

	for i, r := range s1 {
		occurInS1 := 0;
		for _, s := range s1 {
			if r == s {
				occurInS1++
			}
		}
		//fmt.Println(string(r), "在s1中出现了", occurInS1, "次")
		occurInS2 := 0
		for j, t := range s2 {
			if r == t {
				occurInS2++
			}

			if allTheSame && i==j && t!=r{
				allTheSame=false
			}
		}
		//fmt.Println(string(r), "在s2中出现了", occurInS1, "次")
		if occurInS2 != occurInS1 {
			return false
		}

	}

	return !allTheSame
}

func main() {
	b:=c("000a", "000a")
	fmt.Println(b)
}
