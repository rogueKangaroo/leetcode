package simple

import (
	"fmt"
	"testing"
	"time"
)

func TestIsValid(t *testing.T) {
	s := "{}()[]"
	got := IsValid(s)
	fmt.Println(fmt.Sprintf("IsValid result is %v", got))
	time.Sleep(1 * time.Second)

	s = "{()}[]"
	got = IsValid(s)
	fmt.Println(fmt.Sprintf("IsValid result is %v", got))
	time.Sleep(1 * time.Second)

	s = "{(})[]"
	got = IsValid(s)
	fmt.Println(fmt.Sprintf("IsValid result is %v", got))
	time.Sleep(1 * time.Second)
}
