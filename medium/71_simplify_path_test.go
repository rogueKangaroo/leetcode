package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_simplifyPath(t *testing.T) {
	path := "/home/"
	got := simplifyPath(path)
	fmt.Println(fmt.Sprintf("simplifyPath path %v result is %v", path,got))
	time.Sleep(1 * time.Second)

	path = "/../"
	got = simplifyPath(path)
	fmt.Println(fmt.Sprintf("simplifyPath path %v result is %v", path,got))
	time.Sleep(1 * time.Second)

	path = "/home//foo/"
	got = simplifyPath(path)
	fmt.Println(fmt.Sprintf("simplifyPath path %v result is %v", path,got))
	time.Sleep(1 * time.Second)

	path = "/a/./b/../../c/"
	got = simplifyPath(path)
	fmt.Println(fmt.Sprintf("simplifyPath path %v result is %v", path,got))
	time.Sleep(1 * time.Second)

	path = "/a/../../b/../c//.//"
	got = simplifyPath(path)
	fmt.Println(fmt.Sprintf("simplifyPath path %v result is %v", path,got))
	time.Sleep(1 * time.Second)

	path = "/a//b////c/d//././/.."
	got = simplifyPath(path)
	fmt.Println(fmt.Sprintf("simplifyPath path %v result is %v", path,got))
	time.Sleep(1 * time.Second)
}
