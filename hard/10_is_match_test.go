package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_isMatch(t *testing.T) {
	s := "aa"
	p := "a"
	got := isMatch(s, p)
	if got != false {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[false] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "aa"
	p = "a*"
	got = isMatch(s, p)
	if got != true {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[true] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "ab"
	p = ".*"
	got = isMatch(s, p)
	if got != true {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[true] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "aab"
	p = "a*b"
	got = isMatch(s, p)
	if got != true {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[true] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "mississippi"
	p = "mis*is*p*."
	got = isMatch(s, p)
	if got != false {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[false] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "ssippi"
	p = "s*p*."
	got = isMatch(s, p)
	if got != false {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[false] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "aaa"
	p = "aaaa"
	got = isMatch(s, p)
	if got != false {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[false] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "aaa"
	p = "a*a"
	got = isMatch(s, p)
	if got != true {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[true] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "aaa"
	p = "ab*a"
	got = isMatch(s, p)
	if got != false {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[false] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "aab"
	p = "c*a*b"
	got = isMatch(s, p)
	if got != true {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[true] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "aaa"
	p = "ab*a*c*a"
	got = isMatch(s, p)
	if got != true {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[true] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "a"
	p = "a*"
	got = isMatch(s, p)
	if got != true {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[true] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "aaa"
	p = ".*"
	got = isMatch(s, p)
	if got != true {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[true] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "bbbba"
	p = ".*a*a"
	got = isMatch(s, p)
	if got != true {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[true] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "ppi"
	p = "p*."
	got = isMatch(s, p)
	if got != true {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch result[true] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)
}
