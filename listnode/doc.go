// Package listnode provides a set of tools for writing LeetCode problems locally.
//
// Example Usage
//
// ```go
// package example
//
// import "github.com/jtr109/lcutils/listnode"
//
// type ListNode = listnode.ListNode // type alias
//
// func main() {
//   head := []int{1, 2, 3, 4}
//   node := listnode.ConvertArrayToListNode(head) // *listnode.ListNode
//   newHead := listnode.ConvertListNodeToArray(node) // []int{1, 2, 3, 4}
// }
// ```
package listnode
