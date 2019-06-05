package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_numTrees(t *testing.T) {
	n := 3
	got := numTrees(n)
	fmt.Println(fmt.Sprintf("numTrees result is %v", got))
	time.Sleep(1 * time.Second)

	n = 4
	got = numTrees(n)
	fmt.Println(fmt.Sprintf("numTrees result is %v", got))
	time.Sleep(1 * time.Second)
}
