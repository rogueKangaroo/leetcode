package simple

func twoSum(nums []int, target int) []int {
	numMap := make(map[int][]int, 0)
	for i := 0; i < len(nums); i++ {
		array, ok := numMap[nums[i]]
		if ok {
			array = append(array, i)
			numMap[nums[i]] = array
		} else {
			arrayNew := make([]int, 0)
			arrayNew = append(arrayNew, i)
			numMap[nums[i]] = arrayNew
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target-nums[i] {
			if value, exist := numMap[target-nums[i]]; exist && len(value) >= 2 {
				arrayRes := make([]int, 0)
				arrayRes = append(arrayRes, value[0])
				arrayRes = append(arrayRes, value[1])
				return arrayRes
			} else {
				continue
			}
		} else {
			if value, exist := numMap[target-nums[i]]; exist {
				arrayRes := make([]int, 0)
				arrayRes = append(arrayRes, i)
				arrayRes = append(arrayRes, value[0])
				return arrayRes
			} else {
				continue
			}
		}
	}
	return nil
}
