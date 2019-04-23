package medium

import (
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	src := []rune(str)
	if len(src) == 0 || src[0] != rune('-') && src[0] != rune('+') && (src[0] < rune('0') || src[0] > rune('9')) {
		return 0
	}
	res := ""
	flag := true
	for _, item := range src {
		if (item == rune('-') || item == rune('+')) && flag {
			res = res + string(item)
		} else if item >= rune('0') && item <= rune('9') {
			res = res + string(item)
			flag = false
		} else {
			break
		}
	}
	num, _ := strconv.Atoi(res)
	max := ^uint32(0) / 2
	min := 0 - (int64(max) + 1)
	if int64(num) > int64(max) {
		return int(max)
	}
	if int64(num) < int64(min) {
		return int(min)
	}
	return num
}
