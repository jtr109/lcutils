package listnode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSlice(t *testing.T) {
	head := FromSlice([]int{1, 2, 3, 4})
	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	assert.Equal(t, expected, head)
}

func TestFromEmptySlice(t *testing.T) {
	head := FromSlice([]int{})
	assert.Nil(t, head)
}
