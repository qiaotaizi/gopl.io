package word

import "testing"

//运行测试命令:
//go test $GOPATH/src/gopl.io/ch11/word/

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}

	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T){
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T){
	if !IsPalindrome("été"){
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T){
	input:="A man, a plan, a canal: Panama"
	if !IsPalindrome(input){
		t.Errorf(`IsPalindrome("%q") = false`,input)
	}
}

func TestIsPalindromeOpt(t *testing.T) {
	var tests=[]struct{
		input string
		want bool
	}{
		{"",true},
		{"a",true},
		{"aa",true},
		{"ab",false},
		{"kayak",true},
		{"detartrated",true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},// semi-palindrome
		{"北京输油管油输京北",true},
	}

	for _,test:=range tests{
		if got:=IsPalindromeOpt(test.input);got!=test.want{
			t.Errorf(`IsPalindromeOpt(%q) = %v`,test.input,got)
		}
	}
}
