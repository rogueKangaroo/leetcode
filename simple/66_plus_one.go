package simple

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
*/
func plusOne(digits []int) []int {
	newInt := make([]int,0)
	flag := 1
	for i:=len(digits)-1;i>=0;i-- {
		num := digits[i] + flag
		flag = 0
		if num == 10 {
			newInt = append(newInt,0)
			flag = 1
		} else {
			newInt = append(newInt,num)
		}
	}
	if flag == 1 {
		newInt = append(newInt,1)
	}
	resultInt := make([]int,0)

	for i:=len(newInt)-1;i>=0;i--{
		resultInt = append(resultInt,newInt[i])
	}
	return resultInt
}