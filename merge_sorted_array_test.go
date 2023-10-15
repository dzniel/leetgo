package leetgo

import (
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
	index := m + n - 1

	i := m - 1
	j := n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[index] = nums1[i]
			index--
			i--
		} else {
			nums1[index] = nums2[j]
			index--
			j--
		}
	}

	for j >= 0 {
		nums1[index] = nums2[j]
		index--
		j--
	}

	return nums1
}
