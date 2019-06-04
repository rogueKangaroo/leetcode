package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_isScramble(t *testing.T) {
	s1 := "great"
	s2 := "grtae"
	got := isScramble(s1, s2)
	fmt.Println(fmt.Sprintf("isScramble result is %v", got))
	time.Sleep(100 * time.Millisecond)

	s1 = "abcde"
	s2 = "caebd"
	got = isScramble(s1, s2)
	fmt.Println(fmt.Sprintf("isScramble result is %v", got))
	time.Sleep(100 * time.Millisecond)

	s1 = "abb"
	s2 = "bab"
	got = isScramble(s1, s2)
	fmt.Println(fmt.Sprintf("isScramble result is %v", got))
	time.Sleep(100 * time.Millisecond)

	s1 = "abb"
	s2 = "bba"
	got = isScramble(s1, s2)
	fmt.Println(fmt.Sprintf("isScramble result is %v", got))
	time.Sleep(100 * time.Millisecond)

	s1 = "a"
	s2 = "a"
	got = isScramble(s1, s2)
	fmt.Println(fmt.Sprintf("isScramble result is %v", got))
	time.Sleep(100 * time.Millisecond)

	s1 = "abc"
	s2 = "acb"
	got = isScramble(s1, s2)
	fmt.Println(fmt.Sprintf("isScramble result is %v", got))
	time.Sleep(100 * time.Millisecond)

	s1 = "abc"
	s2 = "bca"
	got = isScramble(s1, s2)
	fmt.Println(fmt.Sprintf("isScramble result is %v", got))
	time.Sleep(100 * time.Millisecond)

	s1 = "eat"
	s2 = "tae"
	got = isScramble(s1, s2)
	fmt.Println(fmt.Sprintf("isScramble result is %v", got))
	time.Sleep(100 * time.Millisecond)

}

