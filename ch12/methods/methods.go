package methods

import (
	"fmt"
	"reflect"
	"strings"
)

func Print(x interface{}){
	v:=reflect.ValueOf(x)
	t:=v.Type()
	fmt.Printf("type %s\n",t)
	//t1:=reflect.TypeOf(x)
	//fmt.Printf("type %s\n",t1)
	//fmt.Printf("%d\n",v.NumMethod())
	for i:=0;i<v.NumMethod();i++{
		mt:=v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n",t,t.Method(i).Name,strings.TrimPrefix(mt.String(),"func"))
		fmt.Printf("%s\n",mt.String())
	}

	m,ok:=t.MethodByName("privateMethod")
	if ok{
		fmt.Println(m.Name)
	}else{
		fmt.Println("privateMethod not found")
	}
}
