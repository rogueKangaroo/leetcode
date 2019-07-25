package simple

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	num := GetStrNumArray(needle)
	index := 0
	for i:=0;i<len(haystack);i++ {
		if i==90 {
			fmt.Println("")
		}
		if haystack[i] == needle[index] {
			if index == len(needle) - 1 {
				return i - len(needle) + 1
			}
			index ++
		} else {
			if index > 0 {
				index  = num[index-1]
				i--
			}
		}
	}
	return -1
}

func GetStrNumArray(needle string) []int {
	num := make([]int,len(needle))
	for i:=0;i<len(needle);i++{
		num[i] = 0
	}
	for i := 1;i<len(needle);i++ {
		pre_map := make(map[string]string,0)
		for j:=1;j<=i;j++ {
			keyStr := needle[0:j]
			pre_map[keyStr] = keyStr
		}
		for j:=0;j<=i-1;j++ {
			keyStr := needle[j+1:i+1]
			if _,exist := pre_map[keyStr];exist {
				if num[i] < len(keyStr) {
					num[i] = len(keyStr)
				}
			}
		}
	}
	return num
}
