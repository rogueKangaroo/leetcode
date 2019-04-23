package simple

import "strconv"

func intReverse(x int) int {
	res := ""
	str := strconv.Itoa(x)
	src := []rune(str)
	min := 0
	if string(src[0]) == "-" {
		res = res + "-"
		min = 1
	}
	flag := false
	for i := len(src) - 1; i >= min; i-- {
		if string(src[i]) != "0" || flag {
			res = res + string(src[i])
			flag = true
		}
	}
	num, _ := strconv.Atoi(res)
	max := ^uint32(0)/2 - 1
	minValue := 0 - (int64(max) + 1)
	if int64(num) > int64(max) {
		return 0
	}
	if int64(num) < int64(minValue) {
		return 0
	}
	return num
}
