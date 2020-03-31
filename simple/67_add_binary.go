package simple

import "strconv"

/*
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"

*/

func addBinary(a string, b string) string {
	aRune := []rune(a)
	bRune := []rune(b)
	str := ""
	if len(aRune) >= len(bRune){
		str = GetResStr(aRune,bRune)
	} else {
		str = GetResStr(bRune,aRune)
	}
	resStr := ""
	strRune := []rune(str)
	for i:=len(strRune)-1;i>=0;i--{
		resStr = resStr + string(strRune[i])
	}
	return resStr
}

func GetResStr(aRune,bRune[]rune) string {
	str := ""
	flag := 0
	for i:=0;i<=len(aRune)-1;i++ {
		aNum,_ :=strconv.Atoi(string(aRune[len(aRune)-1-i]))
		sum := 0
		if i >= len(bRune) {
			sum = aNum + flag
		} else {
			bNum,_ :=strconv.Atoi(string(bRune[len(bRune)-1-i]))
			sum = aNum + bNum + flag
		}
		if sum == 3 {
			flag = 1
			str = str + "1"
		} else if sum == 2 {
			flag = 1
			str = str + "0"
		}else if sum == 1 {
			str = str + "1"
			flag = 0
		} else {
			str = str + "0"
			flag = 0
		}
	}
	if flag == 1{
		str = str + "1"
	}
	return str
}
