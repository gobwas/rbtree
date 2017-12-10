package rbtree

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Min() *Node {
	if n != nil && n.Left != nil {
		return n.Left.Min()
	}
	return n
}

func (n *Node) Max() *Node {
	if n != nil && n.Right != nil {
		return n.Right.Min()
	}
	return n
}

func (n *Node) Find(key int) *Node {
	if n == nil {
		return nil
	}
	if key < n.Key {
		return n.Left.Find(key)
	}
	if key > n.Key {
		return n.Right.Find(key)
	}
	return n
}

func (n *Node) Insert(key int) (root *Node, inserted bool) {
	if n == nil {
		return &Node{Key: key}, true
	}
	if key < n.Key {
		n.Left, inserted = n.Left.Insert(key)
		return n, inserted
	}
	if key > n.Key {
		n.Right, inserted = n.Right.Insert(key)
		return n, inserted
	}
	// Duplicates are not allowed in binary trees.
	return n, false
}

func (n *Node) Delete(key int) (root *Node, deleted bool) {
	if n == nil {
		return nil, false
	}
	if key < n.Key {
		n.Left, deleted = n.Left.Delete(key)
		return n, deleted
	}
	if key > n.Key {
		n.Right, deleted = n.Right.Delete(key)
		return n, deleted
	}
	return n.Destroy(), true
}

func (n *Node) Destroy() (root *Node) {
	switch {
	case n.Left != nil && n.Right != nil:
		min := n.Right.Min()
		n.Key = min.Key

		var deleted bool
		n.Right, deleted = n.Right.Delete(min.Key)
		if !deleted {
			panic("inconsistent tree state")
		}

		return n

	case n.Left != nil:
		return n.Left
	case n.Right != nil:
		return n.Right
	default:
		return nil
	}
}

func (n *Node) RotateLeft() (root *Node) {
	var (
		pivot     = n.Right
		pivotLeft = n.Right.Left
	)

	n.Right = pivotLeft
	pivot.Left = n

	return pivot
}

func (n *Node) RotateRight() (root *Node) {
	var (
		pivot      = n.Left
		pivotRight = n.Left.Right
	)

	n.Left = pivotRight
	pivot.Right = n

	return pivot
}

func (n *Node) InOrder(it func(int)) {
	if n == nil {
		return
	}
	n.Left.InOrder(it)
	it(n.Key)
	n.Right.InOrder(it)
}

func (n *Node) PreOrder(it func(int)) {
	if n == nil {
		return
	}
	it(n.Key)
	n.Left.PreOrder(it)
	n.Right.PreOrder(it)
}

func (n *Node) PostOrder(it func(int)) {
	if n == nil {
		return
	}
	n.Left.PostOrder(it)
	n.Right.PostOrder(it)
	it(n.Key)
}
