package simple

import (
	"fmt"
	"strconv"
	"testing"
)

// 1,1,2 ，输出2
func Test_removeDuplicates(t *testing.T) {
	nums := []int{1,1,2}
	got := removeDuplicates(nums)
	fmt.Println(fmt.Sprintf("removeDuplicates result is %v", got))
	for i:=0;i<got;i++ {
		fmt.Print(strconv.Itoa(nums[i]))
	}
}


// 0,0,1,1,1,2,2,3,3,4 ，输出5
func Test_removeDuplicates1(t *testing.T) {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	got := removeDuplicates(nums)
	fmt.Println(fmt.Sprintf("removeDuplicates result is %v", got))
	for i:=0;i<got;i++ {
		fmt.Print(strconv.Itoa(nums[i]))
	}
}