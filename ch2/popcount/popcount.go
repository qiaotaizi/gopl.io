package popcount

//golang的包初始化顺序:
//1.包级变量
//2.init方法初始化

var pc [256]byte

func init(){
	for i:=range pc{
		pc[i]=pc[i/2]+byte(i&1)
	}
}

// 价差一个整数在二进制表示下有多少个1

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
