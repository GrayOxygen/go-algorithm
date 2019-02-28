package array

import (
	"fmt"
	"go-algorithm/src/easy/array"
	"testing"
)

//nums = [0,0,1,1,1,2,2,3,3,4]  --> [0,1,2,3,4]
//去掉有序数组的重复元素，并返回最终长度
func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 4, 5}
	numsLen := array.RemoveDuplicates(nums)
	fmt.Println("长度：", numsLen, "数组：", nums)
}

func TestMaxProfit(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	maxProfits := array.MaxProfit(nums)
	fmt.Println("最大利润：", maxProfits)
}
