package geometry

import (
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
	//Hypot(a,b)函数 返回(a*a+b*b)的算术平方根
}

//将Distance函数绑定至Point类,成为方法,方法中使用p表示this
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//定义Path类型,底层是Point切片
type Path []Point

//给Path类绑定Distance方法
func (p Path) Distance() float64 {
	sum:=0.0
	for i:=range p{
		if i>0{
			sum+=p[i-1].Distance(p[i])//这里调用的是Point类的Distance方法
		}
	}
	return sum
}

//这里为什么要绑定至指针呢?
//方法的调用者也被是为方法的一个参数
//在传参时,如果传递的不是指针,那么调用者对象会被拷贝一份
//在函数中对调用者属性的所有操作都是在拷贝上执行的
//结果调用者本身不会发生变化
//例如下面的ScaleBy2
//调用完毕后再获取X,Y的值还是原值
func (p *Point) ScaleBy(factor float64){
	p.X*=factor
	p.Y*=factor
}
//另外,如果一个类型本身就是指针
//例如
//type T *int
//这种类型是不可以绑定方法的
//调用指针方法时,调用者必须可以取址
//Point{1,2}.ScaleBy(2)
//这种调用将出现编译错误
//因为临时变量Point{1,2}是不能取址的

func (p Point) ScaleBy2(factor float64){
	p.X*=factor
	p.Y*=factor
}

//func (p Point) String() string{
//	return fmt.Sprint(p.X,",",p.Y)
//}
