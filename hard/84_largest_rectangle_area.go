package hard

func largestRectangleArea(x []int) int {
	stack := make([]int, len(x))
	index := 0
	maxRectangle := 0
	for i := 0; i < len(x); i++ {
		if index == 0 || x[stack[index-1]] < x[i] {
			stack[index] = i
			index++
		} else {
			index--
			width := 0
			if index == 0 {
				width = i
			} else {
				width = i - stack[index-1] - 1
			}
			rectangle := width * x[stack[index]]
			if rectangle > maxRectangle {
				maxRectangle = rectangle
			}
			i--
		}
	}
	for i := index - 1; i >= 0; i-- {
		index--
		width := 0
		if index == 0 {
			width = len(x)
		} else {
			width = len(x) - stack[index-1] - 1
		}
		rectangle := width * x[stack[index]]
		if rectangle > maxRectangle {
			maxRectangle = rectangle
		}
	}
	return maxRectangle
}
