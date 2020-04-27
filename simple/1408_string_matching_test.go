package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_stringMatching(t *testing.T) {
	words := []string{"leetcode","et","code"}
	got := stringMatching(words)
	fmt.Println(fmt.Sprintf("stringMatching result is %v", got))
	time.Sleep(1 * time.Second)
}
