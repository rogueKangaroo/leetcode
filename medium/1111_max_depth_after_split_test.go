package medium

import (
	"fmt"
	"testing"
	"time"
)

//示例 1：
//
//输入：seq = "(()())"
//输出：[0,1,1,1,1,0]
func Test_maxDepthAfterSplit(t *testing.T) {
	seq := "(()())"
	got := maxDepthAfterSplit(seq)
	fmt.Println(fmt.Sprintf("maxDepthAfterSplit result is %v", got))
	time.Sleep(1 * time.Second)
}

//示例 2：
//
//输入：seq = "()(())()"
//输出：[0,0,0,1,1,0,1,1]
// 
func Test_maxDepthAfterSplit1(t *testing.T) {
	seq := "()(())()"
	got := maxDepthAfterSplit(seq)
	fmt.Println(fmt.Sprintf("maxDepthAfterSplit result is %v", got))
	time.Sleep(1 * time.Second)
}