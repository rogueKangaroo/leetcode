package simple

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_sortedArrayToBST(t *testing.T) {
	nums := []int{-10,-3,0,5,9}
	got := sortedArrayToBST(nums)
	nodeList := []*TreeNode_sortedArrayToBST{got}
	for {
		if len(nodeList) == 0 {
			break
		} else {
			tempNodeList := make([]*TreeNode_sortedArrayToBST,0)
			for _,node := range nodeList{
				fmt.Print(strconv.Itoa(node.Val)+",")
				if node.Left != nil {
					tempNodeList = append(tempNodeList,node.Left)
				}
				if node.Right != nil {
					tempNodeList = append(tempNodeList,node.Right)
				}
			}
			nodeList = tempNodeList
		}
	}

	time.Sleep(1 * time.Second)
}
