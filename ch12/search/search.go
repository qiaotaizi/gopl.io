package search

import (
	"fmt"
	params "gopl.io/ch12/param"
	"net/http"
)

func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels    []string `http:"l"`
		MaxResult int      `http:"max"`
		Exact     bool     `http:"x"`
	}
	data.MaxResult=10
	if err:=params.Unpack(req,&data);err!=nil{
		http.Error(resp,err.Error(),http.StatusBadRequest)
		return
	}


	fmt.Fprintf(resp,"Search: %+v\n",data)
}
