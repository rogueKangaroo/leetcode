package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_entityParser(t *testing.T) {
	text := "and I quote: &quot;...&quot;"
	got := entityParser(text)
	fmt.Println(fmt.Sprintf("entityParser result is %v", got))
	time.Sleep(1 * time.Second)
}
