package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_longestValidParentheses(t *testing.T) {
	//s := "()()()()"
	//got := longestValidParentheses(s)
	//fmt.Println(fmt.Sprintf("longestValidParentheses result is %v", got))
	//time.Sleep(100 * time.Millisecond)
	//
	//s = "(()"
	//got = longestValidParentheses(s)
	//fmt.Println(fmt.Sprintf("longestValidParentheses result is %v", got))
	//time.Sleep(100 * time.Millisecond)
	//
	//s = ")()())"
	//got = longestValidParentheses(s)
	//fmt.Println(fmt.Sprintf("longestValidParentheses result is %v", got))
	//time.Sleep(100 * time.Millisecond)
	//
	//s = "()(()"
	//got = longestValidParentheses(s)
	//fmt.Println(fmt.Sprintf("longestValidParentheses result is %v", got))
	//time.Sleep(100 * time.Millisecond)

	s := "(()()"
	got := longestValidParentheses(s)
	fmt.Println(fmt.Sprintf("longestValidParentheses result is %v", got))
	time.Sleep(100 * time.Millisecond)
}
