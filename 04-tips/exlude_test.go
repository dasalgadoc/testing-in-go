//go:build exclude
// +build exclude

package _4_tips

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExclude(t *testing.T) {
	t.Parallel()

	assert.True(t, false)
}
