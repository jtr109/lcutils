package listnode

// The ListNode in this package can be used to declare the one in you package:
//
//     type ListNode = listnode.ListNode // type alias
//
//     func foo(head *ListNode) {}
type ListNode struct {
	Val  int
	Next *ListNode
}
