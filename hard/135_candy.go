package hard

/**
老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

你需要按照以下要求，帮助老师给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻的孩子中，评分高的孩子必须获得更多的糖果。
那么这样下来，老师至少需要准备多少颗糖果呢？

示例 1:

输入: [1,0,2]
输出: 5
解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。
示例 2:

输入: [1,2,2]
输出: 4
解释: 你可以分别给这三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这已满足上述两个条件。

1,4,3,2,1,2,2,3,1
**/
func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	last := 1
	sum := 1
	tempIndex := 0
	tempNum := 0
	flag := false
	for i := 1;i < len(ratings);i ++ {
		if ratings[i] > ratings[i-1] {
			last = last + 1
			sum = sum + last
			tempNum = last
			tempIndex = i
			flag = true
		} else if ratings[i] == ratings[i-1] {
			last = 1
			sum = sum + 1
			tempNum = 1
			tempIndex = i
			flag = false
		} else {
			last = 1
			if flag {
				sum = sum + last
				flag = false
			} else {
				if i - tempIndex > tempNum - 1 {
					sum = sum + (i - tempIndex) + 1
					tempNum ++
				} else {
					sum = sum + (i - tempIndex)
				}

			}
		}
	}
	return sum
}