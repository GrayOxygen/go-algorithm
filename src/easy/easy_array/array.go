package easy_array

import "go-algorithm/src/util/array_util"

//问题：去掉有序数组的重复元素，并返回最终长度
//
// nums = [0,0,1,1,1,2,2,3,3,4]  --> [0,1,2,3,4]
// 原地法：利用现有的空间，减少空间复杂度，如下时空复杂均为o(n)
func RemoveDuplicates(nums []int) int {
	newIndex := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[newIndex] {
			newIndex++
			nums[newIndex] = nums[j]
		}
	}
	return newIndex + 1
}

//问题：
//买股票的最大收益：数组元素为顺序的每天股价，由左往右，买入后才能卖出，买入前必须卖出
//[7,1,5,3,6,4]-->4+3=7
//
//解决：买入卖出可看成一个区间，每个区间并无交集，比如X1...Xm-1 Xm Xm+1...Xk...Xp  假设Xp最大，Xm是中间变小的数，X1到Xm-1一直是递增的，Xm+1到Xp是递增的，这符合普遍规律
//那么，最大差值和一定是Xm-1 - X1 的差值加上Xp-Xm+1的差值合，如果是X1到Xm右边某个值的话，如Xk，那么Xk到Xp这个区间前者已经包括，而此时，X1到Xk Xk到Xp，其中X1到Xk包含了Xm-1到Xm减少，故，差值更小
//时间复杂度 O(n*n) 空间复杂度O(n)
func MaxProfit(nums []int) int {
	profits := 0
	for left := 0; left < len(nums); left  ++ {
		max := nums[left]
		for right := left + 1; right < len(nums); right++ {
			if nums[right] < max {
				break
			} else {
				max = nums[right]
			}
		}
		if max > nums[left] {
			profits += max - nums[left]
		}
	}
	return profits
}

//问题：移动/翻转数组 [1,2,3,4,  5,6] -- 向右移动2个位置，就是[5,6,  1,2,3,4]
//
//解决：
func Rotate(nums []int, k int) {
	array_util.Reverse(nums, len(nums)-k%len(nums), len(nums)-1);
	array_util.Reverse(nums, 0, len(nums)-k%len(nums)-1);
	array_util.Reverse(nums, 0, len(nums)-1)
}

//问题：数组是否存在重复元素
//
//解决：先排序，再一次遍历判断，比如快排O(nlogn)那时间复杂度就这么快，主要看排序
//也可用hash来过滤重复元素，时间复杂度为O(n)
func ContainsDuplicate(nums []int) bool {
	eleMap := make(map[int]int, 0)
	for val := range nums {
		if _, ok := eleMap[val]; ok {
			return false
		}
		eleMap[val] = val
	}
	return true
}

//问题：数组中只一个元素仅出现一次，其他元素均出现两次，找到唯一的那个元素
//
//解决：XOR的特性：X^X=0 X^0=X  符合交换律A^B=B^A
func SingleNumber(nums []int) int {
	result := 0
	for val := range nums {
		result = result ^ val
	}
	return result
}

//TODO 两个数组的交集

//问题：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一，[1,2,3]加一就是[1,2,4]
//
//解决：进位的处理
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 { //不用进位直接返回
			digits[i] ++
			return digits
		}
		digits[i] = 0 //逢9进位
	}
	//最高位进位
	result := make([]int, len(digits)+1, len(digits)+1)
	result = append(result, 1)
	return result
}

//问题：移动零，[3,4,0,2,1,1,0,11]-->[3,4,2,1,1,11,0,0]保持非零数的相对顺序不变
//
//解决：双指针法，一个遍历指针，一个指针记住第一个0的位置，起始位置0和1
func MoveZeroes(nums []int) {
	firstZeroIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[firstZeroIndex] == 0 && nums[i] != 0 { //非0向左移动
			nums[firstZeroIndex], nums[i] = nums[i], nums[firstZeroIndex]
			firstZeroIndex++;
		}
		if nums[firstZeroIndex] != 0 && nums[i] == 0 { //定位第一个0
			firstZeroIndex = i;
		}
	}
}

//问题：两数之和，给定目标值，找出和为该目标值的两个元素的下标，假设答案只有一组，[1,2,3]-->0 1
//
//解决：利用hash结构，如下时间复杂度O(n)
func TwoSum(nums []int, target int) []int {
	eleMap := make(map[int]int, 0)
	for index, val := range nums {
		if _, ok := eleMap[target-val]; ok {
			return []int{eleMap[target-val], index}
		}
		eleMap[val] = index
	}
	return nil
}

//TODO 问题：有效的数独，九宫格，任一行中，任一列中，任一3*3的方格中，都不允许出现重复元素，每个位置只能是1-9，如下二位数组表示九宫格
//[
//["5","3",".",".","7",".",".",".","."],
//["6",".",".","1","9","5",".",".","."],
//[".","9","8",".",".",".",".","6","."],
//["8",".",".",".","6",".",".",".","3"],
//["4",".",".","8",".","3",".",".","1"],
//["7",".",".",".","2",".",".",".","6"],
//[".","6",".",".",".",".","2","8","."],
//[".",".",".","4","1","9",".",".","5"],
//[".",".",".",".","8",".",".","7","9"]
//]

//解决：利用bitmap算法思想，将存储的元素映射到数组（对应下标），要注意方块的遍历方式
func IsValidSudoku(board [][]byte) bool {
	//for r := 0; r < 9; r++ {
	//	rowBitMap := make([]int, 9, 9)
	//	columnBitMap := make([]int, 9, 9)
	//	cubeBitMap := make([]int, 9, 9)
	//	for c := 0; c < len(board); c++ {
	//		if board[r][c] != '.' { //行
	//			if rowBitMap[ board[r][c]-1 ] > 0 { //对应位图的位置，如果已经有值，则表示重复了
	//				return false
	//			}
	//			rowBitMap[ board[r][c]-1 ]++
	//		}
	//		if board[c][r] != '.' { //列
	//			if columnBitMap[ board[r][c]-1 ] > 0 { //对应位图的位置，如果已经有值，则表示重复了
	//				return false
	//			}
	//			columnBitMap[ board[r][c]-1 ]++
	//		}
	//		cubeRow := 3*(r/3) + c/3 //每行遍历时，每3个，就要对应cube一次换行
	//		cubeColumn := c % 3
	//		if cubeBitMap[cubeRow][cubeColumn] != '.' {
	//			continue
	//		}
	//		if (cubeBitMap[temp-1] > 0) {
	//			System.out.println(i + "=-=" + j);
	//			return false;
	//		}
	//		cubeBitMap[temp-1] = 1;
	//	}
	//}
return true
}

//旋转图像
func RotateMatrix(matrix [][]int)  {

}
