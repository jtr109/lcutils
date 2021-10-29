package listnode

import (
	"fmt"
)

type Operator struct {
	virtualHead *ListNode
}

func NewOperator() *Operator {
	return &Operator{
		virtualHead: &ListNode{},
	}
}

func (op *Operator) Head() *ListNode {
	return op.virtualHead.Next
}

func (op *Operator) CycleBegin() *ListNode {
	fast, slow := op.virtualHead, op.virtualHead
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			index1 := slow
			index2 := op.virtualHead
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index1
		}
	}
	return nil
}

func (op *Operator) Append(node *ListNode) (ok bool) {
	if op.CycleBegin() != nil {
		// New node cannot be appended because the node list has a cycle.
		return false
	}

	tail := op.virtualHead
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = node
	return true
}

func (op *Operator) getPrevious(index int) (*ListNode, error) {
	if index < 0 {
		return nil, fmt.Errorf("invalid index %d (index must be non-negative", index)
	}
	prev := op.virtualHead
	for i := 0; i < index && prev != nil; i++ {
		prev = prev.Next
		if prev == nil {
			return nil, fmt.Errorf("index %d out of range", index)
		}
	}
	return prev, nil
}

func (op *Operator) Insert(index int, node *ListNode) error {
	prev, err := op.getPrevious(index)
	if err != nil {
		return err
	}
	node.Next = prev.Next
	prev.Next = node
	return nil
}

// func (op *Operator) Delete(index int) {
// 	prev := op.getPrevious(index)
// 	prev.Next = prev.Next.Next
// }

// func (op *Operator) Get(index int) *ListNode {
// 	prev := op.getPrevious(index)
// 	return prev.Next
// }
