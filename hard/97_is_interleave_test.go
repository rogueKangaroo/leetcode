package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_isInterleave(t *testing.T) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	got := isInterleave(s1, s2, s3)
	fmt.Println(fmt.Sprintf("isInterleave result(true) is %v", got))
	time.Sleep(100 * time.Millisecond)

	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbbaccc"
	got = isInterleave(s1, s2, s3)
	fmt.Println(fmt.Sprintf("isInterleave result(false) is %v", got))
	time.Sleep(100 * time.Millisecond)

	s1 = "aa"
	s2 = "ab"
	s3 = "aaba"
	got = isInterleave(s1, s2, s3)
	fmt.Println(fmt.Sprintf("isInterleave result(true) is %v", got))
	time.Sleep(100 * time.Millisecond)
}
