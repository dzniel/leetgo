package leetgo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanJump(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		canJump := CanJump(testCase.nums)
		require.Equal(t, testCase.expected, canJump)
	}
}

func CanJump(nums []int) bool {
	position := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= position {
			position = i
		}
	}

	return position == 0
}
