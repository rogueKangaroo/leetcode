package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_search(t *testing.T) {
	nums := []int{4,5,6,7,0,1,2}
	got := search(nums, 0)
	fmt.Println(fmt.Sprintf("search result is %v", got))
	time.Sleep(1 * time.Second)
}


func Test_search1(t *testing.T) {
	nums := []int{4,5,6,7,0,1,2}
	got := search(nums, 3)
	fmt.Println(fmt.Sprintf("search result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_search2(t *testing.T) {
	nums := []int{1,3}
	got := search(nums, 0)
	fmt.Println(fmt.Sprintf("search result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_search3(t *testing.T) {
	nums := []int{1,3}
	got := search(nums, 3)
	fmt.Println(fmt.Sprintf("search result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_search4(t *testing.T) {
	nums := []int{1,3}
	got := search(nums, 4)
	fmt.Println(fmt.Sprintf("search result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_search5(t *testing.T) {
	nums := []int{1,3,5}
	got := search(nums, 1)
	fmt.Println(fmt.Sprintf("search result is %v", got))
	time.Sleep(1 * time.Second)
}