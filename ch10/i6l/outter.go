package main

import (
	"github.com/qiaotaizi/gopl.io/ch10/i6l/internal"
	"github.com/qiaotaizi/gopl.io/ch10/i6l/internal/inner"
)

func main(){
	//internal包下的所有导出成员仅对i6l目录下有效
	//在其他包使用则出现编译错误
	internal.I6lFunc()
	inner.InnerFunc()
}

