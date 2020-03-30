package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_countAndSay(t *testing.T) {
	n:= 1
	got := countAndSay(n)
	fmt.Println(fmt.Sprintf("countAndSay result is %v", got))
	time.Sleep(1 * time.Second)
}


// 输入4，输出1211
func Test_countAndSay1(t *testing.T) {
	n:= 4
	got := countAndSay(n)
	fmt.Println(fmt.Sprintf("countAndSay 4 result is %v", got))
	time.Sleep(1 * time.Second)
}

// 输入5，输出111221
func Test_countAndSay2(t *testing.T) {
	n:= 5
	got := countAndSay(n)
	fmt.Println(fmt.Sprintf("countAndSay result is %v", got))
	time.Sleep(1 * time.Second)
}

// 输入6，输出312211
func Test_countAndSay3(t *testing.T) {
	n:= 6
	got := countAndSay(n)
	fmt.Println(fmt.Sprintf("countAndSay result is %v", got))
	time.Sleep(1 * time.Second)
}

// 输入6，输出312211
func Test_countAndSayAll(t *testing.T) {
	n := 1
	got := countAndSay(n)
	fmt.Println(fmt.Sprintf("countAndSay 1 result is %v", got))

	n = 2
	got = countAndSay(n)
	fmt.Println(fmt.Sprintf("countAndSay 2 result is %v", got))


	n = 3
	got = countAndSay(n)
	fmt.Println(fmt.Sprintf("countAndSay 3 result is %v", got))


	n = 4
	got = countAndSay(n)
	fmt.Println(fmt.Sprintf("countAndSay 4 result is %v", got))


	n = 5
	got = countAndSay(n)
	fmt.Println(fmt.Sprintf("countAndSay 5 result is %v", got))


	n = 6
	got = countAndSay(n)
	fmt.Println(fmt.Sprintf("countAndSay 6 result is %v", got))
	time.Sleep(1 * time.Second)
}