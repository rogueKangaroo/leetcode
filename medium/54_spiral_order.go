package medium

/*
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	resultArray := make([]int,0)
	spiralOrderArray(matrix,0,0,len(matrix),len(matrix[0]),&resultArray)
	return resultArray
}

func spiralOrderArray(matrix [][]int, mStart,nStart int,mLen,nLen int,resultArray *[]int) {
	if len(*resultArray) == len(matrix)*len(matrix[0]) {
		return
	}
	for i:=0;i<nLen;i++{
		*resultArray = append(*resultArray,matrix[mStart][i+nStart])
	}
	if len(*resultArray) == len(matrix)*len(matrix[0]) {
		return
	}
	for i:=1;i<mLen;i++{
		*resultArray = append(*resultArray,matrix[i+mStart][nStart+nLen-1])
	}
	if len(*resultArray) == len(matrix)*len(matrix[0]) {
		return
	}
	for i:=1;i<nLen;i++{
		*resultArray = append(*resultArray,matrix[mStart+mLen-1][nStart+nLen-1-i])
	}
	if len(*resultArray) == len(matrix)*len(matrix[0]) {
		return
	}
	for i:=1;i<mLen-1;i++{
		*resultArray = append(*resultArray,matrix[mStart+mLen-1-i][nStart])
	}
	spiralOrderArray(matrix,mStart+1,nStart+1,mLen-2,nLen-2,resultArray)
}
