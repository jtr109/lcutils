package treenode

import "github.com/jtr109/lcutils/nilint"

func fromNilInt(val *nilInt) *TreeNode {
	if val.IsNil() {
		return nil
	}
	return &TreeNode{
		Val: val.Val(),
	}
}

func FromSlice(s []nilint.NilInt) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	root := fromNilInt(&s[0])
	nodeQueue := []*TreeNode{root}
	i := 0
	for len(nodeQueue) > 0 && i < len(s) {
		current := nodeQueue[0]
		nodeQueue = nodeQueue[1:] // left pop
		var leftChild *TreeNode
		var rightChild *TreeNode
		i++
		if i < len(s) && !s[i].IsNil() {
			leftChild = fromNilInt(&s[i])
			nodeQueue = append(nodeQueue, leftChild)
		}
		i++
		if i < len(s) && !s[i].IsNil() {
			rightChild = fromNilInt(&s[i])
			nodeQueue = append(nodeQueue, rightChild)
		}
		current.Left = leftChild
		current.Right = rightChild
	}
	return root
}
