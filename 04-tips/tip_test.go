package _4_tips

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssertionEquals(t *testing.T) {
	t.Parallel()

	theShining := NewBookPointer("The Shining")
	tS := NewBookPointer("The Shining")

	assert.Equal(t, theShining, tS)
	assert.NotSame(t, theShining, tS)
}

func TestAssertionEqualsWithPointerConstructor(t *testing.T) {
	t.Parallel()

	theShining := NewBookPointer("The Shining")
	tS := NewBookPointer("The Shining")

	assert.Equal(t, theShining, tS)
	assert.NotSame(t, theShining, tS)
}

func TestAssertionEqualsWithReferences(t *testing.T) {
	t.Parallel()

	theShining := booker
	theShining.Name = "The Shining"
	tS := booker

	assert.Equal(t, theShining, tS)
	assert.Same(t, theShining, tS)
}
