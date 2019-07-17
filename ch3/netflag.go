package main

import "fmt"

//iota常量生成器的使用

type Flags uint

const (
	FlagUp Flags=1<<iota
	FlagBroadCast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool {
	return v&FlagUp==FlagUp
}

func TurnDown(v *Flags){
	*v &^=FlagUp//(&^是按位置零操作,可在gopl中查看对于&^的解释)
}

func SetBroadcast(v *Flags){
	*v |= FlagBroadCast
}

func IsCast(v Flags)bool{
	return v&(FlagBroadCast|FlagMulticast) != 0
}

func main() {
	v := FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}
