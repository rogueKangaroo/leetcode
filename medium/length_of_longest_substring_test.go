package medium

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	s := "abccba"
	got := lengthOfLongestSubstring(s)
	fmt.Println(fmt.Sprintf("lengthOfLongestSubstring result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}
