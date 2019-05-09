package hard

import "fmt"

func isMatch(s string, p string) bool {
	s_rune := []rune(s)
	p_rune := []rune(p)
	if len(s_rune) == 0 && len(p_rune) == 0 {
		return true
	}
	if (len(s_rune) > 0 && len(p_rune) == 0) || (len(s_rune) == 0 && len(p_rune) > 0) {
		return false
	}
	p_rune_list := make([][]rune, 0)
	p_rune_list = append(p_rune_list, p_rune)
	var status []rune
	return sMatch(s_rune, p_rune_list, 0, status)
}
func sMatch(s_rune []rune, p_rune_list [][]rune, index int, status []rune) bool {
	hasMatch := false
	if index == len(s_rune) {
		return toZeroList(p_rune_list)
	}
	s_item := s_rune[index]
	p_rune_list_new_sum := make([][]rune, 0)
	status_new_sum := make([]rune, 0)
	if status == nil {
		status = make([]rune, 0)
		status = append(status, rune(0))
	}
	for i, p_rune_item := range p_rune_list {
		p_str := string(p_rune_item)
		fmt.Println(p_str)
		var hasMatch_new bool
		hasMatch_new, p_rune_list_new, status_new := matchRegular(s_item, p_rune_item, status[i])
		hasMatch = hasMatch || hasMatch_new
		for index, item := range p_rune_list_new {
			item_str := string(item)
			fmt.Println(item_str)
			p_rune_list_new_sum = append(p_rune_list_new_sum, item)
			status_new_sum = append(status_new_sum, status_new[index])
		}
	}
	if hasMatch == false {
		return false
	}
	index = index + 1
	return sMatch(s_rune, p_rune_list_new_sum, index, status_new_sum)
}

func matchRegular(s_item rune, p_rune []rune, status rune) (bool, [][]rune, []rune) {
	p_rune_list := make([][]rune, 0)
	status_list := make([]rune, 0)
	if len(p_rune) == 0 {
		return false, p_rune_list, status_list
	}
	dot := []rune(".")[0]
	starKey := []rune("*")[0]
	if len(p_rune) == 1 {
		if p_rune[0] == s_item || p_rune[0] == dot {
			return true, p_rune_list, status_list
		} else {
			if p_rune[0] == starKey && (status == s_item || status == dot) {
				p_rune_list = append(p_rune_list, getRune(0, p_rune))
				status_list = append(status_list, status)
				return true, p_rune_list, status_list
			}
			return false, p_rune_list, status_list
		}
	}

	p_first := p_rune[0]
	i := 1
	if p_first == dot || p_first == s_item {
		p_rune_list = append(p_rune_list, getRune(i, p_rune))
		status_list = append(status_list, p_first)
		j := i
		startFlag := false
		fist_status := p_first
		for {
			if j == len(p_rune) {
				break
			}
			if p_rune[j] == starKey {
				startFlag = true
				p_rune_list = append(p_rune_list, getRune(j+1, p_rune))
				status_list = append(status_list, fist_status)
				j++
			} else {
				if (p_rune[j] == dot || p_rune[j] == s_item) && startFlag {
					p_rune_list = append(p_rune_list, getRune(j+1, p_rune))
					status_list = append(status_list, p_rune[j])
					fist_status = p_rune[j]
					j++
				} else {
					break
				}
			}
		}
		return true, p_rune_list, status_list
	} else if p_first == starKey && (status == s_item || status == dot) {
		p_rune_list = append(p_rune_list, getRune(0, p_rune))
		status_list = append(status_list, status)
		j := i
		startFlag := true
		fist_status := p_first
		for {
			if j == len(p_rune) {
				break
			}
			if p_rune[j] == starKey {
				startFlag = true
				p_rune_list = append(p_rune_list, getRune(j+1, p_rune))
				status_list = append(status_list, fist_status)
				j++
			} else {
				if (p_rune[j] == dot || p_rune[j] == s_item) && startFlag {
					p_rune_list = append(p_rune_list, getRune(j, p_rune))
					status_list = append(status_list, p_rune[j])
					fist_status = p_rune[j]
					j++
				} else {
					break
				}
			}
		}
		return true, p_rune_list, status_list
	} else if p_first != s_item {
		j := 1
		for {
			if j >= len(p_rune) {
				return false, p_rune_list, status_list
			}
			if p_rune[j] == s_item {
				j++
				break
			}
			j++
		}
		p_rune_list = append(p_rune_list, getRune(j, p_rune))
		status_list = append(status_list, s_item)
		startFlag := false
		fist_status := s_item
		for {
			if j >= len(p_rune) {
				break
			}
			if p_rune[j] == starKey {
				startFlag = true
				p_rune_list = append(p_rune_list, getRune(j+1, p_rune))
				status_list = append(status_list, fist_status)
				j++
			} else {
				if (p_rune[j] == dot || p_rune[j] == s_item) && startFlag {
					p_rune_list = append(p_rune_list, getRune(j+1, p_rune))
					status_list = append(status_list, p_rune[j])
					fist_status = p_rune[j]
					j++
				} else {
					break
				}
			}
		}
		return true, p_rune_list, status_list
	} else {
		return false, p_rune_list, status_list
	}
}

func getRune(index int, p_rune []rune) []rune {
	p_rune_new := make([]rune, 0)
	newIndex := 0
	for i := index; i < len(p_rune); i++ {
		p_rune_new = append(p_rune_new, p_rune[i])
		newIndex++
	}
	return p_rune_new
}

func toZeroList(p_rune_list [][]rune) bool {
	if len(p_rune_list) == 0 {
		return true
	}
	hasMatch := false
	for _, p_rune_item := range p_rune_list {
		var hasMatch_new bool
		hasMatch_new = toZero(p_rune_item)
		hasMatch = hasMatch || hasMatch_new
	}
	return hasMatch
}

func toZero(p_rune_list []rune) bool {
	if len(p_rune_list) == 0 {
		return true
	}
	starKey := []rune("*")[0]
	flag := false
	for _, p_rune_item := range p_rune_list {
		if p_rune_item == starKey {
			flag = true
		} else {
			flag = false
		}
	}
	return flag
}
