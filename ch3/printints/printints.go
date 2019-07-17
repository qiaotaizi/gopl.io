package printints

import (
	"bytes"
	"fmt"
)

func IntsToString(values []int) string {

	var buf bytes.Buffer
	buf.WriteByte('[')//如果确切知道一个字符属于ASCII码,使用WriteByte最好,否则可以用WriteRune函数来写入一个字符
	for i,v := range values{
		if i>0{
			buf.WriteByte(',')
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf,"%d",v)
	}
	buf.WriteByte(']')
	return buf.String()

}
