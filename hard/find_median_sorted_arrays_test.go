package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_findMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	got := findMedianSortedArrays(nums1, nums2)
	fmt.Println(fmt.Sprintf("findMedianSortedArrays result is %v", got))
	time.Sleep(1 * time.Second)
}
