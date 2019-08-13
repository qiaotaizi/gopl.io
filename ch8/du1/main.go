package main

import (
	"flag"
	"fmt"
)

//这个程序的问题在于:
//当扫描的目录中文件特别多
//程序需要等待很长时间才能给出反馈
//给用户一种假死的感觉
//在du2中将会改善这种状况

func main() {
	flag.Parse()
	roots:=flag.Args()
	//fmt.Println(roots)
	if len(roots)==0{
		roots=[]string{"."}//无参数时,根目录指向当前目录
	}
	fileSizes:=make(chan int64)
	go func() {//副routing,遍历目录,将每个文件的结果发送给channel,由主线程进行读取并累加
		for _,root:=range roots{
			walkDir(root,fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles,nbyte int64//文件数,文件目录数
	for size:=range fileSizes{
		nfiles++
		nbyte+=size
	}
	printDiskUsage(nfiles,nbyte)
}

//输出磁盘用量
func printDiskUsage(nfiles,nbytes int64)  {
	fmt.Printf("%d files, %.1f GB\n",nfiles,float64(nbytes)/1e9)
}
