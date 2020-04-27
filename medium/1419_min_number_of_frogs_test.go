package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_minNumberOfFrogs(t *testing.T) {
	croakOfFrogs := "crocakcroraoakk"
	got := minNumberOfFrogs(croakOfFrogs)
	fmt.Println(fmt.Sprintf("minNumberOfFrogs result is %v", got))
	time.Sleep(1 * time.Second)
}
