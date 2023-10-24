package leetgo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxProfit2(t *testing.T) {
	testCases := []struct {
		prices   []int
		expected int
	}{
		{
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, testCase := range testCases {
		maxProfit := MaxProfit2(testCase.prices)
		require.Equal(t, testCase.expected, maxProfit)
	}
}

func MaxProfit2(prices []int) int {
	maxProfit := 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			maxProfit += (prices[i+1] - prices[i])
		}
	}

	return maxProfit
}
