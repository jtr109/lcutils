package treenode

func Equal(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	if !treeNodeEqual(node1.Left, node2.Left) || !treeNodeEqual(node1.Right, node2.Right) {
		return false
	}
	return true
}
