package simple

import (
	"fmt"
	"testing"
	"time"
)

//输入: [1,3,5,6], 5
//输出: 2
func Test_searchInsert(t *testing.T) {
	nums := []int{1,3,5,6}
	target := 5
	got := searchInsert(nums, target)
	fmt.Println(fmt.Sprintf("searchInsert result is %v", got))
	time.Sleep(1 * time.Second)
}


//输入: [1,3,5,6], 2
//输出: 1
func Test_searchInsert1(t *testing.T) {
	nums := []int{1,3,5,6}
	target := 2
	got := searchInsert(nums, target)
	fmt.Println(fmt.Sprintf("searchInsert result is %v", got))
	time.Sleep(1 * time.Second)
}


//输入: [1,3,5,6], 7
//输出: 4
func Test_searchInsert2(t *testing.T) {
	nums := []int{1,3,5,6}
	target := 7
	got := searchInsert(nums, target)
	fmt.Println(fmt.Sprintf("searchInsert result is %v", got))
	time.Sleep(1 * time.Second)
}


//输入: [1,3,5,6], 0
//输出: 0
func Test_searchInsert3(t *testing.T) {
	nums := []int{1,3,5,6}
	target := 0
	got := searchInsert(nums, target)
	fmt.Println(fmt.Sprintf("searchInsert result is %v", got))
	time.Sleep(1 * time.Second)
}