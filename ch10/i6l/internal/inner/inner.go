package inner

import "fmt"

func innerFunc(){
	fmt.Println("innerFunc")
}

func InnerFunc(){
	innerFunc()
}