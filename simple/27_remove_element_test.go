package simple

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_removeElement(t *testing.T) {
	nums := []int{3,2,2,3}
	val := 3
	got := removeElement(nums, val)
	fmt.Println(fmt.Sprintf("removeElement result is %v", got))
	for i:=0;i<got;i++ {
		fmt.Print(strconv.Itoa(nums[i]))
	}
}


func Test_removeElement1(t *testing.T) {
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	got := removeElement(nums, val)
	fmt.Println(fmt.Sprintf("removeElement result is %v", got))
	for i:=0;i<got;i++ {
		fmt.Print(strconv.Itoa(nums[i]))
	}
}