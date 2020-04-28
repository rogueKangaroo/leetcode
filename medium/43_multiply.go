package medium

import "strconv"

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理
*/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	num1Rune := []rune(num1)
	num2Rune := []rune(num2)
	res := ""
	flag := 0
	indexStr := ""
	for i:=0;i<len(num1Rune);i++ {
		resi := ""
		for j:=0;j<len(num2Rune);j++ {
			num1Int,_ := strconv.Atoi(string(num1Rune[len(num1Rune)-1-i]))
			num2Int,_ := strconv.Atoi(string(num2Rune[len(num2Rune)-1-j]))
			item := num1Int * num2Int + flag
			flag = item/10
			item = item%10
			resi = strconv.Itoa(item) + resi
		}
		if flag > 0 {
			resi = strconv.Itoa(flag) + resi
			flag = 0
		}
		res = strAdd(res, resi + indexStr)
		indexStr = indexStr + "0"
	}
	return res
}

func strAdd(num1 string,num2 string) string {
	num1Rune := []rune(num1)
	num2Rune := []rune(num2)
	res := ""
	flag := 0
	if len(num1Rune) > len(num2Rune) {
		for i:=0;i<len(num1Rune);i++ {
			item := 0
			if len(num2Rune) - 1 - i >=0 {
				num1Int,_ := strconv.Atoi(string(num1Rune[len(num1Rune)-1-i]))
				num2Int,_ := strconv.Atoi(string(num2Rune[len(num2Rune)-1-i]))
				item = num1Int + num2Int + flag
			}else {
				num1Int,_ := strconv.Atoi(string(num1Rune[len(num1Rune)-1-i]))
				item = num1Int + flag
			}
			if item >= 10 {
				flag = 1
				item = item - 10
			} else {
				flag = 0
			}
			res = strconv.Itoa(item) + res

		}
	} else {
		for i:=0;i<len(num2Rune);i++ {
			item := 0
			if len(num1Rune) - 1 - i >=0 {
				num1Int,_ := strconv.Atoi(string(num1Rune[len(num1Rune)-1-i]))
				num2Int,_ := strconv.Atoi(string(num2Rune[len(num2Rune)-1-i]))
				item = num1Int + num2Int + flag
			}else {
				num2Int,_ := strconv.Atoi(string(num2Rune[len(num2Rune)-1-i]))
				item = num2Int + flag
			}
			if item >= 10 {
				flag = 1
				item = item - 10
			} else {
				flag = 0
			}
			res = strconv.Itoa(item) + res

		}
	}
	if flag == 1 {
		res = "1" + res
	}
	return res
}