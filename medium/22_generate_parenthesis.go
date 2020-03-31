package medium

/*
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

// 解法：动态规划
*/
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	} else if n == 1 {
		return []string{"()"}
	}
	dpStrArray := make([][]string,0)
	dpStrArray = append(dpStrArray,[]string{})
	dpStrArray = append(dpStrArray,[]string{"()"})
	for i:=2; i<=n; i++ {
		dpTempArray := make([]string,0)
		for j:=0 ;j < i;j++ {
			if len(dpStrArray[j]) == 0 {
				element1 := ""
				for _,element2 := range dpStrArray[i-j-1] {
					tempStr := "(" + element1 + ")" + element2
					dpTempArray = append(dpTempArray,tempStr)
				}
			} else if len(dpStrArray[i-j-1]) == 0 {
				element2 := ""
				for _,element1 := range dpStrArray[j] {
					tempStr := "(" + element1 + ")" + element2
					dpTempArray = append(dpTempArray,tempStr)
				}
			} else {
				for _,element1 := range dpStrArray[j] {
					for _,element2 := range dpStrArray[i-j-1] {
						tempStr := "(" + element1 + ")" + element2
						dpTempArray = append(dpTempArray,tempStr)
					}
				}
			}
		}
		dpStrArray = append(dpStrArray,dpTempArray)
	}
	return dpStrArray[n]
}