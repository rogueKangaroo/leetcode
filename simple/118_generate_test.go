package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_generate(t *testing.T) {
	got := generate(5)
	fmt.Println(fmt.Sprintf("generate result is %v", got))
	time.Sleep(100 * time.Millisecond)
}
