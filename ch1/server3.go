package main

import (
	"fmt"
	"log"
	"net/http"
)

//打印请求参数
//把莉萨如接进去
//运行:
//go run ./ch1/server3.go ./ch1/lissajous.go
func main() {
	http.HandleFunc("/",handler3)
	http.HandleFunc("/lissa", func(writer http.ResponseWriter, request *http.Request) {
		lissajous(writer)
	})
	http.ListenAndServe("localhost:8000",nil)
}

func handler3(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"%s %s %s\n",r.Method,r.URL,r.Proto)
	for k,v:=range w.Header(){
		fmt.Fprintf(w,"Header[%q] = %q\n",k,v)
	}
	fmt.Fprintf(w,"Host = %q\n",r.Host)
	fmt.Fprintf(w,"RemoteAddr = %q\n",r.RemoteAddr)

	if err:=r.ParseForm();err!=nil{
		log.Print(err)
	}
	for k,v:=range r.Form{
		fmt.Fprintf(w,"Form[%q] = %q\n",k,v)
	}
}
