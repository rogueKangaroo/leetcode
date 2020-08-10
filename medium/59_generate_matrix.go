package medium

/*
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]

*/
func generateMatrix(n int) [][]int {
	result := make([][]int,0)
	for i:=0;i<n;i++{
		item := make([]int,n)
		result = append(result,item)
	}
	generateMatrixN(1,n,0,&result,n*n)
	return result

}

func generateMatrixN(start,n,index int, result *[][]int,length int)  {
	if start > length {
		return
	}
	for i:=0;i<n;i++{
		(*result)[index][index+i] = start
		start ++
	}
	if start > length {
		return
	}
	for i:=1;i<n;i++{
		(*result)[index+i][index+n-1] = start
		start ++
	}
	if start > length {
		return
	}
	for i:=1;i<n;i++{
		(*result)[index+n-1][index+n-1-i] = start
		start ++
	}
	if start > length {
		return
	}
	for i:=1;i<n-1;i++{
		(*result)[index+n-1-i][index] = start
		start ++
	}
	generateMatrixN(start,n-2,index+1,result,length)
}