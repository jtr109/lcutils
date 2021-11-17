package treenode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val *NilInt) *TreeNode {
	if val.isNil {
		return nil
	}
	return &TreeNode{
		Val: val.Val(),
	}
}
