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

func (op *Operator) FromSlice(s []int) *Operator {
	if len(s) == 0 {
		return nil
	}
	head := &ListNode{
		Val: s[0],
	}
	current := head
	for i := 1; i < len(s); i++ {
		current.Next = &ListNode{
			Val: s[i],
		}
		current = current.Next
	}
	op.virtualHead.Next = head
	return op
}

func (op *Operator) ToSlice() ([]int, error) {
	// failed if the node list has a cycle
	if op.cycleFlagNode() != nil {
		return nil, fmt.Errorf("failed because the node list has a cycle")
	}
	current := op.Head()
	result := []int{}
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result, nil
}

func (op *Operator) Head() *ListNode {
	return op.virtualHead.Next
}

func (op *Operator) CycleBegin() *ListNode {
	sign := op.cycleFlagNode()
	if sign == nil {
		// no cycle in the node list
		return nil
	}
	index1 := op.virtualHead
	index2 := sign
	for index1 != index2 {
		index1 = index1.Next
		index2 = index2.Next
	}
	return index1
}

func (op *Operator) cycleFlagNode() *ListNode {
	fast, slow := op.virtualHead, op.virtualHead
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return fast
		}
	}
	return nil

}

func (op *Operator) Append(node *ListNode) error {
	if op.cycleFlagNode() != nil {
		return fmt.Errorf("appending failed when the node list has a cycle")
	}

	tail := op.virtualHead
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = node
	return nil
}

func (op *Operator) Get(index int) (*ListNode, error) {
	if index < 0 {
		return nil, fmt.Errorf("invalid index %d (index must be non-negative", index)
	}
	current := op.virtualHead
	for i := 0; i <= index; i++ {
		current = current.Next
		if current == nil {
			return nil, fmt.Errorf("index %d out of range", index)
		}
	}
	return current, nil
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

func (op *Operator) Delete(index int) error {
	prev, err := op.getPrevious(index)
	if err != nil {
		return err
	}
	if prev.Next == nil {
		return fmt.Errorf("index %d out of range", index)
	}
	target := prev.Next
	prev.Next = prev.Next.Next
	target.Next = nil
	return nil
}
