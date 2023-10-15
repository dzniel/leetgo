package leetgo

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
	}

	for index, testCase := range testCases {
		result := Merge(testCase.nums1, testCase.m, testCase.nums2, testCase.n)
		require.Equal(t, testCase.expected, result, index+1)
	}
}

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	if len(nums1) > m {
		nums1 = nums1[:m]
	}

	if len(nums2) > n {
		nums2 = nums2[:n]
	}

	nums1 = append(nums1, nums2...)

	sort.Ints(nums1)
	return nums1
}
