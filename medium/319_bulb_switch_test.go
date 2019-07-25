package medium

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_bulbSwitch(t *testing.T) {
	start := time.Now()
	for i:=0;i<1000;i++ {
		got := bulbSwitch(i)
		fmt.Println(fmt.Sprintf("bulbSwitch result is %v", got))
		fmt.Println(fmt.Sprintf("bulbSwitch time is %v",  time.Since(start)))

	}
	time.Sleep(1 * time.Second)
}


func Test_bulbSwitchTest(t *testing.T) {
	start := time.Now()
	for i:=0;i<99999;i++ {
		sum := 0
		for j:=0;j<100000;j++ {
			sum = sum + 1
		}
	}
	fmt.Println(fmt.Sprintf("bulbSwitch time is %v",  time.Since(start)))


	start1 := time.Now()
	wc := &sync.WaitGroup{}
	for i:=0;i<99999;i++ {
		wc.Add(1)
		go func(wc *sync.WaitGroup) {
			wc.Done()
			sum1 := 0
			for j:=0;j<100000;j++ {
				sum1 = sum1 + 1
			}
		}(wc)
	}
	wc.Wait()
	fmt.Println(fmt.Sprintf("bulbSwitch time is %v",  time.Since(start1)))
}

