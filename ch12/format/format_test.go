package format

import "testing"

func TestFormatAtom(t *testing.T){
	tests:=[]struct{
		input interface{}
		want string
	}{
		{"ok","\"ok\""},
		{1,"1"},
		{3.1415926,"3.14"},
		{user{1, "jzh"},"format.user value"},
		{[]string{"风","花","雪","月"},"??"},
		{complex(1.0,2.5),"(1.00,2.50)"},
	}
	for _,test:=range tests{
		got:=Any(test.input)
		if got!=test.want{
			t.Errorf("Any(%q) got %s, want %s",test.input,got,test.want)
		}
	}
}

type user struct {
	id int
	name string
}

