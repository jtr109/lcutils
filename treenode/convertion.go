package treenode

import "github.com/jtr109/lcutils/nilint"

func fromNilInt(val *nilint.NilInt) *TreeNode {
	if val.IsNil() {
		return nil
	}
	return &TreeNode{
		Val: val.Val(),
	}
}

// Create a TreeNode from a slice.
//
// For example, a root is provided as a abstract slice which contains level scan result of a binary tree, refers to the problem Maximum Depth of Binary Tree: https://leetcode.com/problems/maximum-depth-of-binary-tree/.
//
// We can follow the code below to create a TreeNode of root:
//
//     root := treenode.FromSlice([]nilint.NilInt{
//         nilint.NewInt(3),
//         nilint.NewInt(9),
//         nilint.NewInt(20),
//         nilint.NewNil(),
//         nilint.NewNil(),
//         nilint.NewInt(15),
//         nilint.NewInt(7),
//     })
//
// Explanation:
//
// We cannot use []int as the type of the slice root because it contains null. We define a type nilint.NilInt to represent it, and now we can define the slice as type []nilint.NilInt. We provide function NewInt and NewNil to create integer and null value.
//
// The function treenode.NewOperator returns an instance of type treenode.Operator with a lot of features. But in our case, we only need to get the root TreeNode we need from the method Root.
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

// Convert the TreeNode into a slice.
//
// For example, we have the root of a tree with type TreeNode, we can make a convertion like below:
//
//    ToSlice(root)
//
// The output will be a slice. As is shown in all LeetCode problems, such as https://leetcode.com/problems/maximum-depth-of-binary-tree/.
func ToSlice(root *TreeNode) []nilint.NilInt {
	slice := []nilint.NilInt{}
	nextLevel := []*TreeNode{root}
	for len(nextLevel) > 0 {
		currentLevel := nextLevel
		nextLevel = []*TreeNode{}
		for _, node := range currentLevel {
			if node == nil {
				slice = append(slice, nilint.NewNil())
				continue
			}
			slice = append(slice, nilint.NewInt(node.Val))
			nextLevel = append(nextLevel, node.Left, node.Right)
		}
	}
	// strip tail nils
	length := 0
	for i, n := range slice {
		if !n.IsNil() {
			length = i + 1
		}
	}
	return slice[:length]
}
