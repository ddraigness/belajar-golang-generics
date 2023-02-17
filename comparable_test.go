package belajargolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T) {
	assert.Equal(t, true, IsSame[string]("surya", "surya"))
	assert.Equal(t, true, IsSame[int](100, 100))
	assert.Equal(t, true, IsSame[bool](true, true))
}
