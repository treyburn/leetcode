package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_firstBadVersion(t *testing.T) {
	input := 22
	WANT = 11

	got := firstBadVersion(input)

	assert.Equal(t, WANT, got)
}
