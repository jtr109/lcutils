// Helper utils for the type TreeNode in LeetCode.
//
// Decleration of TreeNode
//
// The TreeNode in this package can be used to declare the one in you package:
//
//     type TreeNode = treenode.TreeNode
//
//     func foo(root *TreeNode) {}
//
// Convertion From Slice
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
//
// Comparing Two TreeNode
//
// Assume we want to check if two TreeNode is equal, use the following example:
//
//     p := treenode.FromSlice([]nilint.NilInt{
//         nilint.NewInt(1),
//         nilint.NewInt(2),
//         nilint.NewInt(3),
//     })
//     q := treenode.FromSlice([]nilint.NilInt{
//         nilint.NewInt(1),
//         nilint.NewInt(2),
//         nilint.NewInt(3),
//     })
//     Equal(p, q) // true
package treenode
