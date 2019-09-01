package main

//注意这里导入的是html/template
//html/template与text/template的区别
//会将html中的特殊字符进行转义处理

import (
	"github.com/qiaotaizi/gopl.io/ch4/github"
	"html/template"
	"log"
	"os"
)

//html模板数据填充

var htmlTempl = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

var issueList = template.Must(template.New("issueList").Parse(htmlTempl))

//执行程序并输出至文件
//go run issuehtml.go 参数列表 >xxx.html

func main() {
	result,err:=github.SearchIssues(os.Args[1:])
	if err!=nil{
		log.Fatal(err)
	}
	if err=issueList.Execute(os.Stdout,result);err!=nil{
		log.Fatal(err)
	}
}
