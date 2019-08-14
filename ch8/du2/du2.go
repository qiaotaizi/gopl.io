package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//此方法递归地遍历目标文件夹
//每发现一个文件,将其大小发送至参数channel(只写)
func walkDir(dir string , fileSizes chan<- int64){
	//遍历文件系统头文件内容
	for _,entry:=range dirent(dir){
		//如果当前对象是目录
		if entry.IsDir(){
			//将目录名与前缀路径拼接,得到新的目录名
			subdir:=filepath.Join(dir,entry.Name())
			//递归调用
			walkDir(subdir,fileSizes)
		}else{
			//是文件
			//将文件大小发送给channel
			fileSizes<-entry.Size()
		}
	}
}

//返回文件系统头文件的信息
func dirent(dir string) []os.FileInfo{
	entries,err:=ioutil.ReadDir(dir)
	if err!=nil{
		fmt.Fprintf(os.Stderr,"du2: %v\n",err)
		return nil
	}
	return entries
}
