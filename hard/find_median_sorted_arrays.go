package hard

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		nums1 = nums2
	} else if len(nums2) == 0 {
		nums2 = nums1
	}
	sumLen := len(nums1) + len(nums2)
	i := 0
	j := 0
	num := 0
	next := 0
	for {
		if j == len(nums2) {
			num = nums1[i]
			i = i + 1
		} else if i == len(nums1) {
			num = nums2[j]
			j = j + 1
		} else if i < len(nums1) && nums1[i] < nums2[j] {
			num = nums1[i]
			i = i + 1
		} else if j < len(nums2) && nums1[i] >= nums2[j] {
			num = nums2[j]
			j = j + 1
		}
		if sumLen%2 == 0 && i+j == sumLen/2 {
			if i < len(nums1) && j < len(nums2) {
				if nums1[i] < nums2[j] {
					next = nums1[i]
				} else {
					next = nums2[j]
				}
			} else if i == len(nums1) {
				next = nums2[j]
			} else {
				next = nums1[i]
			}
			return float64(num+next) / 2
		} else if sumLen%2 == 1 && i+j == sumLen/2 {
			if i < len(nums1) && j < len(nums2) {
				if nums1[i] < nums2[j] {
					next = nums1[i]
				} else {
					next = nums2[j]
				}
			} else if i == len(nums1) {
				next = nums2[j]
			} else {
				next = nums1[i]
			}
			return float64(next)
		}
	}
	return 0
}
