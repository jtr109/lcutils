package treenode

type Operator struct {
	root *TreeNode
}

func NewOperator() *Operator {
	return &Operator{}
}

func (op *Operator) Root() *TreeNode {
	return op.root
}

func (op *Operator) FromSlice(s []NilInt) {
	op.root = nil
	if len(s) == 0 {
		return
	}
	op.root = NewTreeNode(&s[0])
	nodeQueue := []*TreeNode{op.root}
	i := 0
	for len(nodeQueue) > 0 && i < len(s) {
		current := nodeQueue[0]
		var leftChild *TreeNode
		var rightChild *TreeNode
		i++
		if i < len(s) && !s[i].IsNil() {
			leftChild = NewTreeNode(&s[i])
			nodeQueue = append(nodeQueue, leftChild)
		}
		i++
		if i < len(s) && !s[i].IsNil() {
			rightChild = NewTreeNode(&s[i])
			nodeQueue = append(nodeQueue, rightChild)
		}
		current.Left = leftChild
		current.Right = rightChild
		nodeQueue = nodeQueue[1:] // left pop
	}
}
