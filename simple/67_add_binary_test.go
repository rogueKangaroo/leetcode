package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_addBinary(t *testing.T) {
	a := "11"
	b := "1"
	got := addBinary(a, b)
	fmt.Println(fmt.Sprintf("addBinary result is %v", got))
	time.Sleep(1 * time.Second)
}

// 10101
func Test_addBinary1(t *testing.T) {
	a := "1010"
	b := "1011"
	got := addBinary(a, b)
	fmt.Println(fmt.Sprintf("addBinary result is %v", got))
	time.Sleep(1 * time.Second)
}

// 10
func Test_addBinary2(t *testing.T) {
	a := "1111"
	b := "1111"
	got := addBinary(a, b)
	fmt.Println(fmt.Sprintf("addBinary result is %v", got))
	time.Sleep(1 * time.Second)
}