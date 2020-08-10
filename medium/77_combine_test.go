package medium

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

func Test_combine(t *testing.T) {
	n := 4
	k := 2
	got := combine(n, k)
	fmt.Println(fmt.Sprintf("combine result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}


func Test_combine1(t *testing.T) {
	n := 4
	k := 3
	got := combine(n, k)
	fmt.Println(fmt.Sprintf("combine result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

func Test_combine2(t *testing.T) {
	n := 4
	k := 1
	got := combine(n, k)
	fmt.Println(fmt.Sprintf("combine result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

func Test_combine3(t *testing.T) {
	n := 5
	k := 3
	got := combine(n, k)
	fmt.Println(fmt.Sprintf("combine result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}