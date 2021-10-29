package listnode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewListNode(t *testing.T) {
	val := newRand().Int()
	node := NewListNode(val)
	assert.Equal(t, val, node.Val)
	assert.Nil(t, node.Next)
}
