package display

import (
	"fmt"
	"github.com/qiaotaizi/gopl.io/ch12/format"
	"reflect"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	//获取接口值的值
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
	//获取接口值的值的类型(说的有点绕,但是应该这样说)
	switch v.Kind() {
	case reflect.Invalid:
			fmt.Printf("%s = invalid\n",path)
	case reflect.Slice,reflect.Array:
		for i:=0;i<v.Len();i++{
			display(fmt.Sprintf("%s[%d]",path,i),v.Index(i))//value.Index方法仅能对Slice\数组\字符串类型的值调用,否则
			// 将产生panic
		}
	case reflect.Struct:
		//value.NumField方法返回结构体中成员的数量
		for i:=0;i<v.NumField();i++{
			fieldPath:=fmt.Sprintf("%s.%s",path,v.Type().Field(i).Name)//v.Type().Field(i)获取值的类型的第i个参数的类型,调用其Name属性
			display(fieldPath,v.Field(i))//v.Field(i)获取第i个参数的值
		}
	case reflect.Map:
		for _,key:=range v.MapKeys(){
			display(fmt.Sprintf("%s[%s]",path,format.FormatAtom(key)),v.MapIndex(key))//Mapindex()方法用于获取map类型的值中,键key对应的值
		}
	case reflect.Ptr:
		if v.IsNil(){//测试空指针
			fmt.Printf("%s = nil\n",path)
		}else{
			display(fmt.Sprintf("(*%s)",path),v.Elem())//Elem()方法返回指针指向的变量
		}
	case reflect.Interface:
		if v.IsNil(){//测试nil接口
			fmt.Printf("%s = nil\n",path)
		}else{
			fmt.Printf("%s.type = %s\n",path,v.Elem().Type())//获取接口值对应的动态值的类型
			display(path+".value",v.Elem())//获取接口值对应的动态值
		}
	default:
		fmt.Printf("%s = %s\n",path,format.FormatAtom(v))
	}
}
