package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_divide(t *testing.T) {
	dividend := 11
	divisor := 3
	got := divide(dividend, divisor)
	fmt.Println(fmt.Sprintf("divide result is %v", got))
	time.Sleep(1 * time.Second)
}


func Test_divide1(t *testing.T) {
	dividend := 7
	divisor := -2
	got := divide(dividend, divisor)
	fmt.Println(fmt.Sprintf("divide result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_divide2(t *testing.T) {
	dividend := 10000000
	divisor := -2
	got := divide(dividend, divisor)
	fmt.Println(fmt.Sprintf("divide result is %v", got))
	time.Sleep(1 * time.Second)
}


func Test_divide3(t *testing.T) {
	fmt.Println(fmt.Sprintf("divide result is %v", 2147483648/3))
	time.Sleep(1 * time.Second)
}

// -2147483648
//-1
func Test_divide4(t *testing.T) {
	dividend := -2147483648
	divisor := -1
	got := divide(dividend, divisor)
	fmt.Println(fmt.Sprintf("divide result is %v", got))
	time.Sleep(1 * time.Second)
}

// -2147483648
// 2
func Test_divide5(t *testing.T) {
	dividend := 10000
	divisor := 4
	got := divide(dividend, divisor)
	fmt.Println(fmt.Sprintf("divide result is %v", got))
	time.Sleep(1 * time.Second)
}