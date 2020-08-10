package medium

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

func Test_generateMatrix(t *testing.T) {
	n:=3
	got := generateMatrix(n)
	for i:=0;i<n;i++ {
		fmt.Println(fmt.Sprintf("%v", render.Render(got[i])))
	}
	time.Sleep(1 * time.Second)
}

func Test_generateMatrix1(t *testing.T) {
	n:=4
	got := generateMatrix(n)
	for i:=0;i<n;i++ {
		fmt.Println(fmt.Sprintf("%v", render.Render(got[i])))
	}
	time.Sleep(1 * time.Second)
}

func Test_generateMatrix2(t *testing.T) {
	n:=5
	got := generateMatrix(n)
	for i:=0;i<n;i++ {
		fmt.Println(fmt.Sprintf("%v", render.Render(got[i])))
	}
	time.Sleep(1 * time.Second)
}