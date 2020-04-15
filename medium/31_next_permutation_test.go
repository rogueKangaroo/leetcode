package medium

import (
	"fmt"
	"testing"
	"time"
)
// 1,2,3 → 1,3,2
//3,2,1 → 1,2,3
//1,1,5 → 1,5,1
//1,1,5,4,6,5 -> 1,1,5,5,4,6
func Test_nextPermutation(t *testing.T) {
	nums := []int{1,1,5,4,6,5}
	nextPermutation(nums)
	fmt.Println(fmt.Sprintf("nextPermutation result is %v", nums))
	time.Sleep(1 * time.Second)
}

func Test_nextPermutation1(t *testing.T) {
	nums := []int{1,2,3}
	nextPermutation(nums)
	fmt.Println(fmt.Sprintf("nextPermutation result is %v", nums))
	time.Sleep(1 * time.Second)
}

//1,3,5,4,3,2 -> 1,4,2,3,3,5
func Test_nextPermutation2(t *testing.T) {
	nums := []int{1,3,5,4,3,2}
	nextPermutation(nums)
	fmt.Println(fmt.Sprintf("nextPermutation result is %v", nums))
	time.Sleep(1 * time.Second)
}

func Test_nextPermutation3(t *testing.T) {
	nums := []int{1,3,2}
	nextPermutation(nums)
	fmt.Println(fmt.Sprintf("nextPermutation result is %v", nums))
	time.Sleep(1 * time.Second)
}