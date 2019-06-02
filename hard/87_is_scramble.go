package hard

import "strconv"

/*
给定一个字符串 s1，我们可以把它递归地分割成两个非空子字符串，从而将其表示为二叉树。

下图是字符串 s1 = "great" 的一种可能的表示形式。

    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
在扰乱这个字符串的过程中，我们可以挑选任何一个非叶节点，然后交换它的两个子节点。

例如，如果我们挑选非叶节点 "gr" ，交换它的两个子节点，将会产生扰乱字符串 "rgeat" 。

    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
我们将 "rgeat” 称作 "great" 的一个扰乱字符串。

同样地，如果我们继续将其节点 "eat" 和 "at" 进行交换，将会产生另一个新的扰乱字符串 "rgtae" 。

    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
我们将 "rgtae” 称作 "great" 的一个扰乱字符串。

给出两个长度相等的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。

示例 1:

输入: s1 = "great", s2 = "rgeat"
输出: true
示例 2:

输入: s1 = "abcde", s2 = "caebd"
输出: false

*/
func isScramble(s1 string, s2 string) bool {
	str1 := ""
	str2 := ""
	map1 := make(map[uint8][]int,0)
	map2 := make(map[uint8][]int,0)
	for i:=0;i<len(s1);i++ {
		if _,ok := map1[s1[i]];ok {
			array := map1[s1[i]]
			array = append(array,i)
			map1[s1[i]] = array
		} else {
			array := make([]int,0)
			array = append(array,i)
			map1[s1[i]] =array
		}
	}
	for i:=0;i<len(s2);i++ {
		if _,ok := map2[s2[i]];ok {
			array := map2[s2[i]]
			array = append(array,i)
			map2[s2[i]] = array
		} else {
			array := make([]int,0)
			array = append(array,i)
			map2[s2[i]] =array
		}
	}
	for i:=0;i<len(s1);i++ {
		if array,ok := map2[s1[i]];ok {
			flag := false
			for index,item := range array {
				if item != -1 {
					flag = true
					str1 = str1 + strconv.Itoa(item)
					array[index] = -1
					map2[s1[i]] = array
					break
				}
			}
			if !flag {
				return false
			}
		} else {
			return false
		}
	}
	for i:=0;i<len(s2);i++ {
		if array,ok := map1[s2[i]];ok {
			flag := false
			for index,item := range array {
				if item != -1 {
					flag = true
					str2 = str2 + strconv.Itoa(item)
					array[index] = -1
					map1[s2[i]] = array
					break
				}
			}
			if !flag {
				return false
			}
		} else {
			return false
		}
	}
	if (s1 != "" && str1 == "") || (s2 != "" && str2 == "") {
		return false
	}
	if str1 == str2 {
		return true
	} else {
		return false
	}
}
