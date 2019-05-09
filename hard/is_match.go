package hard

// 动态规划
/* 确认递归关系：

f(n) = match[i][j] = match[i][j+2] || (match[i+1][j]&&curMatch) // p[j+1] == '*',因为a*，*表示0到正无穷
f(n) = match[i][j] = match[i+1][j+1] && curMatch                // 当前位匹配且f(n-1)匹配。则f(n)匹配
f(0) = true     // 空值匹配空值
*/
func isMatch(s string, p string) bool {
	s_rune := []rune(s)
	p_rune := []rune(p)
	starKey := []rune("*")[0]
	dot := []rune(".")[0]
	var match [][]bool
	for i := 0; i < len(s_rune)+1; i++ {
		match_p := make([]bool, len(p_rune)+1)
		for j := 0; j < len(match_p); j++ {
			match_p[j] = false
		}
		match = append(match, match_p)
	}
	match[len(s_rune)][len(p_rune)] = true
	for i := len(s_rune); i >= 0; i = i - 1 {
		for j := len(p_rune) - 1; j >= 0; j = j - 1 {
			curMatch := false
			if i < len(s_rune) {
				curMatch = (s_rune[i] == p_rune[j]) || (p_rune[j] == dot)
			}
			if j < len(p_rune)-1 && p_rune[j+1] == starKey {
				if i == len(s_rune) {
					match[i][j] = match[i][j+2]
				} else {
					match[i][j] = match[i][j+2] || (match[i+1][j] && curMatch)
				}
			} else {
				if i == len(s_rune) {
					match[i][j] = false
				} else {
					match[i][j] = curMatch && match[i+1][j+1]
				}
			}
		}
	}
	return match[0][0]
}
