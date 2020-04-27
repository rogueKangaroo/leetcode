package medium

/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

 

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2
 

提示：

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。

*/


func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend > 0 && divisor > 0 {
		if dividend == 2147483648 && divisor == 1 {
			return 2147483647
		} else {
			return div(dividend,divisor)
		}
	} else if dividend < 0 && divisor < 0 {
		if dividend == -2147483648 && divisor == -1 {
			return 2147483647
		} else {
			return div(0-dividend,0-divisor)
		}
	} else if dividend > 0 && divisor < 0 {
		return 0 - div(dividend,0-divisor)
	} else if dividend < 0 && divisor > 0 {
		return 0 - div(0-dividend,divisor)
	}

	return 1
}

func div(dividend int, divisor int) int  {
	flagDivisor := divisor
	if divisor > dividend {
		return 0
	} else if (divisor + divisor) > dividend {
		return 1
	}
	start := 1
	end := start
	startFlagDivisor := 0
	for {
		if flagDivisor == dividend {
			return start + start
		} else if flagDivisor >= dividend {
			break
		} else {
			startFlagDivisor = flagDivisor
			flagDivisor = flagDivisor + flagDivisor
			start = end
			end = end + end
		}
	}
	i:=start
	for ;i < end;i++ {
		startFlagDivisor = startFlagDivisor + divisor
		if startFlagDivisor > dividend {
			break
		}
	}

	return i
}
