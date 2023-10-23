package leetgo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		nums         []int
		k            int
		expectedNums []int
	}{
		{
			nums:         []int{1, 2, 3, 4, 5, 6, 7},
			k:            3,
			expectedNums: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums:         []int{-1, -100, 3, 99},
			k:            2,
			expectedNums: []int{3, 99, -1, -100},
		},
	}

	for _, testCase := range testCases {
		nums := Rotate(testCase.nums, testCase.k)
		require.Equal(t, testCase.expectedNums, nums)
	}
}

func Rotate(nums []int, k int) []int {
	k %= len(nums)
	Reverse(nums, 0, len(nums)-1)
	Reverse(nums, 0, k-1)
	Reverse(nums, k, len(nums)-1)
	return nums
}

func Reverse(nums []int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
