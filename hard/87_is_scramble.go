package hard

import (
	"sort"
)

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

解法：
	如果s1和s2是扰乱字符，则必然存在从第i个字符分隔，使得s1(0->i)和s2(0->i)是扰乱字符且s1(i->len(s1))和s2(i->len(s2))是扰乱字符；或者s1(0->i)和s2(len(s2)-i->len(s2))是扰乱字符且s1(i->len(s1))和s2(0->len(s2)-i)是扰乱字符
	其余做好剪枝操作，提升效率
*/
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	if SortString(s1) != SortString(s2) {
		return false
	}

	for i:=1;i<len(s1);i++ {
		if (isScramble(s1[0:i],s2[0:i]) && isScramble(s1[i:] ,s2[i:])) || (isScramble(s1[0:i] , s2[len(s1) - i:]) && isScramble(s1[i:] , s2[0:len(s1) - i])) {
			return true
		}
	}

	return false
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}