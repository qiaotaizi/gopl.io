package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

//这个版本的du
//允许并发遍历多个目录
//但是限制最高并发遍历的目录数量

func main() {
	defer func() func() {
		start:=time.Now()
		return func() {
			dur:=time.Since(start)
			fmt.Printf("spent %vs\n",dur.Seconds())
		}
	}()()
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."} //无参数时,根目录指向当前目录
	}
	fileSizes := make(chan int64)
	var n sync.WaitGroup //额外开启一个线程,等待所有目录遍历完成,关闭fileSizes channel
	for _, root := range roots {
		n.Add(1) //遍历参数目录,n+1
		go walkDir(root, &n, fileSizes)//开启线程遍历目录
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbyte int64 //文件数,文件目录数
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				//这里为什么用ok标志判断信道是否关闭呢?
				//因为外层for循环没有range fileSize
				//如果这里不显示判断一下,讲会导致文件读取完毕时线程挂起
				//fileSizes信道已经关闭
				break loop //在select中跳出外层for循环
			}
			nfiles++
			nbyte += size
		case <-tick://如果verbose为false,tick为nil,select不会选择这个case
			//每当计时器跳数,输出一次目录信息
			printDiskUsage(nfiles, nbyte)
		}

	}
	printDiskUsage(nfiles, nbyte) //打印总量
}

//输出磁盘用量
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
