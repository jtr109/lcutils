package listnode_test

import (
	"testing"

	"github.com/jtr109/lcutils/listnode"
)

func TestTypeAlias(t *testing.T) {
	type ListNode = listnode.ListNode
	_ = &ListNode{
		Next: &ListNode{},
	}
}
