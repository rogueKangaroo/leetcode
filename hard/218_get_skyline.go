package hard
/**
城市的天际线是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。现在，假设您获得了城市风光照片（图A）上显示的所有建筑物的位置和高度，请编写一个程序以输出由这些建筑物形成的天际线（图B）。



每个建筑物的几何信息用三元组 [Li，Ri，Hi] 表示，其中 Li 和 Ri 分别是第 i 座建筑物左右边缘的 x 坐标，Hi 是其高度。可以保证 0 ≤ Li, Ri ≤ INT_MAX, 0 < Hi ≤ INT_MAX 和 Ri - Li > 0。您可以假设所有建筑物都是在绝对平坦且高度为 0 的表面上的完美矩形。

例如，图A中所有建筑物的尺寸记录为：[ [2 9 10], [3 7 15], [5 12 12], [15 20 10], [19 24 8] ] 。

输出是以 [ [x1,y1], [x2, y2], [x3, y3], ... ] 格式的“关键点”（图B中的红点）的列表，它们唯一地定义了天际线。关键点是水平线段的左端点。请注意，最右侧建筑物的最后一个关键点仅用于标记天际线的终点，并始终为零高度。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。

例如，图B中的天际线应该表示为：[ [2 10], [3 15], [7 12], [12 0], [15 10], [20 8], [24, 0] ]。

说明:

任何输入列表中的建筑物数量保证在 [0, 10000] 范围内。
输入列表已经按左 x 坐标 Li  进行升序排列。
输出列表必须按 x 位排序。
输出天际线中不得有连续的相同高度的水平线。例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...] 是不正确的答案；三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]
**/
type Node struct {
	Height int
	Next *Node
}

func getSkyline(buildings [][]int) [][]int {
	res := make([][]int,0)
	if len(buildings) == 0 {
		return res
	}
	newBuildings := handleSkyline(buildings)
	heightNode := &Node{
		Height:-newBuildings[0][1],
	}
	res0 := make([]int,2)
	res0[0] = newBuildings[0][0]
	res0[1] = -newBuildings[0][1]
	res = append(res,res0)
	last := make([]int,2)
	last[0] = newBuildings[0][0]
	last[1] = -newBuildings[0][1]
	for i:=1;i<len(newBuildings);i++ {
		if newBuildings[i][1] < 0 {
			heightNode = InsertHeight(heightNode,-newBuildings[i][1])
		} else {
			heightNode = DelHeight(heightNode,newBuildings[i][1])
		}
		if heightNode != nil {
			maxHeight := heightNode.Height
			if last[1] != maxHeight {
				last[0] = newBuildings[i][0]
				last[1] = maxHeight
				resItem := make([]int,2)
				resItem[0] = newBuildings[i][0]
				resItem[1] = maxHeight
				res = append(res,resItem)
			}
		} else {
			last[0] = newBuildings[i][0]
			last[1] = 0
			resItem := make([]int,2)
			resItem[0] = newBuildings[i][0]
			resItem[1] = 0
			res = append(res,resItem)
		}

	}
	return handleRes(res)
}

func handleRes(res [][]int) [][]int{
	newRes := make([][]int,0)
	index := 0
	last_X := -1
	for i:=0;i<len(res);i++ {
		if res[i][0] == last_X {
			index --
			newRes[index][1] = res[i][1]
			index ++
		} else {
			newResItem := make([]int,2)
			newResItem[0] = res[i][0]
			newResItem[1] = res[i][1]
			newRes = append(newRes,newResItem)
			last_X = res[i][0]
			index ++
		}
	}
	return newRes
}

func handleSkyline(buildings [][]int) [][]int {
	res := make([][]int,0)
	for i:=0;i<len(buildings);i++ {
		resItem0 := make([]int,2)
		resItem0[0] = buildings[i][0]
		resItem0[1] = -buildings[i][2]
		resItem1 := make([]int,2)
		resItem1[0] = buildings[i][1]
		resItem1[1] = buildings[i][2]
		res = append(res,resItem0)
		res = append(res,resItem1)
	}
	temp := make([][]int,2 * len(buildings))
	sortBuildings(&res,0,len(res)-1,temp)
	return res
}

// 使用归并排序
func sortBuildings(buildingPoints *[][]int,start,end int,temp [][]int){
	if start < end {
		mid := (end+start)/2
		sortBuildings(buildingPoints,start,mid,temp)
		sortBuildings(buildingPoints,mid+1,end,temp)
		mergeList(*buildingPoints,start,mid,end,temp)
	}
}

func mergeList(buildingPoints [][]int,start,mid,end int,temp[][]int)  {
	s := start
	m1 := mid
	m2 := mid + 1
	l := end
	i := 0
	for s <= m1 && l >= m2{
		if buildingPoints[s][0] < buildingPoints[m2][0] {
			temp[i] = buildingPoints[s]
			i = i + 1
			s = s + 1
		} else {
			temp[i] = buildingPoints[m2]
			i = i + 1
			m2 = m2 + 1
		}
	}
	for s <= m1 {
		temp[i] = buildingPoints[s]
		i = i + 1
		s = s + 1
	}

	for m2 <= l {
		temp[i] = buildingPoints[m2]
		i = i + 1
		m2 = m2 + 1
	}
	for j := 0;j < i;j++ {
		buildingPoints[start+j] = temp[j]
	}
}
func DelHeight(heightNode *Node,height int) *Node  {
	if heightNode.Height == height {
		return heightNode.Next
	}
	node := heightNode
	for {
		if node.Next != nil && node.Next.Height != height {
			node = node.Next
		} else {
			if node.Next != nil && node.Next.Height == height {
					node.Next = node.Next.Next
			}
			break
		}
	}
	return heightNode
}

func InsertHeight(heightNode *Node,height int) *Node {
	if heightNode == nil || heightNode.Height < height {
		newNode := &Node{
			Height:height,
			Next:heightNode,
		}
		return newNode
	}
	node := heightNode
	for {
		if node.Next != nil && node.Next.Height >= height {
			node = node.Next
		} else {
			newNode := &Node{
				Height:height,
			}
			if node.Next == nil {
				node.Next = newNode
			} else {
				newNode.Next = node.Next
				node.Next = newNode
			}
			break
		}
	}
	return heightNode
}
