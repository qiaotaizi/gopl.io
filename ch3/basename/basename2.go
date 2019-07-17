package basename

import "strings"

//使用strings库函数进行字符串操作,实现与BaseName1一样的功能
func BaseName2(s string) string{
	//最后一个'/'的位置
	slash:=strings.LastIndex(s,"/")
	s=s[slash+1:]
	dot:=strings.LastIndex(s,".")
	s=s[:dot]
	return s
}
