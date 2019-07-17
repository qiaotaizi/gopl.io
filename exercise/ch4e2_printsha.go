package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var a = flag.Bool("sha512",false,"默认以SHA256对参数进行摘要,若有此参数,则使用SHA512进行摘要")

func main() {
	flag.Parse()
	if *a{
		if len(os.Args)>2{
			fmt.Printf("%x\n",sha512.Sum512([]byte(os.Args[2])))
		}else{
			fmt.Println("请输入要加密的参数")
		}
	}else{
		if len(os.Args)>1{
			fmt.Printf("%x\n",sha256.Sum256([]byte(os.Args[1])))
		}else{
			fmt.Println("请输入要加密的参数")
		}
	}
}
