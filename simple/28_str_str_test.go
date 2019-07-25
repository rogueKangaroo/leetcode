package simple

import (
	"fmt"
	"testing"
	"time"
)

func TestGetStrNumArray(t *testing.T) {
	needle := "ABCDABD"
	got := GetStrNumArray(needle)
	fmt.Println(fmt.Sprintf("GetStrNumArray result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_strStr(t *testing.T) {
	haystack := "BBC ABCDAB ABCDABCDABDE"
	needle := "ABCDABD"
	got := strStr(haystack,needle)
	fmt.Println(fmt.Sprintf("strStr result is %v", got))
	time.Sleep(1 * time.Second)

	haystack = "hello"
	needle = "ll"
	fmt.Println(haystack[1:])
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
