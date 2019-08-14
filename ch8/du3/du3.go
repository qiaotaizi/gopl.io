package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

//此方法递归地遍历目标文件夹
//每发现一个文件,将其大小发送至参数channel(只写)
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()//当本目录遍历结束时,n-1
	//遍历文件系统头文件内容
	for _, entry := range dirent(dir) {
		//如果当前对象是目录
		if entry.IsDir() {
			n.Add(1)//子目录,又要开启一个线程遍历,n+1
			//将目录名与前缀路径拼接,得到新的目录名
			subdir := filepath.Join(dir, entry.Name())
			//递归调用
			go walkDir(subdir, n, fileSizes)
		} else {
			//是文件
			//将文件大小发送给channel
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)//一个缓冲区为20的channel

//返回文件系统头文件的信息
func dirent(dir string) []os.FileInfo {
	sema <- struct{}{}//每当要遍历一个目录时,会调用dirent方法,向sema发送一个消息,如果并发数已经达到20,这里会挂起,直到有目录遍历完成
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3: %v\n", err)
		return nil
	}
	return entries
}
