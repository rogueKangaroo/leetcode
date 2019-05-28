package hard

/*
	我们用dp[i]表示以i结尾的最长有效括号;

	当s[i] 为(,dp[i] 必然等于0,因为不可能组成有效的括号;

	那么s[i] 为)

	2.1 当s[i-1] 为(,那么dp[i] = dp[i-2] + 2;

	2.2 当s[i-1] 为)并且 s[i-dp[i-1] - 1]为(,那么dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
*/
func longestValidParentheses(s string) int {
	left_rune := []rune("(")[0]
	right_rune := []rune(")")[0]
	s_rune := []rune(s)
	if len(s_rune) <= 1 {
		return 0
	}
	dp := make([]int, len(s_rune))
	for index, _ := range dp {
		dp[index] = 0
	}
	for i := 1; i < len(s_rune); i++ {
		if s_rune[i] == right_rune {
			if s_rune[i-1] == left_rune {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if s_rune[i-1] == right_rune && i-dp[i-1]-1 >= 0 && s_rune[i-dp[i-1]-1] == left_rune {
				if i-dp[i-1]-2 > 0 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
	}
	max := 0
	for _, item := range dp {
		if item > max {
			max = item
		}
	}
	return max
}
