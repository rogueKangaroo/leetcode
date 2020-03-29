package medium

import (
	"fmt"
	"testing"
	"time"
)

// 输入3，输出"III"
func Test_intToRoman(t *testing.T) {
	num := 3
	got := intToRoman(num)
	fmt.Println(fmt.Sprintf("intToRoman result is %v", got))
	time.Sleep(1 * time.Second)
}


// 输入: 4
// 输出: "IV"
func Test_intToRoman1(t *testing.T) {
	num := 4
	got := intToRoman(num)
	fmt.Println(fmt.Sprintf("intToRoman result is %v", got))
	time.Sleep(1 * time.Second)
}


//输入: 9
//输出: "IX"
func Test_intToRoman2(t *testing.T) {
	num := 9
	got := intToRoman(num)
	fmt.Println(fmt.Sprintf("intToRoman result is %v", got))
	time.Sleep(1 * time.Second)
}


//输入: 58
//输出: "LVIII"
func Test_intToRoman3(t *testing.T) {
	num := 58
	got := intToRoman(num)
	fmt.Println(fmt.Sprintf("intToRoman result is %v", got))
	time.Sleep(1 * time.Second)
}


//输入: 1994
//输出: "MCMXCIV"
func Test_intToRoman4(t *testing.T) {
	num := 1994
	got := intToRoman(num)
	fmt.Println(fmt.Sprintf("intToRoman result is %v", got))
	time.Sleep(1 * time.Second)
}

