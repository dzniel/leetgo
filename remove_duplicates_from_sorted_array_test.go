package leetgo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums         []int
		k            int
		expectedNums []int
	}{
		{
			nums:         []int{1, 1, 2},
			k:            2,
			expectedNums: []int{1, 2, 2},
		},
		{
			nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			k:            5,
			expectedNums: []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4},
		},
	}

	for _, testCase := range testCases {
		k, nums := RemoveDuplicates(testCase.nums)
		require.Equal(t, testCase.k, k)
		require.Equal(t, testCase.expectedNums[:testCase.k], nums[:k])
	}
}

func RemoveDuplicates(nums []int) (int, []int) {
	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k, nums
}
