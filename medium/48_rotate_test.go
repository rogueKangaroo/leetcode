package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_swap(t *testing.T) {
	a, b, c, d := 0, 1, 2, 3
	d = a ^ d
	a = a ^ d
	d = a ^ d
	c = d ^ c
	d = d ^ c
	c = d ^ c
	b = c ^ b
	c = c ^ b
	b = c ^ b
	fmt.Println(fmt.Sprintf("swap result is a:%v,b:%v,c:%v,d:%v", a, b, c, d))
	time.Sleep(1 * time.Second)
}

//给定 matrix =
//[
//  [ 5, 1, 9,11],
//  [ 2, 4, 8,10],
//  [13, 3, 6, 7],
//  [15,14,12,16]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [15,13, 2, 5],
//  [14, 3, 4, 1],
//  [12, 6, 8, 9],
//  [16, 7,10,11]
//]
func Test_rotate(t *testing.T) {
	matrix := [][]int{
		{ 5, 1, 9,11},
		{ 2, 4, 8,10},
		{13, 3, 6, 7},
		{15,14,12,16},
	}
	rotate(matrix)
	fmt.Println(fmt.Sprintf("rotate result is %v", matrix))
	time.Sleep(1 * time.Second)
}

