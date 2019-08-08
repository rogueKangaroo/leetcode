package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_candy(t *testing.T) {
	ratings := []int{1,0,2}
	got := candy(ratings)
	fmt.Println(fmt.Sprintf("candy result(5) is %v", got))
	time.Sleep(100 * time.Millisecond)

	ratings = []int{1,2,2}
	got = candy(ratings)
	fmt.Println(fmt.Sprintf("candy result(4) is %v", got))
	time.Sleep(100 * time.Millisecond)

	ratings = []int{1,3,2,2,1}
	got = candy(ratings)
	fmt.Println(fmt.Sprintf("candy result(7) is %v", got))
	time.Sleep(100 * time.Millisecond)

	ratings = []int{1,4,3,2,1,2,2,3,1}
	got = candy(ratings)
	fmt.Println(fmt.Sprintf("candy result(17) is %v", got))
	time.Sleep(100 * time.Millisecond)

	ratings = []int{25,94,89,54,26,54,54,99,64,71,76,63,81,82,60,64,29,51,87,87,72}
	got = candy(ratings)
	fmt.Println(fmt.Sprintf("candy result(41) is %v", got))
	time.Sleep(100 * time.Millisecond)

	ratings = []int{58,21,72,77,48,9,38,71,68,77,82,47,25,94,89,54,26,54,54,99,64,71,76,63,81,82,60,64,29,51,87,87,72,12,16,20,21,54,43,41,83,77,41,61,72,82,15,50,36,69,49,53,92,77,16,73,12,28,37,41,79,25,80,3,37,48,23,10,55,19,51,38,96,92,99,68,75,14,18,63,35,19,68,28,49,36,53,61,64,91,2,43,68,34,46,57,82,22,67,89}
	got = candy(ratings)
	fmt.Println(fmt.Sprintf("candy result(208) is %v", got))
	time.Sleep(100 * time.Millisecond)
}
