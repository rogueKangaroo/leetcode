package simple

/*
	给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

	示例:

	输入: [-2,1,-3,4,-1,2,1,-5,4],
	输出: 6
	解释: 连续子数组 [4,-1,2,1] 的和最大，为 6

	dp解法:
	temp = temp + |nums[i]| // nums[i] <= 0
	temp = 0              // nums[i] >0
	f(n) = max(i) = 0 			// nums[i] < 0
	f(n) = max(i) = max(i-1) + nums[i]  // nums[i] >= 0, max(i-1) >= temp
	f(n) = max(i) = nums[i] // nums[i]>=0,max(i-1)<temp
	max(0) = nums[0]
*/
func maxSubArray(nums []int) int {
	max := make([]int, 0)
	index := 0
	temp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			temp = temp - nums[i]
		} else {
			if index == 0 {
				max = append(max, nums[i])
				temp = 0
				index++
				continue
			}
			if max[index-1] >= temp {
				max = append(max, max[index-1]-temp+nums[i])
			} else {
				max = append(max, nums[i])
			}
			index++
			temp = 0
		}
	}
	if len(max) == 0 {
		maxNumber := nums[0]
		for i := 0; i < len(nums); i++ {
			if maxNumber < nums[i] {
				maxNumber = nums[i]
			}
		}
		return maxNumber
	}
	maxNumber := max[0]
	for i := 0; i < len(max); i++ {
		if maxNumber < max[i] {
			maxNumber = max[i]
		}
	}
	return maxNumber
}
