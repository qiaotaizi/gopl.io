package httpgetbody

import (
	"io/ioutil"
	"net/http"
)

//这里为什么返回interface{},error
//而不是像ioutil.ReadAll那样返回byte[],error呢?
// TODO
func HttpGetBody(url string)(interface{},error){
	resp,err:=http.Get(url)
	if err!=nil{
		return nil,err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
