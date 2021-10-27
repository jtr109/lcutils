// Package listnode provides a set of tools for writing LeetCode problems locally.
//
// Example Usage
//
// ```go
// package example
//
// import (
// 	"testing"
//
// 	"github.com/jtr109/lcutils/listnode"
// 	"github.com/stretchr/testify/assert"
// )
//
// func main() {
//   head := []int{1, 2, 3, 4}
//   node := listnode.ConvertArrayToListNode(head) // *listnode.ListNode
//   newHead := listnode.ConvertListNodeToArray(node) // []int{1, 2, 3, 4}
// }
// ```
package listnode
