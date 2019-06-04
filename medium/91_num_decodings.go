package medium

import "strconv"

/*
一条包含字母 A-Z 的消息通过以下方式进行了编码：

'A' -> 1
'B' -> 2
...
'Z' -> 26
给定一个只包含数字的非空字符串，请计算解码方法的总数。

示例 1:

输入: "12"
输出: 2
解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
示例 2:

输入: "226"
输出: 3
解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。

解法：
	f(n) = numDecodings(s) = numDecodings(1:len(s)) + numDecodings(2:len(s))  // s[0->2] <= 26
	f(n) = numDecodings(s) = numDecodings(1:len(s))                           // s[0->2] > 26
*/

func numDecodings(s string) int {
	if s[0] == "0"[0] {
		return 0
	}
	num,_ := strconv.Atoi(s)
	if num == 0 {
		return 0
	}
	if num <= 10 || num == 20{
		return 1
	}
	if num > 10 && num <=26 {
		return 2
	}
	flagNum,_ := strconv.Atoi(s[0:2])
	if flagNum <= 26 {
		return numDecodings(s[1:]) + numDecodings(s[2:])
	} else {
		return numDecodings(s[1:])
	}
}
