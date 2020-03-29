package simple

import (
	"fmt"
	"testing"
	"time"
)

// III ，输出3
func Test_romanToInt(t *testing.T) {
	s := "III"
	got := romanToInt(s)
	fmt.Println(fmt.Sprintf("romanToInt result is %v", got))
	time.Sleep(1 * time.Second)
}


// IV ，输出4
func Test_romanToInt1(t *testing.T) {
	s := "IV"
	got := romanToInt(s)
	fmt.Println(fmt.Sprintf("romanToInt result is %v", got))
	time.Sleep(1 * time.Second)
}

// IX ，输出9
func Test_romanToInt2(t *testing.T) {
	s := "IX"
	got := romanToInt(s)
	fmt.Println(fmt.Sprintf("romanToInt result is %v", got))
	time.Sleep(1 * time.Second)
}

// LVIII ，输出58
func Test_romanToInt3(t *testing.T) {
	s := "LVIII"
	got := romanToInt(s)
	fmt.Println(fmt.Sprintf("romanToInt result is %v", got))
	time.Sleep(1 * time.Second)
}

// MCMXCIV ，输出1994
func Test_romanToInt4(t *testing.T) {
	s := "MCMXCIV"
	got := romanToInt(s)
	fmt.Println(fmt.Sprintf("romanToInt result is %v", got))
	time.Sleep(1 * time.Second)
}

// IV ，输出4
func Test_romanToInt5(t *testing.T) {
	s := "IV"
	got := romanToInt(s)
	fmt.Println(fmt.Sprintf("romanToInt result is %v", got))
	time.Sleep(1 * time.Second)
}