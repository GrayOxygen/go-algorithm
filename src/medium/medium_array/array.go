package medium_array

import (
	"go-algorithm/src/util/math_util"
)

//问题：给定一个字符串，找到不含重复字符的最长子串
//
//解决：
func LengthOfLongestSubstring(s string) int {
	oldIndex := 0
	maxSubLen := 0
	words := []rune(s)
	subMap := make(map[string]int, 0)
	for i := 0; i < len(words); i++ {
		if _, ok := subMap[string(words[i])]; !ok {
			subMap[ string(words[i])] = i
			maxSubLen = math_util.Max(maxSubLen, len(subMap))
			continue
		}
		delete(subMap, string(words[i])) //遇到重复时，即子字符串终止，移除元素从下一个元素开始统计子串
		oldIndex++
	}
	return maxSubLen
}
