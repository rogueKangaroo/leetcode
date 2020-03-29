package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_longestCommonPrefix(t *testing.T) {
	strs :=[]string{"flower","flow","flight"}
	got := longestCommonPrefix(strs)
	fmt.Println(fmt.Sprintf("longestCommonPrefix result is %v", got))
	time.Sleep(1 * time.Second)
}


func Test_longestCommonPrefix1(t *testing.T) {
	strs :=[]string{"dog","racecar","car"}
	got := longestCommonPrefix(strs)
	fmt.Println(fmt.Sprintf("longestCommonPrefix result is %v", got))
	time.Sleep(1 * time.Second)
}