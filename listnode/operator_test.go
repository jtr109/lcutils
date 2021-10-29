package listnode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOperator(t *testing.T) {
	op := NewOperator()
	assert.Nil(t, op.virtualHead.Next)
}

func TestOperatorHead(t *testing.T) {
	op := NewOperator()
	head := NewListNode(newRand().Int())
	op.Append(head)
	assert.Equal(t, head, op.Head()) // node is appended as head when the node list is empty
	op.Append(NewListNode(newRand().Int()))
	assert.Equal(t, head, op.Head()) // head doesn't change after appending
}

func TestOperatorCycleBegin(t *testing.T) {
	// 3 -> 2 -> 0 -> -4
	//       ^        /
	//        \      /
	//         '----'
	op := NewOperator()
	head := NewListNode(3)
	node1 := NewListNode(2)
	node2 := NewListNode(0)
	node3 := NewListNode(-4)
	ok := op.Append(head)
	assert.True(t, ok)
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1
	assert.Equal(t, node1, op.CycleBegin())
	// Appending failed when the node list has a cycle
	node4 := NewListNode(newRand().Int())
	assert.False(t, op.Append(node4))
}

func TestOperatorAppend(t *testing.T) {
	op := NewOperator()
	head := NewListNode(newRand().Int())
	op.Append(head)
	node1 := NewListNode(newRand().Int())
	op.Append(node1)
	assert.Equal(t, node1, head.Next)
	node2 := NewListNode(newRand().Int())
	op.Append(node2)
	assert.Equal(t, node1, head.Next)
	assert.Equal(t, node2, head.Next.Next)
}

func TestOperatorGetPrevious(t *testing.T) {
	op := NewOperator()
	head := NewListNode(newRand().Int())
	node1 := NewListNode(newRand().Int())
	node2 := NewListNode(newRand().Int())
	op.Append(head)
	op.Append(node1)
	op.Append(node2)
	prev, err := op.getPrevious(0)
	assert.Equal(t, op.virtualHead, prev)
	assert.Nil(t, err)
	prev, err = op.getPrevious(1)
	assert.Equal(t, head, prev)
	assert.Nil(t, err)
	prev, err = op.getPrevious(2)
	assert.Equal(t, node1, prev)
	assert.Nil(t, err)
	prev, err = op.getPrevious(3)
	assert.Equal(t, node2, prev)
	assert.Nil(t, err)
	for i := -1; i > -3; i-- {
		prev, err = op.getPrevious(i)
		assert.Nil(t, prev)
		assert.Errorf(t, err, fmt.Sprintf("invalid index %d (index must be non-negative", i))
	}
	for i := 4; i < 10; i++ {
		prev, err = op.getPrevious(i)
		assert.Nil(t, prev)
		assert.Errorf(t, err, fmt.Sprintf("index %d out of range", i))
	}
}

func TestOperatorInsert(t *testing.T) {
	// 3 -> 2 - - x - -> 0 --> -4
	//       \          ^
	//        \        /
	//         '-> 1 -'
	op := NewOperator()
	head := NewListNode(3)
	node1 := NewListNode(2)
	node2 := NewListNode(0)
	node3 := NewListNode(-4)
	op.Append(head)
	op.Append(node1)
	op.Append(node2)
	op.Append(node3)
	node4 := NewListNode(1)
	op.Insert(2, node4)
	assert.Equal(t, node4, node1.Next)
	assert.Equal(t, node2, node4.Next)
}

func TestOperatorInsertAtHead(t *testing.T) {
	//       3 --> 2
	//       ^
	//      /
	// 1 --'
	op := NewOperator()
	head := NewListNode(3)
	node1 := NewListNode(2)
	op.Append(head)
	op.Append(node1)
	node2 := NewListNode(1)
	op.Insert(0, node2)
	assert.Equal(t, node2, op.Head())
}

func TestOperatorInsertAtTail(t *testing.T) {
	// 3 --> 2
	//        \
	//         '--> 1
	op := NewOperator()
	head := NewListNode(3)
	node1 := NewListNode(2)
	node2 := NewListNode(1)
	op.Append(head)
	op.Append(node1)
	op.Insert(2, node2)
	assert.Equal(t, node2, node1.Next)
	assert.Nil(t, node2.Next)
}

func TestOperatorInsertWithInvalidIndex(t *testing.T) {
	op := NewOperator()
	head := NewListNode(3)
	node1 := NewListNode(2)
	node2 := NewListNode(1)
	op.Append(head)
	op.Append(node1)
	for i := -1; i > -3; i-- {
		err := op.Insert(i, node2)
		assert.Errorf(t, err, fmt.Sprintf("invalid index %d (index must be non-negative", i))
	}
	for i := 3; i < 10; i++ {
		err := op.Insert(i, node2)
		assert.Errorf(t, err, fmt.Sprintf("index %d out of range", i))
	}

}
