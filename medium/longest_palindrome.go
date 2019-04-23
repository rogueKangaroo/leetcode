package medium

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	if isPalindrome([]rune(s)) {
		return s
	}
	src := []rune(s)
	max := 1
	str := string(src[0])
	for index, _ := range src {
		if index > 0 && index < (len(src)-1) {
			i := 1
			for {
				if !isPalindrome(src[index-i : index+i+1]) {
					break
				}
				i = i + 1
				if i > index || index+i > len(src) {
					break
				}
			}
			i = i - 1
			if 2*i+1 > max {
				max = 2*i + 1
				str = string(src[index-i : index+i+1])
			}

			j := 1
			for {
				if !isPalindrome(src[index-(j-1) : index+j+1]) {
					break
				}
				j = j + 1
				if j-1 > index || index+j > len(src) {
					break
				}
			}
			j = j - 1
			if 2*j > max {
				max = 2 * j
				str = string(src[index-(j-1) : index+j+1])
			}

		} else if index == 0 {
			k := 1
			for {
				if !isPalindrome(src[0 : k+1]) {
					break
				}
				if k == len(src) {
					break
				}
				k = k + 1
			}
			k = k - 1
			if k+1 > max {
				max = k + 1
				str = string(src[0 : k+1])
			}
		}
	}
	return str
}

func isPalindrome(src []rune) bool {
	srcNew := make([]rune, 2*len(src)+1)
	i := 0
	index := 0
	for {
		srcNew[i] = rune('#')
		i = i + 1
		srcNew[i] = src[index]
		i = i + 1
		index = index + 1
		if index >= len(src) {
			break
		}
	}
	srcNew[i] = rune('#')
	for i := 0; i < len(srcNew)/2; i++ {
		if srcNew[i] != srcNew[len(srcNew)-1-i] {
			return false
		}
	}

	return true
}
