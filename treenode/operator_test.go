package treenode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSlice(t *testing.T) {
	slice := []NilInt{NewInt(3), NewInt(9), NewInt(20), NewNil(), NewNil(), NewInt(15), NewInt(7)}
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
	slice := []NilInt{}
	op := NewOperator()
	op.FromSlice(slice)
	assert.Nil(t, op.Root())
}

func TestFromPartialSlice(t *testing.T) {
	slice := []NilInt{NewInt(1), NewInt(2)}
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
	slice := []NilInt{
		NewInt(1),
		NewInt(2),
		NewNil(),
		NewInt(3),
		NewNil(),
		NewInt(4),
		NewNil(),
		NewInt(5),
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
