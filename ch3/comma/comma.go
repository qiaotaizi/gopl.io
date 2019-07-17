package comma

//给出一个整数,每三位加一个逗号
func Comma(s string) string{
	n:=len(s)
	if n<=3{
		return s
	}
	return Comma(s[:n-3])+","+s[n-3:]
}

