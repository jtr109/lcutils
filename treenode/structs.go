package treenode

import "github.com/jtr109/lcutils/nilint"

type nilInt = nilint.NilInt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val *nilInt) *TreeNode {
	if val.IsNil() {
		return nil
	}
	return &TreeNode{
		Val: val.Val(),
	}
}
