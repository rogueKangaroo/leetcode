package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_getMaxRectangle(t *testing.T) {
	x := []int{0}
	got := getMaxRectangle(x)
	fmt.Println(fmt.Sprintf("getMaxRectangle result is %v", got))
	time.Sleep(100 * time.Millisecond)

	//x = []int{2, 1, 4, 5, 1, 3, 3}
	//got = getMaxRectangle(x)
	//fmt.Println(fmt.Sprintf("getMaxRectangle result is %v", got))
	//time.Sleep(100 * time.Millisecond)
}

func Test_maximalRectangle(t *testing.T) {
	matrix := make([][]byte, 0)
	matrix = append(matrix, []byte{'0', '0', '0', '1', '0', '1', '0'})
	matrix = append(matrix, []byte{'0', '1', '0', '0', '0', '0', '0'})
	matrix = append(matrix, []byte{'0', '1', '0', '1', '0', '0', '1'})
	matrix = append(matrix, []byte{'0', '0', '1', '1', '0', '0', '1'})
	matrix = append(matrix, []byte{'1', '1', '1', '1', '1', '1', '0'})
	matrix = append(matrix, []byte{'1', '0', '0', '1', '0', '1', '1'})
	matrix = append(matrix, []byte{'0', '1', '0', '0', '1', '0', '1'})
	matrix = append(matrix, []byte{'1', '1', '0', '1', '1', '1', '0'})
	matrix = append(matrix, []byte{'1', '0', '1', '0', '1', '0', '1'})
	matrix = append(matrix, []byte{'1', '1', '1', '0', '0', '0', '0'})
	got := maximalRectangle(matrix)
	fmt.Println(fmt.Sprintf("maximalRectangle result is %v", got))
	time.Sleep(100 * time.Millisecond)

	matrix = make([][]byte, 0)
	matrix = append(matrix, []byte{'0', '1', '1', '0', '1'})
	matrix = append(matrix, []byte{'1', '1', '0', '1', '0'})
	matrix = append(matrix, []byte{'0', '1', '1', '1', '0'})
	matrix = append(matrix, []byte{'1', '1', '1', '1', '0'})
	matrix = append(matrix, []byte{'1', '1', '1', '1', '1'})
	matrix = append(matrix, []byte{'0', '0', '0', '0', '0'})
	got = maximalRectangle(matrix)
	fmt.Println(fmt.Sprintf("maximalRectangle result is %v", got))
	time.Sleep(100 * time.Millisecond)
}
