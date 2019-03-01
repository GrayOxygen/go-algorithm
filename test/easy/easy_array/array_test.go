package easy_array_test

import (
	"fmt"
	"go-algorithm/src/easy/easy_array"
	"testing"
)

//nums = [0,0,1,1,1,2,2,3,3,4]  --> [0,1,2,3,4]
//去掉有序数组的重复元素，并返回最终长度
func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 4, 5}
	numsLen := easy_array.RemoveDuplicates(nums)
	fmt.Println("长度：", numsLen, "数组：", nums)
}

func TestMaxProfit(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	maxProfits := easy_array.MaxProfit(nums)
	fmt.Println("最大利润：", maxProfits)

}

func TestMoveZeroes(t *testing.T) {
	nums := []int{1, 0, 1}
	easy_array.MoveZeroes(nums)
	fmt.Println("移动零：", nums)
}
