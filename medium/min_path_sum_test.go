package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_minPathSum(t *testing.T) {
	obstacleGrid := make([][]int, 0)
	obstacleGrid = append(obstacleGrid, []int{1, 3, 1})
	obstacleGrid = append(obstacleGrid, []int{1, 5, 1})
	obstacleGrid = append(obstacleGrid, []int{4, 2, 1})
	got := minPathSum(obstacleGrid)
	fmt.Println(fmt.Sprintf("minPathSum result is %v", got))
	time.Sleep(1 * time.Second)
}
