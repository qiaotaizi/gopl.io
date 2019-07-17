package main

import "fmt"

func main() {
	strings:=[]string{0:"ok",3:"hello",9:"world"}

	strings=nonempty(strings)

	fmt.Println(strings)
}


func nonempty(strings []string)[]string{
	i:=0
	for _,s:=range strings{
		if s!=""{
			strings[i]=s
			i++
		}
	}

	return strings[:i]
}
