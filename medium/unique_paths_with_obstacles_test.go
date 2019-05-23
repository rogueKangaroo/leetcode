package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_uniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := make([][]int, 0)
	obstacleGrid = append(obstacleGrid, []int{0, 0, 0})
	obstacleGrid = append(obstacleGrid, []int{0, 1, 0})
	obstacleGrid = append(obstacleGrid, []int{0, 0, 0})
	got := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(fmt.Sprintf("uniquePathsWithObstacles result is %v", got))
	time.Sleep(1 * time.Second)

	obstacleGrid = make([][]int, 0)
	obstacleGrid = append(obstacleGrid, []int{1})
	obstacleGrid = append(obstacleGrid, []int{0})
	got = uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(fmt.Sprintf("uniquePathsWithObstacles result is %v", got))
	time.Sleep(1 * time.Second)

	obstacleGrid = make([][]int, 0)
	obstacleGrid = append(obstacleGrid, []int{1, 0})
	got = uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(fmt.Sprintf("uniquePathsWithObstacles result is %v", got))
	time.Sleep(1 * time.Second)
}
