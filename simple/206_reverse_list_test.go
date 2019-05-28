package simple

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

func Test_reverseList(t *testing.T) {
	node13 := &ListNode{3, nil}
	node12 := &ListNode{4, node13}
	node11 := &ListNode{2, node12}
	got := reverseList(node11)
	fmt.Println(fmt.Sprintf("reverseList result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

func Test_reverseList2(t *testing.T) {
	node23 := &ListNode{4, nil}
	node22 := &ListNode{6, node23}
	node21 := &ListNode{5, node22}
	got := reverseList2(node21)
	fmt.Println(fmt.Sprintf("reverseList2 result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}
