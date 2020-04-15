package simple

// 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
//
//
//
//在杨辉三角中，每个数是它左上方和右上方的数的和。
//
//示例:
//
//输入: 5
//输出:
//[
//     [1],
//    [1,1],
//   [1,2,1],
//  [1,3,3,1],
// [1,4,6,4,1]
//]
//
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	} else if numRows == 1 {
		return [][]int{[]int{1}}
	} else if numRows == 2 {
		return [][]int{{1},{1,1}}
	}
	result := make([][]int,0)
	result = append(result,[]int{1})
	result = append(result,[]int{1,1})
	for i:=2;i<numRows;i++ {
		temp := make([]int,0)
		for j:=0;j<=i;j++{
			if j == 0 {
				temp = append(temp,1)
			} else if j == i {
				temp = append(temp,1)
			} else {
				temp = append(temp,result[i-1][j-1] + result[i-1][j])
			}
		}
		result = append(result,temp)
	}
	return result
}