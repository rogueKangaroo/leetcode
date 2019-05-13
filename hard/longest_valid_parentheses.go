package hard

func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	s_rune := []rune(s)
	left := []rune("(")[0]
	index := 0
	flag := false
	flagNumber := 0
	maxNumber := 0
	number := 0
	for i := 0; i < len(s_rune); i++ {
		if s_rune[i] == left {
			stack = appendStack(stack, 2)
			index++
		} else {
			if index == 0 {
				continue
			}
			tempNumber := stack[index-1]
			if flag && index == 1 {
				number = tempNumber + flagNumber
				flagNumber = 0
			} else {
				number = tempNumber
			}
			stack = stack[:len(stack)-1]
			index--
			if index == 0 {
				flag = true
				flagNumber = number
			}
		}
		if number > maxNumber {
			maxNumber = number
		}
	}
	return maxNumber
}

func appendStack(stack []int, number int) []int {
	for i := 0; i < len(stack); i++ {
		stack[i] = stack[i] + number
	}
	stack = append(stack, number)
	return stack
}
