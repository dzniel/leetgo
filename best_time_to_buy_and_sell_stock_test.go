package leetgo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		prices   []int
		expected int
	}{
		{
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, testCase := range testCases {
		maxProfit := MaxProfit(testCase.prices)
		require.Equal(t, testCase.expected, maxProfit)
	}
}

func MaxProfit(prices []int) int {
	lowestPrice := 10001
	maxProfit := 0

	for _, price := range prices {
		if price < lowestPrice {
			lowestPrice = price
		}

		currentProfit := price - lowestPrice

		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
	}

	return maxProfit
}
