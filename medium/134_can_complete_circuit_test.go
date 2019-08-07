package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_canCompleteCircuit(t *testing.T) {
	gas := []int{1,2,3,4,5}
	cost := []int{3,4,5,1,2}
	got := canCompleteCircuit(gas, cost)
	fmt.Println(fmt.Sprintf("canCompleteCircuit result is %v", got))
	time.Sleep(1 * time.Second)

	gas = []int{2,3,4}
	cost = []int{3,4,3}
	got = canCompleteCircuit(gas, cost)
	fmt.Println(fmt.Sprintf("canCompleteCircuit result is %v", got))
	time.Sleep(1 * time.Second)

	gas = []int{2}
	cost = []int{2}
	got = canCompleteCircuit(gas, cost)
	fmt.Println(fmt.Sprintf("canCompleteCircuit result is %v", got))
	time.Sleep(1 * time.Second)
}
