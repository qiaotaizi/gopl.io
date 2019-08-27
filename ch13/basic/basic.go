//unsafe包的基础知识

package basic

import (
	"fmt"
	"unsafe"
)

type testStruct1 struct {
	param1 bool
	param2 float64
	param3 int16
}

type testStruct2 struct {
	param1 float64
	param2 bool
	param3 int16
}

func SizeOf(){
	tests:=[]interface{}{
		float64(0),
		int8(1),
		testStruct1{
			param1: false,
			param2: 0,
			param3: 0,
		},
		testStruct2{
			param1: 0,
			param2: false,
			param3: 0,
		},
	}
	for _,test:=range tests{
		printSize(test)
	}
}

func printSize(x interface{}){
	fmt.Printf("size of %T = %d\n",x,unsafe.Sizeof(x))
}
