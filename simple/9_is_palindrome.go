package simple

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	src := []rune(str)
	srcNew := make([]rune, 2*len(src)+1)
	i := 0
	for _, char := range src {
		srcNew[i] = rune('#')
		i = i + 1
		srcNew[i] = char
		i = i + 1
	}
	srcNew[i] = rune('#')
	for index := 0; index < len(srcNew)/2; index++ {
		if srcNew[index] != srcNew[len(srcNew)-1-index] {
			return false
		}
	}
	return true
}
