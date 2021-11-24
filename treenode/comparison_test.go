package treenode

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	root1 := FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(2),
		nilint.NewInt(3),
	})
	root2 := FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(2),
		nilint.NewInt(3),
	})
	assert.True(t, Equal(root1, root2))
}

func TestNotEqual(t *testing.T) {
	root1 := FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(2),
	})
	root2 := FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewNil(),
		nilint.NewInt(2),
	})
	assert.False(t, Equal(root1, root2))
}

func TestNotEqual2(t *testing.T) {
	root1 := FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(2),
		nilint.NewInt(1),
	})
	root2 := FromSlice([]nilint.NilInt{
		nilint.NewInt(1),
		nilint.NewInt(1),
		nilint.NewInt(2),
	})
	assert.False(t, Equal(root1, root2))
}
