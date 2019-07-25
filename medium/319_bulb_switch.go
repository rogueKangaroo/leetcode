package medium

func bulbSwitch(n int) int {
	if n <=1 {
		return n
	}
	num := n/2
	temp := num
	for {
		if num * num <= n && (num + 1) * (num + 1) > n {
			return num
		} else if num * num < n && (num + 1) * (num + 1) == n {
			return num + 1
		} else if num * num < n && (num + 1) * (num + 1) < n {
			num = num + (temp - num)/2
		} else {
			temp = num
			num = num/2
		}
	}
}


