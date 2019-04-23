package medium

func zConvert(s string, numRows int) string {
	src := []rune(s)
	num := numRows*2 - 2
	if num == 0 {
		return s
	}
	str := make([]string, numRows)
	res := ""
	for index, char := range src {
		divideNum := index % num
		if divideNum < numRows {
			str[divideNum] = str[divideNum] + string(char)
		} else {
			row := numRows - (divideNum - numRows + 1) - 1
			str[row] = str[row] + string(char)
		}
	}
	for _, strItem := range str {
		res = res + strItem
	}
	return res
}
