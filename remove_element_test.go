package leetgo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		nums         []int
		val          int
		k            int
		expectedNums []int
	}{
		{
			nums:         []int{3, 2, 2, 3},
			val:          2,
			k:            2,
			expectedNums: []int{3, 3, 2, 3},
		},
		{
			nums:         []int{3, 2, 2, 3},
			val:          3,
			k:            2,
			expectedNums: []int{2, 2, 2, 3},
		},
		{
			nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:          2,
			k:            5,
			expectedNums: []int{0, 1, 3, 0, 4, 0, 4, 2},
		},
	}

	for _, testCase := range testCases {
		k, nums := RemoveElement(testCase.nums, testCase.val)
		require.Equal(t, testCase.k, k)
		require.Equal(t, testCase.expectedNums[:testCase.k], nums[:k])
	}
}

func RemoveElement(nums []int, val int) (int, []int) {
	k := 0

	for i := range nums {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k, nums
}
