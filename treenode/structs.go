package treenode

// The TreeNode in this package can be used to declare the one in you package:
//
//     type TreeNode = treenode.TreeNode // type alias
//
//     func foo(root *TreeNode) {}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
