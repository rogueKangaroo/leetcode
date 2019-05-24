package simple

/*


解法：
	dp:
	f(n) = stair[i] = stair[i-1] + stair[i-2] // i > 2
*/
func climbStairs(m int) int {
	res := make([]int,0)
	res = append(res,1)
	res = append(res,1)
	for i:=2;i<=m;i++ {
		res = append(res,res[i-1] + res[i-2])
	}
	return res[m]
}
