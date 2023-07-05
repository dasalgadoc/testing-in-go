package _4_tips

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSkipThisTest(t *testing.T) {
	t.Parallel()
	t.Skip("This test is skipped")

	assert.True(t, false)
}

func TestNoSkipThisTest(t *testing.T) {
	t.Parallel()

	assert.True(t, true, "All true is true")
}
