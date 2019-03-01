package medium_array

import (
	"fmt"
	"go-algorithm/src/medium/medium_array"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	len := medium_array.LengthOfLongestSubstring("abbccde")
	fmt.Println("最大子串：", len)
}
