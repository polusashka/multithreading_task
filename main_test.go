package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements_edgeSizes(t *testing.T) {
	edgeSizes := []int{0, -100}

	for _, size := range edgeSizes {
		result := generateRandomElements(size)
		require.Len(t, result, 0)
	}
}

func TestGenerateRandomElements_sizes(t *testing.T) {
	sizes := []int{1, 100, 10000}
	for _, size := range sizes {
		result := generateRandomElements(size)
		require.Len(t, result, size)
	}
}

func TestMaximum(t *testing.T) {
	testData := map[int][]int{
		0: {},
		5: {5},
		7: {7, 2, 5, 7, 3},
	}
	for i, v := range testData {
		result := maximum(v)
		require.Equal(t, result, i)
	}
}
