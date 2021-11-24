package listnode

// Create a pointer of ListNode from a slice contains integer.
//
// Refers to the problem Remove Linked List Elements: https://leetcode.com/problems/remove-linked-list-elements/.
//
// We can follow the code below to create a ListNode from the slice:
//
//     listnode.FromSlice([]int{1, 2, 3, 4}) // *listnode.ListNode
func FromSlice(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	head := &ListNode{
		Val: s[0],
	}
	current := head
	for i := 1; i < len(s); i++ {
		current.Next = &ListNode{
			Val: s[i],
		}
		current = current.Next
	}
	return head
}
