package listnode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	head1 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				nil,
			},
		},
	}
	head2 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				nil,
			},
		},
	}
	assert.Equal(t, head1, head2)
}

func TestEqualListNodeWithCycle(t *testing.T) {
	head1 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				nil,
			},
		},
	}
	head1.Next.Next.Next = head1
	head2 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				nil,
			},
		},
	}
	assert.NotEqual(t, head1, head2)
	head2.Next.Next.Next = head2
	assert.Equal(t, head1, head2)
}
