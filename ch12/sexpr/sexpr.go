package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

func encode(buf *bytes.Buffer,v reflect.Value) error{
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		fmt.Fprintf(buf,"%d",v.Int())

	case reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64,reflect.Uintptr:
		fmt.Fprintf(buf,"%d",v.Uint())

	case reflect.Bool:
		var r rune
		if v.Bool(){
			r='t'
		}else{
			r='f'
		}
		fmt.Fprintf(buf,"%c",r)

	case reflect.String:
		fmt.Fprintf(buf,"%s",v.String())

	case reflect.Ptr:
		encode(buf,v.Elem())//获取指针指向的数据,进行递归encode

	case reflect.Array,reflect.Slice:
		buf.WriteByte('(')
		for i:=0;i<v.Len();i++{
			if i>0{
				buf.WriteByte(' ')
			}
			if err:=encode(buf,v.Index(i));err!=nil{
				return err
			}
			buf.WriteByte(')')
		}

	case reflect.Struct:
		buf.WriteByte('(')
		for i:=0;i<v.NumField();i++{
			if i>0{
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf,"(%s ",v.Type().Field(i).Name)
			if err:=encode(buf,v.Field(i));err!=nil{
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Map:
		buf.WriteByte('(')
		for i,key:=range v.MapKeys(){
			if i>0{
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err:=encode(buf,key);err!=nil{
				return err
			}
			buf.WriteByte(' ')
			if err:=encode(buf,v.MapIndex(key));err!=nil{
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	default:
		return fmt.Errorf("unsupport type %s",v.Type())
	}
	return nil
}

func Marshal(v interface{})([]byte,error){
	var buf bytes.Buffer
	if err:=encode(&buf,reflect.ValueOf(v));err!=nil{
		return nil,err
	}
	return buf.Bytes(),nil
}
