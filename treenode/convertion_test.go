package treenode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSlice(t *testing.T) {
	root := FromSlice([]nilInt{newInt(3), newInt(9), newInt(20), newNil(), newNil(), newInt(15), newInt(7)})
	assert.Equal(t, 3, root.Val)
	assert.Equal(t, 9, root.Left.Val)
	assert.Nil(t, root.Left.Left)
	assert.Nil(t, root.Left.Right)
	assert.Equal(t, 20, root.Right.Val)
	assert.Equal(t, 15, root.Right.Left.Val)
	assert.Nil(t, root.Right.Left.Left)
	assert.Nil(t, root.Right.Left.Right)
	assert.Equal(t, 7, root.Right.Right.Val)
	assert.Nil(t, root.Right.Right.Left)
	assert.Nil(t, root.Right.Right.Right)
}

func TestFromEmptySlice(t *testing.T) {
	root := FromSlice([]nilInt{})
	assert.Nil(t, root)
}

func TestFromPartialSlice(t *testing.T) {
	root := FromSlice([]nilInt{newInt(1), newInt(2)})
	assert.Equal(t, 1, root.Val)
	assert.Equal(t, 2, root.Left.Val)
	assert.Nil(t, root.Right)
}

func TestFromOnlyLeftSlice(t *testing.T) {
	//         1
	//        /
	//       2
	//      /
	//     3
	//    /
	//   4
	//  /
	// 5
	root := FromSlice([]nilInt{
		newInt(1),
		newInt(2),
		newNil(),
		newInt(3),
		newNil(),
		newInt(4),
		newNil(),
		newInt(5),
	})
	assert.Equal(t, 1, root.Val)
	assert.Equal(t, 2, root.Left.Val)
	assert.Nil(t, root.Right)
	assert.Equal(t, 3, root.Left.Left.Val)
	assert.Nil(t, root.Left.Right)
	assert.Equal(t, 4, root.Left.Left.Left.Val)
	assert.Nil(t, root.Left.Left.Right)
	assert.Equal(t, 5, root.Left.Left.Left.Left.Val)
	assert.Nil(t, root.Left.Left.Left.Left.Left)
	assert.Nil(t, root.Left.Left.Left.Left.Right)
}
