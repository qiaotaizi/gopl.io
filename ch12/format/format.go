package format

import (
	"reflect"
	"strconv"
)

// 将任意类型的参数格式化为string
func Any(value interface{})string{
	return formatAtom(reflect.ValueOf(value))
}

func FormatAtom(v reflect.Value)string{
	return formatAtom(v)
}

//原子格式化
//因为本函数将入参视为一个整体来进行格式化
func formatAtom(v reflect.Value) string {
	//获取值的类型
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		return strconv.FormatInt(v.Int(),10)
	case reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64,reflect.Uintptr:
		return strconv.FormatUint(v.Uint(),10)
	case reflect.Float32,reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f',2,64)
	case reflect.Complex64,reflect.Complex128:
		c:=v.Complex()
		return "("+strconv.FormatFloat(real(c),'f',2,64)+","+strconv.FormatFloat(imag(c),'f',2,64)+")"
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan,reflect.Func,reflect.Ptr,reflect.Slice,reflect.Map:
		return v.Type().String()+" 0x"+strconv.FormatUint(uint64(v.Pointer()),16)
	default://reflect.Array,reflect.Struct,reflect.Interface简单处理
		return v.Type().String()+" value"
	}
}
