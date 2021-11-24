package listnode

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
