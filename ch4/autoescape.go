package main

import (
	"html/template"
	"log"
	"os"
)

const testTempl= `<p>A: {{.A}}</p>
<p>B: {{.B}}</p>`

func main() {

	t:=template.Must(template.New("escape").Parse(testTempl))
	var data struct{
		A string  //普通字符串
		B template.HTML  //受信任的html代码
	}
	data.A="<b>Hello!</b>"//由于A是不受信任的html代码,<>符号会被转义为&lt;&gt;
	data.B="<b>Hello!</b>"//B是受信任的html代码,不会发生转义
	if err:=t.Execute(os.Stdout,data);err!=nil{
		log.Fatal(err)
	}
	
}
