package _1_basics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_All_Starts_Well(t *testing.T) {
	assert.True(t, true, "True is true!")
}

func Test_Any_Number_Should_Return_Itself(t *testing.T) {
	assert.Equal(t, 1, FizzBuzz(1))
	assert.Equal(t, 2, FizzBuzz(2))
	assert.Equal(t, 4, FizzBuzz(4))
}
