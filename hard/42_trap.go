package hard

/**
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
**/
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	max := 0
	sum := 0
	for i:=0;i<len(height);i++ {
		if height[i] > height[max] {
			max = i
		}
		sum = sum + height[i]
	}
	left_sum := 0
	left_max_height := 0
	for i:=0;i<max;i++ {
		if height[i] > left_max_height  {
			left_max_height = height[i]
		}
		left_sum = left_sum + left_max_height
	}

	right_sum := 0
	right_max_height := 0
	for i:=len(height)-1;i>max;i-- {
		if height[i] > right_max_height  {
			right_max_height = height[i]
		}
		right_sum = right_sum + right_max_height
	}
	return left_sum + right_sum - sum + height[max]
}