package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_plusOne(t *testing.T) {
	digits := []int{1,2,3}
	got := plusOne(digits)
	fmt.Println(fmt.Sprintf("plusOne result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_plusOne1(t *testing.T) {
	digits := []int{4,3,2,1}
	got := plusOne(digits)
	fmt.Println(fmt.Sprintf("plusOne result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_plusOne2(t *testing.T) {
	digits := []int{8,9,9,9}
	got := plusOne(digits)
	fmt.Println(fmt.Sprintf("plusOne result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_plusOne3(t *testing.T) {
	digits := []int{9,9,9,9}
	got := plusOne(digits)
	fmt.Println(fmt.Sprintf("plusOne result is %v", got))
	time.Sleep(1 * time.Second)
}