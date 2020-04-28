package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_groupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	got := groupAnagrams(strs)
	fmt.Println(fmt.Sprintf("addTwoNumbers result is %v", got))
	time.Sleep(1 * time.Second)
}
