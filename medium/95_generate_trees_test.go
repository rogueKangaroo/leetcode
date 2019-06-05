package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_generateTrees(t *testing.T) {
	n := 3
	got := generateTrees(n)
	fmt.Println(fmt.Sprintf("generateTrees result is %v", got))
	time.Sleep(1 * time.Second)
}
