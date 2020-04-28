package medium

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_strAdd(t *testing.T) {
	got := strAdd("1111", "999")
	fmt.Println(fmt.Sprintf("strAdd result is %v", got))
	time.Sleep(1 * time.Second)
}

//输入: num1 = "123", num2 = "456"
//输出: "56088"
func Test_multiply(t *testing.T) {
	num1 := "123"
	num2 := "456"
	got := multiply(num1, num2)
	fmt.Println(fmt.Sprintf("strAdd result is %v", got))
	num1Int,_ := strconv.Atoi(num1)
	num2Int,_ := strconv.Atoi(num2)
	fmt.Println(fmt.Sprintf("strAdd right result is %v", num1Int * num2Int))
	time.Sleep(1 * time.Second)
}


func Test_multiply1(t *testing.T) {
	num1 := "2"
	num2 := "12"
	got := multiply(num1, num2)
	fmt.Println(fmt.Sprintf("strAdd result is %v", got))
	num1Int,_ := strconv.Atoi(num1)
	num2Int,_ := strconv.Atoi(num2)
	fmt.Println(fmt.Sprintf("strAdd right result is %v", num1Int * num2Int))
	time.Sleep(1 * time.Second)
}