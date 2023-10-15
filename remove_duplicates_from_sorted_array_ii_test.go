package leetgo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates2(t *testing.T) {
	testCases := []struct {
		nums         []int
		k            int
		expectedNums []int
	}{
		{
			nums:         []int{1, 1, 1, 2, 2, 3},
			k:            5,
			expectedNums: []int{1, 1, 2, 2, 3, 3},
		},
		{
			nums:         []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			k:            7,
			expectedNums: []int{0, 0, 1, 1, 2, 3, 3, 3, 3},
		},
	}

	for _, testCase := range testCases {
		k, nums := RemoveDuplicates2(testCase.nums)
		require.Equal(t, testCase.k, k)
		require.Equal(t, testCase.expectedNums[:testCase.k], nums[:k])
	}
}

func RemoveDuplicates2(nums []int) (int, []int) {
	k := 1

	freq := 1
	for i := 1; i < len(nums); i++ {
		if freq == 2 && nums[i] == nums[i-1] {
			continue
		} else if freq < 2 && nums[i] == nums[i-1] {
			nums[k] = nums[i]
			freq++
			k++
		} else {
			nums[k] = nums[i]
			freq = 1
			k++
		}
	}

	return k, nums
}
