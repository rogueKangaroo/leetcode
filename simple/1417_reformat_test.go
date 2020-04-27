package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_reformat(t *testing.T) {
	s := "ab123"
	got := reformat(s)
	fmt.Println(fmt.Sprintf("reformat result is %v", got))
	time.Sleep(1 * time.Second)
}


func Test_reformat1(t *testing.T) {
	s := "a0b1c2"
	got := reformat(s)
	fmt.Println(fmt.Sprintf("reformat result is %v", got))
	time.Sleep(1 * time.Second)
}