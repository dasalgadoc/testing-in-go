package _1_basics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_All_Starts_Well(t *testing.T) {
	assert.True(t, true, "True is true!")
}

func Test_Any_Number_Should_Return_Itself(t *testing.T) {
	assert.Equal(t, "1", FizzBuzz(1))
	assert.Equal(t, "2", FizzBuzz(2))
	assert.Equal(t, "4", FizzBuzz(4))
}

func Test_Return_Fizz_If_Divisible_By_3(t *testing.T) {
	assert.Equal(t, "Fizz", FizzBuzz(3))
	assert.Equal(t, "Fizz", FizzBuzz(6))
	assert.Equal(t, "Fizz", FizzBuzz(9))
}

func Test_Return_Buzz_If_Divisible_By_3(t *testing.T) {
	assert.Equal(t, "Buzz", FizzBuzz(5))
	assert.Equal(t, "Buzz", FizzBuzz(10))
	assert.Equal(t, "Buzz", FizzBuzz(20))
}
