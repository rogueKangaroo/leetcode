package simple

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true

**/
func IsValid(s string) bool {
	small_left_rune := []rune("(")[0]
	small_right_rune := []rune(")")[0]
	middle_left_rune := []rune("[")[0]
	middle_right_rune := []rune("]")[0]
	big_left_rune := []rune("{")[0]
	big_right_rune := []rune("}")[0]
	index := -1
	s_list := make([]rune,len(s))
	s_rune := []rune(s)
	for i:=0;i<len(s_rune);i++ {
		if s_rune[i] == small_left_rune || s_rune[i] == middle_left_rune || s_rune[i] == big_left_rune {
			index ++
			s_list[index] = s_rune[i]
		} else if s_rune[i] == small_right_rune {
			if index >= 0 && s_list[index] == small_left_rune {
				index --
			} else {
				return false
			}
		}else if s_rune[i] == middle_right_rune {
			if index >= 0 && s_list[index] == middle_left_rune {
				index --
			} else {
				return false
			}
		}else if s_rune[i] == big_right_rune {
			if index >= 0 && s_list[index] == big_left_rune {
				index --
			} else {
				return false
			}
		}
	}
	if index >= 0 {
		return false
	}
	return true
}
