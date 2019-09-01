package main

import (
	"fmt"
	"github.com/qiaotaizi/gopl.io/thumbnail"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	defer func() func() {
		start := time.Now()
		fmt.Println("start=", start)
		return func() {
			end := time.Now()
			fmt.Println("end=", end)
			fmt.Println("last=", end.Sub(start))
		}
	}()()
	filenames := []string{
		"G:\\thumbnails\\avatar1.jpeg",
		"G:\\thumbnails\\avatar2.jpeg",
		"G:\\thumbnails\\avatar3.png",
		"G:\\thumbnails\\avatar4.jpg",
	}
	//makeThumbnails1(filenames)
	//makeThumbnails2(filenames)
	makeThumbnails3(filenames)
}

//串行创建图片
func makeThumbnails1(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

//简单并发创建图片
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go func(file string) {
			_, err := thumbnail.ImageFile(file)
			log.Println(err)
		}(f)
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) //这里忽略了异常处理
			ch <- struct{}{}
		}(f) //传递一个拷贝至goroutine函数,以免内部函数执行时,f的引用发生变化
	}

	for range filenames { //协程通知主协程执行完毕
		<-ch
	}
}

//goroutine泄露的一个实例
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err
			//这个函数这样返回是有问题的,
			// 这里返回后,其他未完成的产生缩略图的协程如果产生错误,
			// 向errors信道传递消息时,会阻塞,导致这些协程被一直挂起,无法被释放
			//这种bug被成为goroutine泄露
		}
	}

	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
			//这里return不会造成goroutine泄露
			//因为channel有足够大的缓冲区
			//保证所有的goroutine不会挂起
			//函数运行完毕后
			//ch的作用域失效
			//会被gc回收
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes:=make(chan int64)
	var wg sync.WaitGroup
	for f:=range filenames{
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb,err:=thumbnail.ImageFile(f)
			if err!=nil{
				log.Println(err)
				return
			}
			info,_:=os.Stat(thumb)//获取生成的缩略图文件的信息
			sizes<-info.Size()//通知缩略图文件大小
		}(f)
	}

	go func() {
		wg.Wait()//直到与wg.Add等量的wg.Done被调用前,wg.Wait都不会通过
		close(sizes)
	}()

	var total int64
	for size:=range sizes{
		total+=size
	}

	return total
}
