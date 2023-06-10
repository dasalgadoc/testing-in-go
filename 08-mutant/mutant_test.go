package _8_mutant

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompleteCoverageButNoMutantProof(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
	}{
		{"1 + 1", 1, 1},
		{"2 + 2", 2, 2},
		{"3 + 3", 3, 3},
		{"4 + 4", 4, 4},
		{"5 + 5", 5, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sumFunction(test.a, test.b)
		})
	}
}

func TestMutantTest_ThisTestIsIncorrect(t *testing.T) {
	sumFunction = func(a, b int) int {
		return a - b
	}
	defer restoreSumFunction()

	tests := []struct {
		name string
		a    int
		b    int
	}{
		{"1 + 1", 1, 1},
		{"2 + 2", 2, 2},
		{"3 + 3", 3, 3},
		{"4 + 4", 4, 4},
		{"5 + 5", 5, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sumFunction(test.a, test.b)
		})
	}
}

func TestTwoNumbersShouldBeAddingInSumFunction(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"1 + 1 = 2", 1, 1, 2},
		{"2 + 2 = 4", 2, 2, 4},
		{"3 + 3 = 6", 3, 3, 6},
		{"4 + 4 = 8", 4, 4, 8},
		{"5 + 5 = 10", 5, 5, 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sumFunction(test.a, test.b)
			assert.Equal(t, test.want, got,
				"should be equal got %d want %d in %s test", got, test.want, test.name)
		})
	}
}

func restoreSumFunction() {
	sumFunction = func(a, b int) int {
		return a + b
	}
}
