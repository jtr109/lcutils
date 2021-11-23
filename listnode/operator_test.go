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
	op := NewOperator().FromSlice([]int{3, 2, 0, -4})
	node1, err := op.Get(1)
	assert.Nil(t, err)
	node3, err := op.Get(3)
	assert.Nil(t, err)
	assert.Nil(t, op.CycleBegin()) // return nil if the list node doesn't have cycle
	node3.Next = node1
	assert.Equal(t, node1, op.CycleBegin())
	// Appending failed when the node list has a cycle
	node4 := NewListNode(newRand().Int())
	assert.Errorf(t, op.Append(node4), "appending failed when the node list has a cycle")
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

func TestOperatorGet(t *testing.T) {
	op := NewOperator()
	head := NewListNode(3)
	node1 := NewListNode(2)
	op.Append(head)
	op.Append(node1)
	node, err := op.Get(0)
	assert.Equal(t, head, node)
	assert.Nil(t, err)
	node, err = op.Get(1)
	assert.Equal(t, node1, node)
	assert.Nil(t, err)
	for i := -1; i > -3; i-- {
		_, err := op.Get(i)
		assert.Errorf(t, err, fmt.Sprintf("invalid index %d (index must be non-negative", i))
	}
	for i := 2; i < 10; i++ {
		_, err := op.Get(i)
		assert.Errorf(t, err, fmt.Sprintf("index %d out of range", i))
	}
}

func TestOperatorGetPrevious(t *testing.T) {
	op := NewOperator()
	head := NewListNode(newRand().Int())
	node1 := NewListNode(newRand().Int())
	op.Append(head)
	op.Append(node1)
	prev, err := op.getPrevious(0)
	assert.Equal(t, op.virtualHead, prev)
	assert.Nil(t, err)
	prev, err = op.getPrevious(1)
	assert.Equal(t, head, prev)
	assert.Nil(t, err)
	prev, err = op.getPrevious(2)
	assert.Equal(t, node1, prev)
	assert.Nil(t, err)
	for i := -1; i > -3; i-- {
		prev, err = op.getPrevious(i)
		assert.Nil(t, prev)
		assert.Errorf(t, err, fmt.Sprintf("invalid index %d (index must be non-negative", i))
	}
	for i := 3; i < 10; i++ {
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
	op := NewOperator().FromSlice([]int{3, 2, 0, -4})
	node1, _ := op.Get(1)
	node2, _ := op.Get(2)
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
	op := NewOperator().FromSlice([]int{3, 2})
	node0, _ := op.Get(0)
	node2 := NewListNode(1)
	op.Insert(0, node2)
	assert.Equal(t, node2, op.Head())
	assert.Equal(t, node0, node2.Next)
}

func TestOperatorInsertAtTail(t *testing.T) {
	// 3 --> 2
	//        \
	//         '--> 1
	op := NewOperator().FromSlice([]int{3, 2})
	node1, _ := op.Get(1)
	node2 := NewListNode(1)
	op.Insert(2, node2)
	assert.Equal(t, node2, node1.Next)
	assert.Nil(t, node2.Next)
}

func TestOperatorInsertWithInvalidIndex(t *testing.T) {
	node2 := NewListNode(1)
	op := NewOperator().FromSlice([]int{3, 2})
	for i := -1; i > -3; i-- {
		err := op.Insert(i, node2)
		assert.Errorf(t, err, fmt.Sprintf("invalid index %d (index must be non-negative", i))
	}
	for i := 3; i < 10; i++ {
		err := op.Insert(i, node2)
		assert.Errorf(t, err, fmt.Sprintf("index %d out of range", i))
	}

}

func TestOperatorDelete(t *testing.T) {
	// 3 ----> 2 - - x - -> 0 - - x - -> -4
	//          \                        ^
	//           \                      /
	//            '--------------------'
	op := NewOperator().FromSlice([]int{3, 2, 0, -4})
	node1, _ := op.Get(1)
	node2, _ := op.Get(2)
	node3, _ := op.Get(3)
	err := op.Delete(2)
	assert.Nil(t, err)
	assert.Equal(t, node3, node1.Next)
	assert.Nil(t, node2.Next)
}

func TestOperatorDeleteHead(t *testing.T) {
	op := NewOperator()
	head := NewListNode(3)
	node1 := NewListNode(2)
	op.Append(head)
	op.Append(node1)
	assert.Equal(t, head, op.Head())
	err := op.Delete(0)
	assert.Nil(t, err)
	assert.Equal(t, node1, op.Head())
}

func TestOperatorDeleteTail(t *testing.T) {
	op := NewOperator()
	head := NewListNode(3)
	node1 := NewListNode(2)
	op.Append(head)
	op.Append(node1)
	assert.Equal(t, head, op.Head())
	err := op.Delete(1)
	assert.Nil(t, err)
	assert.Nil(t, head.Next)
	assert.Nil(t, node1.Next)
}

func TestOperatorDeleteWithInvalidIndex(t *testing.T) {
	op := NewOperator()
	head := NewListNode(3)
	node1 := NewListNode(2)
	op.Append(head)
	op.Append(node1)
	for i := -1; i > -3; i-- {
		err := op.Delete(i)
		assert.Errorf(t, err, fmt.Sprintf("invalid index %d (index must be non-negative", i))
	}
	for i := 2; i < 10; i++ {
		err := op.Delete(i)
		assert.Errorf(t, err, fmt.Sprintf("index %d out of range", i))
	}
}

func TestOperatorToSlice(t *testing.T) {
	s := []int{3, 2, 0, -4}
	op := NewOperator().FromSlice(s)
	newS, err := op.ToSlice()
	assert.Nil(t, err)
	assert.Equal(t, s, newS)
}

func TestOperatorToSliceWithCycle(t *testing.T) {
	// 3 -> 2 -> 0 -> -4
	//       ^        /
	//        \      /
	//         '----'
	s := []int{3, 2, 0, -4}
	op := NewOperator().FromSlice(s)
	node1, _ := op.Get(1)
	node3, _ := op.Get(3)
	node3.Next = node1
	newS, err := op.ToSlice()
	assert.Errorf(t, err, "failed because the node list has a cycle")
	assert.Nil(t, newS)
}
