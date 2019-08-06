package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_maxProfit(t *testing.T) {
	prices := []int{7,1,5,3,6,4}
	got := maxProfit(prices)
	fmt.Println(fmt.Sprintf("maxProfit result is %v", got))
	time.Sleep(1 * time.Second)
}
