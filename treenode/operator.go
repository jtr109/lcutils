package treenode

import "github.com/jtr109/lcutils/nilint"

type Operator struct {
	root *TreeNode
}

func NewOperator() *Operator {
	return &Operator{}
}

func (op *Operator) SetRoot(root *TreeNode) {
	op.root = root
}

func (op *Operator) Root() *TreeNode {
	return op.root
}

func (op *Operator) FromSlice(s []nilInt) {
	op.root = nil
	if len(s) == 0 {
		return
	}
	op.root = NewTreeNode(&s[0])
	nodeQueue := []*TreeNode{op.root}
	i := 0
	for len(nodeQueue) > 0 && i < len(s) {
		current := nodeQueue[0]
		var leftChild *TreeNode
		var rightChild *TreeNode
		i++
		if i < len(s) && !s[i].IsNil() {
			leftChild = NewTreeNode(&s[i])
			nodeQueue = append(nodeQueue, leftChild)
		}
		i++
		if i < len(s) && !s[i].IsNil() {
			rightChild = NewTreeNode(&s[i])
			nodeQueue = append(nodeQueue, rightChild)
		}
		current.Left = leftChild
		current.Right = rightChild
		nodeQueue = nodeQueue[1:] // left pop
	}
}

func (op *Operator) ToSlice() (result []nilInt) {
	// maybe two possible mode
	// 1. all nil in last level is tripped
	// 2. all tail nils is tripped
	// more examples are required for a decision
	// we choose the second option for now
	if op.Root() == nil {
		return []nilInt{}
	}
	nextLevel := []*TreeNode{op.Root()}
	for len(nextLevel) != 0 {
		current := nextLevel
		nextLevel = []*TreeNode{}
		// // stop convertion if all nodes in current level is nil
		// allNil := true
		// for _, node := range current {
		// 	if node != nil {
		// 		allNil = false
		// 		break
		// 	}
		// }
		// if allNil {
		// 	break
		// }
		for _, node := range current {
			if node == nil {
				result = append(result, nilint.NewNil())
				continue
			}
			result = append(result, nilint.NewInt(node.Val))
			nextLevel = append(nextLevel, node.Left, node.Right)
		}
	}
	for len(result) > 0 && result[len(result)-1] == nilint.NewNil() {
		result = result[:len(result)-1]
	}
	return
}
