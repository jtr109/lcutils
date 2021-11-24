// ## Example
//
// ## Create TreeNode from Slice
//
// For example, a `root` is provided as a abstract slice which contains level scan result of a binary tree, refers to the problem [Maximum Depth of Binary Tree - LeetCode](https://leetcode.com/problems/maximum-depth-of-binary-tree/).
//
// We can follow the code below to create a `TreeNode` of `root`:
//
// ```go
// root := treenode.NewOperator().FromSlice([]nilint.NilInt{
// 		nilint.NewInt(3),
// 		nilint.NewInt(9),
// 		nilint.NewInt(20),
// 		nilint.NewNil(),
// 		nilint.NewNil(),
// 		nilint.NewInt(15),
// 		nilint.NewInt(7),
// 	}).Root()
// ```
//
// Explanation:
//
// We cannot use `[]int` as the type of the slice `root` because it contains `null`. We define a type `nilint.NilInt` to represent it, and now we can define the slice as type `[]nilint.NilInt`. We provide function `NewInt` and `NewNil` to create integer and null value.
//
// The function `treenode.NewOperator` returns an instance of type `treenode.Operator` with a lot of features. But in our case, we only need to get the root `TreeNode` we need from the method `Root`.
package treenode
