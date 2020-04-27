package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_combinationSum(t *testing.T) {
	candidates := []int{2,3,6,7}
	got := combinationSum(candidates, 7)
	fmt.Println(fmt.Sprintf("combinationSum result is %v", got))
	time.Sleep(1 * time.Second)
}


func Test_slice(t *testing.T) {
	candidates := []int{2,3,6,7}
	candidates = candidates[:len(candidates)-1]
	fmt.Println(fmt.Sprintf("test result is %v", candidates))
	time.Sleep(1 * time.Second)
}


func Test_combinationSum1(t *testing.T) {
	candidates := []int{2,3,5}
	got := combinationSum(candidates, 8)
	fmt.Println(fmt.Sprintf("combinationSum result is %v", got))
	time.Sleep(1 * time.Second)
}