package medium

/*
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]

解法：首先了解IP规则，IP地址是32位二进制数，通常被分割为4个8位二进制数，换成十进制就是不大于256
回溯算法
*/
func restoreIpAddresses(s string) []string {
	return []string{}
}

func restoreIpAddressesStr(numArray []rune,str string,ipLength int) (bool,string) {
	if ipLength == 4 && len(numArray) > 0{
		return false,""
	}
	if numArray[0] > []rune("5")[0] {
		return false,""
	}else {
		ipLength ++
		return restoreIpAddressesStr(numArray[1:],str + "." + string(numArray[0]),ipLength)
	}
}