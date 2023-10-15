package leetgo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMajorityElement(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
	}

	for _, testCase := range testCases {
		result := MajorityElement(testCase.nums)
		require.Equal(t, testCase.expected, result)
	}
}

func MajorityElement(nums []int) int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
		if value, ok := freq[num]; ok && value > (len(nums)/2) {
			return num
		}
	}

	return 0
}
