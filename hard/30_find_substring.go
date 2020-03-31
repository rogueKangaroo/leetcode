package hard
/*
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

 

示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]

解法：先枚举随机串联字符串，然后遍历
*/
func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	resultInt := make([]int,0)
	allLength := len(words)*len(words[0])
	strMap := make(map[string]int)

	for i:=0;i<len(words);i++{
		if v,exist := strMap[words[i]];exist{
			v = v + 1
			strMap[words[i]] = v
		} else {
			strMap[words[i]] = 1
		}
	}

	for i:=0;i<=len(s)-allLength;i++{
		temp := s[i:i+len(words[0])]
		if _,exist := strMap[temp];exist{
			tempS := s[i:i+allLength]
			newStr := getSubString(tempS,len(words[0]))
			if subStringCompare(strMap,newStr){
				resultInt = append(resultInt,i)
			}
		}
	}

	return resultInt
}

func getSubString(s string,length int) []string {
	sArray := []rune(s)
	newStr := make([]string,0)
	for i:=0;i<len(sArray);i=i+length{
		str := ""
		for j:=0;j<length;j++{
			str = str + string(sArray[i+j])
		}
		newStr = append(newStr,str)
	}
	return newStr
}

func subStringCompare(strMap map[string]int, compareWords []string) bool {
	newMap := make(map[string]int)
	for i:=0;i<len(compareWords);i++ {
		if v,exist := strMap[compareWords[i]];exist {
			if num,existed := newMap[compareWords[i]];existed{
				num = num + 1
				newMap[compareWords[i]] = num
				if num > v {
					return false
				}
			} else {
				newMap[compareWords[i]] = 1
			}
		} else {
			return false
		}
	}
	return true
}