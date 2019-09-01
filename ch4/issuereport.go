package main
//注意这里导入的是text/template
import (
	"github.com/qiaotaizi/gopl.io/ch4/github"
	"log"
	"os"
	"text/template"
	"time"
)

func main() {
	result,err:=github.SearchIssues(os.Args[1:])
	if err!=nil{
		log.Fatal(err)
	}
	//模板对象将结果值设置进模板,并输出至对应的writer
	if err=report.Execute(os.Stdout,result);err!=nil{
		log.Fatal(err)
	}
}

const templ = `{{.TotalCount}} issues:
{{range .Items}}--------------------------
Number:	{{.Number}}
User:	{{.User.Login}}
Title:	{{.Title | printf "%.64s"}}
Age:	{{.CreatedAt | daysAgo}} days
{{end}}
`

//模板中的{{range}}{{end}}对表示循环
//{{.Value | function}}表示将Value传入function计算结果并填写至模板
//函数需要进行注册,方式见report变量的初始化

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report=template.Must(
	template.New("issueList").//对象初始化并命名
	Funcs(template.FuncMap{"daysAgo":daysAgo}).//函数映射
	Parse(templ))//初始化文本模板
