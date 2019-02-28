package array

// 原地法：利用现有的空间，减少空间复杂度

// nums = [0,0,1,1,1,2,2,3,3,4]  --> [0,1,2,3,4]
//去掉有序数组的重复元素，并返回最终长度   时空复杂均为o(n)
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

//[7,1,5,3,6,4]-->4+3=7
//买股票的最大收益：数组元素为顺序的每天股价，由左往右，买入后才能卖出，买入前必须卖出

//思路：买入卖出可看成一个区间，每个区间并无交集，比如X1...Xm-1 Xm Xm+1...Xk...Xp  假设Xp最大，Xm是中间变小的数，X1到Xm-1一直是递增的，Xm+1到Xp是递增的，这符合普遍规律
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

//public void rotate(int[] nums, int k) {
//reverse(nums,nums.length-k%nums.length,nums.length-1);
//reverse(nums,0,nums.length-k%nums.length-1);
//reverse(nums,0,nums.length-1);
//}
//public void reverse(int[] nums,int start,int end){
//while(start<end){
//int temp = nums[start];
//nums[start++] = nums[end];
//nums[end--] = temp;
//}
//}

func Rotate(nums []int) int {
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
func reverse(nums []int) {
	for i := 0; i < len(nums); i++ {
		//nums[len(nums)/2-1]
	}
}
