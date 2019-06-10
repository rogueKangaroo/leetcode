package hard

/*
给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。

示例 1:

输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出: true
示例 2:

输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出: false

解法：

*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool,len(s1) + 1)
	for i:=0;i<len(s1)+1;i++ {
		dp_row := make([]bool,len(s2) + 1)
		dp[i] = dp_row
	}
	dp[0][0] = true
	for i:=1;i<=len(s1);i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for i:=1;i<=len(s2);i++ {
		dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
	}
	for i:=1;i<=len(s1);i++ {
		for j:=1;j<=len(s2);j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i-1+j]) || (dp[i][j-1] && s2[j-1] == s3[i-1+j])
		}
	}
	return dp[len(s1)][len(s2)]
}