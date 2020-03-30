package medium

/*

给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]

解法：dp动态规划，本次结果依赖上次结果+本次值的组合
*/
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	numStrMap := map[string][]string{"2":{"a","b","c"},"3":{"d","e","f"},"4":{"g","h","i"},"5":{"j","k","l"},"6":{"m","n","o"},"7":{"p","q","r","s"},"8":{"t","u","v"},"9":{"w","x","y","z"}}
	numStrArray := []rune(digits)
	temp := numStrMap[string(numStrArray[0])]
	for i:=1;i<len(numStrArray);i++ {
		newTemp := make([]string,0)
		for j:=0;j<len(temp);j++ {
			currentNumArray := numStrMap[string(numStrArray[i])]
			for k:=0;k<len(currentNumArray);k++{
				newTemp = append(newTemp,temp[j]+currentNumArray[k])
			}
		}
		temp = newTemp
	}
	return temp
}
