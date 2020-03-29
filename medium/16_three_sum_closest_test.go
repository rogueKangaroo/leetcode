package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_threeSumClosest(t *testing.T) {
	nums := []int{-1,2,1,-4}
	target := 1
	got := threeSumClosest(nums, target)
	fmt.Println(fmt.Sprintf("threeSumClosest result is %v", got))
	time.Sleep(1 * time.Second)
}


func Test_threeSumClosest1(t *testing.T) {
	nums := []int{0,1,2}
	target := 0
	got := threeSumClosest(nums, target)
	fmt.Println(fmt.Sprintf("threeSumClosest result is %v", got))
	time.Sleep(1 * time.Second)
}


func Test_threeSumClosest2(t *testing.T) {
	nums := []int{-3,0,1,2}
	target := 0
	got := threeSumClosest(nums, target)
	fmt.Println(fmt.Sprintf("threeSumClosest result is %v", got))
	time.Sleep(1 * time.Second)
}

//[-3,-2,-5,3,-4]
//-1
func Test_threeSumClosest3(t *testing.T) {
	nums := []int{-5,-4,-3,-2,3}
	target := -1
	got := threeSumClosest(nums, target)
	fmt.Println(fmt.Sprintf("threeSumClosest result is %v", got))
	time.Sleep(1 * time.Second)
}

//[1,2,5,10,11]
//12
func Test_threeSumClosest4(t *testing.T) {
	nums := []int{1,2,5,10,11}
	target := 12
	got := threeSumClosest(nums, target)
	fmt.Println(fmt.Sprintf("threeSumClosest result is %v", got))
	time.Sleep(1 * time.Second)
}