package methods

import (
	"testing"
	"time"
)

type timeUtils struct {}

func (t timeUtils) Format(time time.Time)string{
	return ""
}

func (t timeUtils) Parse(str string)time.Time{
	return time.Time{}
}

func (t timeUtils) privateMethod(){

}

func TestPrint(t *testing.T) {
	var util timeUtils
	Print(util)//只能打印绑定至timeUtil的方法,无法打印绑定至*timeUtil的方法
	//只能打印导出的方法,无法打印私有的方法
}
