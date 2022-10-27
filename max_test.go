package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {

	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{5, 2, -10, 4, 8, 9, -20, 5, 20, 4, 8},
			expected: 20,
		},
		{
			numbers:  []int{5, 2, -5, 20, 4, 8},
			expected: 20,
		},
		{
			numbers:  []int{5, 2, -10, 4, 8, 8},
			expected: 8,
		},
	}

	for _, testCase := range testTable {
		result := Max(testCase.numbers)
		t.Logf("Calling Max(%v), result %d\n", testCase.numbers, result)

		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrect result %d, got %d", testCase.expected, result))

	}

}

func TestMaxIndex(t *testing.T) {
	numbers := []int{5, 2, -10, 4, 8, 9, -20, 5, 20, 4, 8}
	expected := 8

	result := MaxIndex(numbers)

	if expected != result {
		t.Errorf("Incorrect result %d, got %d", expected, result)
	}
}
