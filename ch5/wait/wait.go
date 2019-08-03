package wait

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//尝试在1分钟之内打通一个url
func WaitForServer(url string) error{
	const timeout  = 1*time.Minute
	deadline:=time.Now().Add(timeout)
	for tries:=0;time.Now().Before(deadline);tries++{
		_,err:=http.Head(url)
		if err==nil{
			return nil;
		}
		log.Printf("server not responding (%s); retrying...",err)
		time.Sleep(time.Second<<uint(tries))//每次循环将等待更长的时间
	}
	return fmt.Errorf("server %s failed to response after %s",url,deadline)
}
