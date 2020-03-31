package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_mySqrt(t *testing.T) {
	x := 9
	got := mySqrt(x)
	fmt.Println(fmt.Sprintf("mySqrt result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_mySqrt1(t *testing.T) {
	x := 6
	got := mySqrt(x)
	fmt.Println(fmt.Sprintf("mySqrt result is %v", got))
	time.Sleep(1 * time.Second)
}


func Test_mySqrt2(t *testing.T) {
	x := 10
	got := mySqrt(x)
	fmt.Println(fmt.Sprintf("mySqrt result is %v", got))
	time.Sleep(1 * time.Second)
}