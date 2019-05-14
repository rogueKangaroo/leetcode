package hard

/*
	给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

	'?' 可以匹配任何单个字符。
	'*' 可以匹配任意字符串（包括空字符串）。
	两个字符串完全匹配才算匹配成功。

	说明:

	s 可能为空，且只包含从 a-z 的小写字母。
	p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。

	dp解法:
	curMatch = (p(j) == ? && p(j) == s(i))
	f(n) = match(i,j) = match(i+1,j+1) || match(i+1,j) || match(i,j+1)  // p(j) == *
	f(n) = match(i,j) = match(i+1.j+1) && curMatch  // p(j) == ? or a-z
*/
func isMatch44(s string, p string) bool {
	s_rune := []rune(s)
	p_rune := []rune(p)
	questionMark := []rune("?")[0]
	starKey := []rune("*")[0]
	match := make([][]bool, 0)
	for i := 0; i < len(s_rune)+1; i++ {
		match_p := make([]bool, 0)
		for j := 0; j < len(p_rune)+1; j++ {
			match_p = append(match_p, false)
		}
		match = append(match, match_p)
	}
	match[len(s_rune)][len(p_rune)] = true
	for i := len(s_rune); i >= 0; i-- {
		for j := len(p_rune) - 1; j >= 0; j-- {
			if i == len(s_rune) {
				if p_rune[j] == starKey {
					match[i][j] = match[i][j+1]
				}
			} else {
				curMatch := p_rune[j] == questionMark || p_rune[j] == s_rune[i]
				if p_rune[j] == starKey {
					match[i][j] = match[i+1][j+1] || match[i+1][j] || match[i][j+1]
				} else {
					match[i][j] = match[i+1][j+1] && curMatch
				}
			}
		}
	}
	return match[0][0]
}
