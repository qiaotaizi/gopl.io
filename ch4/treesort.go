package main

import "fmt"

func main() {
	values:=[]int{2,4,6,4,5,1,0,45,67,4}
	Sort(values)
	fmt.Println(values)
}

//结构体不能包含一个自身类型的成员对象(但是这么写也不会引发编译错误)
//但可以包含自身指针类型的对象

type tree struct {
	value int
	left,right *tree
}

func Sort(values []int){
	var root *tree
	for _,v:=range values{
		root=add(root,v)
	}
	appendValues(values[:0],root)
}

func add(t *tree,value int) *tree{
	if t==nil{
		t=new(tree)
		t.value=value
		return t
	}
	if value<t.value{
		t.left=add(t.left,value)
	}else{
		t.right=add(t.right,value)
	}
	return t
}

func appendValues(values []int,t *tree)[]int{
	if t!=nil{
		values=appendValues(values,t.left)
		values=append(values,t.value)
		values=appendValues(values,t.right)
	}
	return values
}
