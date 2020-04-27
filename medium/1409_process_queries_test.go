package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_processQueries(t *testing.T) {
	queries := []int{7,5,5,8,3}
	got := processQueries(queries,8)
	fmt.Println(fmt.Sprintf("processQueries result is %v", got))
	time.Sleep(1 * time.Second)
}


func Test_processQueries1(t *testing.T) {
	queries := []int{3,1,2,1}
	got := processQueries(queries,5)
	fmt.Println(fmt.Sprintf("processQueries result is %v", got))
	time.Sleep(1 * time.Second)
}