package simple

import "strconv"

/*
二进制手表顶部有 4 个 LED 代表小时（0-11），底部的 6 个 LED 代表分钟（0-59）。

每个 LED 代表一个 0 或 1，最低位在右侧。



例如，上面的二进制手表读取 “3:25”。

给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。

案例:

输入: n = 1
返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
 

注意事项:

输出的顺序没有要求。
小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。
分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。

*/
func readBinaryWatch(num int) []string {
	if num > 8 {
		return []string{}
	}
	if num == 0 {
		return []string{"0:00"}
	}
	timeList := []int{1,2,4,8,1,2,4,8,16,32}
	allResult := make([]string,0)
	tempResMap := make(map[int]string)
	backTrackForReadBinaryWatch(timeList,tempResMap,&allResult,num,0)
	return allResult
}

func backTrackForReadBinaryWatch(timeList []int,tempResMap map[int]string,allResList *[]string, n int,index int)  {
	if getMapValidLength(tempResMap) == n {
		timeStr := convertIntToTime(tempResMap,timeList)
		*allResList = append(*allResList,timeStr)
		return
	}

	for i:=index;i<len(timeList);i++ {
		if isValidTime(tempResMap,i,timeList) {
			tempResMap[i] = "1"
			backTrackForReadBinaryWatch(timeList,tempResMap,allResList,n,i)
			tempResMap[i] = ""
		}
	}
}

func convertIntToTime(tempResMap map[int]string,timeList []int) string {
	hour := 0
	minute := 0
	for k,v := range  tempResMap {
		if v == "1" {
			if k <= 3 {
				hour = hour + timeList[k]
			} else {
				minute = minute + timeList[k]
			}
		}
	}
	timeStr := strconv.Itoa(hour) + ":"
	if minute >= 10 {
		timeStr =  timeStr + strconv.Itoa(minute)
	} else {
		timeStr = timeStr +"0" +  strconv.Itoa(minute)
	}

	return timeStr
}

func isValidTime(tempResMap map[int]string,index int,timeList []int) bool {
	if value,exist := tempResMap[index];exist && value == "1"{
		return false
	} else {
		hour := 0
		minute := 0
		for k,v := range tempResMap{
			if v == "1" {
				if k <= 3 {
					hour = hour + timeList[k]
				} else {
					minute = minute + timeList[k]
				}
			}
		}
		if index <= 3 {
			return hour + timeList[index] <= 11
		} else {
			return minute + timeList[index] <= 59
		}
	}
}

func getMapValidLength(tempResMap map[int]string) int {
	length := 0
	for _,v := range tempResMap{
		if v == "1" {
			length ++
		}
	}
	return length
}