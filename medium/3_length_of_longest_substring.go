package medium

// tmmzuxt
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	strMap := make(map[rune]int)
	max := 0
	temp := 0
	r := []rune(s)
	flag := 0
	for index, str := range r {
		if value, exist := strMap[str]; exist {
			if value >= flag {
				if temp > max {
					max = temp
					temp = 0
				}
				strMap[str] = index
				temp = index - value
				flag = value
			} else {
				temp = temp + 1
				strMap[str] = index
			}
		} else {
			temp = temp + 1
			strMap[str] = index
		}
	}
	if temp > max {
		max = temp
	}
	return max
}
