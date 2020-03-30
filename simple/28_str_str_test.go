package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_strStr(t *testing.T) {
	haystack := "BBC ABCDAB ABCDABCDABDE"
	needle := "ABCDABD"
	got := strStr(haystack,needle)
	fmt.Println(fmt.Sprintf("strStr result is %v", got))
	time.Sleep(1 * time.Second)

	haystack = "hello"
	needle = "ll"
	got = strStr(haystack,needle)
	fmt.Println(fmt.Sprintf("strStr result is %v", got))
	time.Sleep(1 * time.Second)

	haystack = "aaaaa"
	needle = "bba"
	got = strStr(haystack,needle)
	fmt.Println(fmt.Sprintf("strStr result is %v", got))
	time.Sleep(1 * time.Second)

	haystack = "aabaaabaaac"
	needle = "aabaaac"
	got = strStr(haystack,needle)
	fmt.Println(fmt.Sprintf("strStr result is %v", got))
	time.Sleep(1 * time.Second)

	haystack = "ababaabbbbababbaabaaabaabbaaaabbabaabbbbbbabbaabbabbbabbbbbaaabaababbbaabbbabbbaabbbbaaabbababbabbbabaaabbaabbabababbbaaaaaaababbabaababaabbbbaaabbbabb"
	needle = "abbabbbabaa"
	fmt.Println(haystack[89:])
	got = strStr(haystack,needle)
	fmt.Println(fmt.Sprintf("strStr result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_strStr1(t *testing.T)  {
	haystack := "a"
	needle := "a"
	got := strStr(haystack,needle)
	fmt.Println(fmt.Sprintf("strStr result is %v", got))
	time.Sleep(1 * time.Second)
}
