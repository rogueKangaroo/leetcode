package simple

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 num1 成为一个有序数组。

 

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 

示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]


nums1 = [1,2,3,4,7,8,9,0,0,0], m = 7
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]

解法：遍历倒叙插入
*/
func merge(nums1 []int, m int, nums2 []int, n int)  {
	index := m + n - 1
	for {
		if n == 0 {
			break
		}
		if m==0 || nums1[m-1] <= nums2[n-1] {
			nums1[index] = nums2[n-1]
			index --
			n --
		} else {
			nums1[index] = nums1[m-1]
			index --
			m --
		}
	}
}
