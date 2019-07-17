package main

import (
	"fmt"
	"log"
	"net/http"
)

//go run src/gopl.io/ch1/server1/main.go &
//后台运行服务

func main() {
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
}
