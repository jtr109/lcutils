package treenode

import (
	"testing"

	"github.com/jtr109/lcutils/nilint"
	"github.com/stretchr/testify/assert"
)

var newInt = nilint.NewInt
var newNil = nilint.NewNil

func TestFromSlice(t *testing.T) {
	slice := []nilInt{newInt(3), newInt(9), newInt(20), newNil(), newNil(), newInt(15), newInt(7)}
	op := NewOperator()
	op.FromSlice(slice)
	assert.Equal(t, 3, op.Root().Val)
	assert.Equal(t, 9, op.Root().Left.Val)
	assert.Nil(t, op.Root().Left.Left)
	assert.Nil(t, op.Root().Left.Right)
	assert.Equal(t, 20, op.Root().Right.Val)
	assert.Equal(t, 15, op.Root().Right.Left.Val)
	assert.Nil(t, op.Root().Right.Left.Left)
	assert.Nil(t, op.Root().Right.Left.Right)
	assert.Equal(t, 7, op.Root().Right.Right.Val)
	assert.Nil(t, op.Root().Right.Right.Left)
	assert.Nil(t, op.Root().Right.Right.Right)
}

func TestFromEmptySlice(t *testing.T) {
	slice := []nilInt{}
	op := NewOperator()
	op.FromSlice(slice)
	assert.Nil(t, op.Root())
}

func TestFromPartialSlice(t *testing.T) {
	slice := []nilInt{newInt(1), newInt(2)}
	op := NewOperator()
	op.FromSlice(slice)
	assert.Equal(t, 1, op.Root().Val)
	assert.Equal(t, 2, op.Root().Left.Val)
	assert.Nil(t, op.Root().Right)
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
	slice := []nilInt{
		newInt(1),
		newInt(2),
		newNil(),
		newInt(3),
		newNil(),
		newInt(4),
		newNil(),
		newInt(5),
	}
	op := NewOperator()
	op.FromSlice(slice)
	assert.Equal(t, 1, op.Root().Val)
	assert.Equal(t, 2, op.Root().Left.Val)
	assert.Nil(t, op.Root().Right)
	assert.Equal(t, 3, op.Root().Left.Left.Val)
	assert.Nil(t, op.Root().Left.Right)
	assert.Equal(t, 4, op.Root().Left.Left.Left.Val)
	assert.Nil(t, op.Root().Left.Left.Right)
	assert.Equal(t, 5, op.Root().Left.Left.Left.Left.Val)
	assert.Nil(t, op.Root().Left.Left.Left.Left.Left)
	assert.Nil(t, op.Root().Left.Left.Left.Left.Right)
}

func TestNilToSlice(t *testing.T) {
	assert.Equal(t, []nilInt{}, NewOperator().ToSlice())
}

func TestToSlice(t *testing.T) {
	//         3
	//       /   \
	//      /     \
	//     9       20
	//    /  \     / \
	//   /    \   /   \
	// null  null 15   7
	slice := []nilInt{newInt(3), newInt(9), newInt(20), newNil(), newNil(), newInt(15), newInt(7)}
	op := NewOperator()
	op.FromSlice(slice)
	assert.Equal(t, slice, op.ToSlice())
}

func TestOnlyLeftToSlice(t *testing.T) {
	//         1
	//        /
	//       2
	//      /
	//     3
	//    /
	//   4
	//  /
	// 5
	slice := []nilInt{
		newInt(1),
		newInt(2),
		newNil(),
		newInt(3),
		newNil(),
		newInt(4),
		newNil(),
		newInt(5),
	}
	op := NewOperator()
	op.FromSlice(slice)
	assert.Equal(t, slice, op.ToSlice())
}
