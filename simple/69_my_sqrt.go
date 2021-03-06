package simple

/*
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。

解法:
*/

func mySqrt(x int) int {
	if x ==0 {
		return 0
	} else if x < 4 {
		return 1
	}
	num := x/2
	end := 0
	for {
		if num*num > x{
			end = num
			num = num/2
		}else {
			break
		}
	}
	i:=num
	for ;i<=end;i++{
		if i*i > x{
			return i-1
		} else if i*i == x {
			return i
		}
	}
	return i
}
