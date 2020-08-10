package medium

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

/*
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"
*/
func Test_getPermutation(t *testing.T) {
	n := 3
	k := 3
	got := getPermutation(n, k)
	fmt.Println(fmt.Sprintf("getPermutation result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

func Test_getPermutation1(t *testing.T) {
	n := 4
	k := 9
	got := getPermutation(n, k)
	fmt.Println(fmt.Sprintf("getPermutation result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

func Test_getPermutation2(t *testing.T) {
	n := 2
	k := 1
	got := getPermutation(n, k)
	fmt.Println(fmt.Sprintf("getPermutation result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

func Test_getPermutation3(t *testing.T) {
	n := 3
	k := 2
	got := getPermutation(n, k)
	fmt.Println(fmt.Sprintf("getPermutation result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}