package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayRemoved(t *testing.T) {
	for i, v := range []struct {
		arr      []int
		idx      int
		expected []int
	}{
		{
			arr:      []int{1, 2, 3, 4, 5},
			idx:      0,
			expected: []int{2, 3, 4, 5},
		},
		{
			arr:      []int{1, 2, 3, 4, 5},
			idx:      1,
			expected: []int{1, 3, 4, 5},
		},
		{
			arr:      []int{1, 2, 3, 4, 5},
			idx:      2,
			expected: []int{1, 2, 4, 5},
		},
		{
			arr:      []int{1, 2, 3, 4, 5},
			idx:      3,
			expected: []int{1, 2, 3, 5},
		},
		{
			arr:      []int{1, 2, 3, 4, 5},
			idx:      4,
			expected: []int{1, 2, 3, 4},
		},
	} {
		testCase := v
		t.Run(fmt.Sprintf("testcase-%d", i), func(t *testing.T) {
			assert.Equal(t, testCase.expected, arrayExcludingIndex(testCase.arr, testCase.idx))
		})
	}
}
