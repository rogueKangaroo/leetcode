package simple

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/
func longestCommonPrefix(strs []string) string {
	// 特殊情况处理
	if len(strs) == 0 {
		return ""
	}else if len(strs) == 1 {
		return strs[0]
	}

	// 多重遍历
	j := 0
	sumStr := ""
	currentItem := ""
	strMap := make(map[string][]rune)
	for {
		for i:=0;i<len(strs);i++ {
			if len(strs[i]) == 0 {
				return ""
			} else {
				if i == 0 {
					if len(strs[i]) == j {
						return sumStr
					} else {
						currentItem = string([]rune(strs[i])[j])
					}
				}
				if j == 0 {
					strMap[strs[i]] = []rune(strs[i])
					if string([]rune(strs[i])[j]) != currentItem {
						return sumStr
					}
				} else {
					v,_ := strMap[strs[i]]
					if len(v) == j || string([]rune(strs[i])[j]) != currentItem{
						return sumStr
					}
				}
			}
		}
		j = j+ 1
		sumStr = sumStr + currentItem
	}
	return sumStr
}