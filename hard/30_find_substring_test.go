package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_findSubstring(t *testing.T) {
	s := "barfoothefoobarman"
	words := []string{"foo","bar"}
	got := findSubstring(s, words)
	fmt.Println(fmt.Sprintf("findSubstring result is %v", got))
	time.Sleep(100 * time.Millisecond)
}


func Test_findSubstrin1(t *testing.T) {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word","good","best","word"}
	got := findSubstring(s, words)
	fmt.Println(fmt.Sprintf("findSubstring result is %v", got))
	time.Sleep(100 * time.Millisecond)
}

// 预期是8
func Test_findSubstrin2(t *testing.T) {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word","good","best","good"}
	got := findSubstring(s, words)
	fmt.Println(fmt.Sprintf("findSubstring result is %v", got))
	time.Sleep(100 * time.Millisecond)
}

//"lingmindraboofooowingdingbarrwingmonkeypoundcake"
//["fooo","barr","wing","ding","wing"]
// 预期13
func Test_findSubstrin3(t *testing.T) {
	s := "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	words := []string{"fooo","barr","wing","ding","wing"}
	got := findSubstring(s, words)
	fmt.Println(fmt.Sprintf("findSubstring result is %v", got))
	time.Sleep(100 * time.Millisecond)
}