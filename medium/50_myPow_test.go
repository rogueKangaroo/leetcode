package medium

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

func Test_myPow(t *testing.T) {
	x := float64(2)
	n := 10
	got := myPow(x, n)
	fmt.Println(fmt.Sprintf("myPow result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

func Test_myPow1(t *testing.T) {
	x := float64(2.1)
	n := 3
	got := myPow(x, n)
	fmt.Println(fmt.Sprintf("myPow result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

func Test_myPow2(t *testing.T) {
	x := float64(2)
	n := -2
	got := myPow(x, n)
	fmt.Println(fmt.Sprintf("myPow result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

// 0.00001
//2147483647

func Test_myPow3(t *testing.T) {
	x := float64(0.00001)
	n := 2147483647
	got := myPow(x, n)
	fmt.Println(fmt.Sprintf("myPow result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}