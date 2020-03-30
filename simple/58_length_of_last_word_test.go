package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_lengthOfLastWord(t *testing.T) {
	s := "Hello World"
	got := lengthOfLastWord(s)
	fmt.Println(fmt.Sprintf("lengthOfLastWord result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_lengthOfLastWord1(t *testing.T) {
	s := " q"
	got := lengthOfLastWord(s)
	fmt.Println(fmt.Sprintf("lengthOfLastWord result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_lengthOfLastWord2(t *testing.T) {
	s := "b a "
	got := lengthOfLastWord(s)
	fmt.Println(fmt.Sprintf("lengthOfLastWord result is %v", got))
	time.Sleep(1 * time.Second)
}