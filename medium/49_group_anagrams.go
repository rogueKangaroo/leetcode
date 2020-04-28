package medium

/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。
*/
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	charList := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	for i:=0;i<len(strs);i++ {
		key := sortString(strs[i],charList)
		if value,exist := strMap[key];exist{
			value = append(value,strs[i])
			strMap[key] = value
		} else {
			item := []string{strs[i]}
			strMap[key] = item
		}
	}
	res := make([][]string,0)
	for _,v := range strMap {
		res = append(res,v)
	}
	return res
}

func sortString(str string,charList []string) string {
	charMap := make(map[string]int)
	strRunes := []rune(str)
	for i:=0;i<len(strRunes);i++ {
		if value,exist := charMap[string(strRunes[i])];exist {
			charMap[string(strRunes[i])] = value + 1
		} else {
			charMap[string(strRunes[i])] = 1
		}
	}
	res := ""
	for i:=0;i<len(charList);i++ {
		if value,exist := charMap[charList[i]];exist {
			for j:=0;j<value;j++{
				res = res + charList[i]
			}
		}
	}
	return res
}