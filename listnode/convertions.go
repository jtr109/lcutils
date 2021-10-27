package listnode

func ConvertListNodeToArray(head *ListNode) []int {
	current := head
	result := []int{}
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func ConvertArrayToListNode(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}
	head := &ListNode{
		Val: array[0],
	}
	current := head
	for i := 1; i < len(array); i++ {
		current.Next = &ListNode{
			Val: array[i],
		}
		current = current.Next
	}
	return head
}
