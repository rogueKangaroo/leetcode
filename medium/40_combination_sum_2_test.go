package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_combinationSum2(t *testing.T) {
	candidates := []int{10,1,2,7,6,1,5}
	got := combinationSum2(candidates, 8)
	fmt.Println(fmt.Sprintf("combinationSum2 result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_combinationSum21(t *testing.T) {
	candidates := []int{2,5,2,1,2}
	got := combinationSum2(candidates, 5)
	fmt.Println(fmt.Sprintf("combinationSum2 result is %v", got))
	time.Sleep(1 * time.Second)
}