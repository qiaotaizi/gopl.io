package main

import "fmt"

func main() {
	//包含匿名成员结构体的初始化方法
	w := Wheel{
		Circle: Circle{
			Point:Point{0, 1},
			Radius:10},
		Spokes: 20}
	fmt.Println(w.Y)//直接调用匿名成员的叶子节点
	fmt.Printf("%#v",w)//结构体如何打印所有成员
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point //匿名成员 简化对象成员的链式调用语法,并且使对象获得匿名成员的所有方法
	Radius int
}

type Wheel struct {
	Spokes int
	Circle
}
