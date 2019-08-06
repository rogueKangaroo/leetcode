package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_handleSkyline(t *testing.T) {
	buildings := [][]int{{2, 10}, {9, 10}, {3, 15}, {7, 15}, {5, 12}, {12, 12}, {15, 10}, {20, 10}, {19, 8}, {24, 8}}
	temp := make([][]int, len(buildings))
	sortBuildings(&buildings, 0, len(buildings)-1, temp)
	fmt.Println(fmt.Sprintf("handleSkyline result is %v", buildings))
	time.Sleep(100 * time.Millisecond)
}

func Test_getSkyline(t *testing.T) {
	buildings := [][]int{{2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}}
	got := getSkyline(buildings)
	fmt.Println(fmt.Sprintf("getSkyline result is %v", got))
	time.Sleep(100 * time.Millisecond)
}

func Test_getSkyline1(t *testing.T) {
	buildings := [][]int{{3,7,8},{3,8,7},{3,9,6},{3,10,5},{3,11,4},{3,12,3},{3,13,2},{3,14,1}}
	got := getSkyline(buildings)
	fmt.Println(fmt.Sprintf("getSkyline result is %v", got))
	time.Sleep(100 * time.Millisecond)
}
