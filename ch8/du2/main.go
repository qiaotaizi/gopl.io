package main

import (
	"flag"
	"fmt"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."} //无参数时,根目录指向当前目录
	}
	fileSizes := make(chan int64)
	go func() { //副routing,遍历目录,将每个文件的结果发送给channel,由主线程进行读取并累加
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
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
		case <-tick:
			//每当计时器跳数,输出一次目录信息
			printDiskUsage(nfiles, nbyte)
		}

	}
	printDiskUsage(nfiles, nbyte)//打印总量
}

//输出磁盘用量
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
