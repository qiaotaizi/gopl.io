package basename

//不使用任何库函数,从一个完整文件路径下取得一个文件名
func BaseName1(s string) string{

	//找到最后一个'/'并向后取子串
	for i:=len(s)-1;i>=0;i--{
		if s[i]=='/'{
			s=s[i+1:];
			break
		}
	}

	//找到最后一个'.'并向前取子串
	for i:=len(s)-1;i>=0;i--{
		if s[i]=='.'{
			s=s[:i]
			break
		}
	}

	return s
}
