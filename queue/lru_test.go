package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LRUCache(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	assert.Equal(t, 1, len(lru.set))

	lru.Put(2, 2)
	assert.Equal(t, 2, len(lru.set))

	got := lru.Get(1)
	assert.Equal(t, 1, got)

	lru.Put(3, 3)
	assert.Equal(t, 2, len(lru.set))

	got = lru.Get(2)
	assert.Equal(t, -1, got)

	lru.Put(4, 4)
	assert.Equal(t, 2, len(lru.set))

	got = lru.Get(1)
	assert.Equal(t, -1, got)

	got = lru.Get(3)
	assert.Equal(t, 3, got)

	got = lru.Get(4)
	assert.Equal(t, 4, got)

	lru.Put(3, 4)

	lru.Put(5, 1)

	got = lru.Get(3)
	assert.Equal(t, 4, got)

	got = lru.Get(3)
	assert.Equal(t, 4, got)

	got = lru.Get(3)
	assert.Equal(t, 4, got)

	got = lru.Get(3)
	assert.Equal(t, 4, got)

	lru.Put(3, 5)

	got = lru.Get(3)
	assert.Equal(t, 5, got)

	got = lru.Get(5)
	assert.Equal(t, 1, got)
}
