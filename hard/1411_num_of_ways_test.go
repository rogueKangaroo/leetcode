package hard

import (
	"fmt"
	"testing"
	"time"
)
//30228214
func Test_numOfWays(t *testing.T) {
	got := numOfWays(5000)
	fmt.Println(fmt.Sprintf("numOfWays result is %v", got))
	time.Sleep(1 * time.Second)
}
