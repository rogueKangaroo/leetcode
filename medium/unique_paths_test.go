package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_uniquePaths(t *testing.T) {
	m := 4
	n := 4
	got := uniquePaths(m, n)
	fmt.Println(fmt.Sprintf("uniquePaths result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_uniquePaths2(t *testing.T) {
	m := 4
	n := 4
	got := uniquePaths2(m, n)
	fmt.Println(fmt.Sprintf("uniquePaths result is %v", got))
	time.Sleep(1 * time.Second)
}
