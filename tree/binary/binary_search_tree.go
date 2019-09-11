package binary

// SearchTree 是二叉搜索树
type SearchTree struct {
	*TreeNode
}

// TreeNode 是一棵树的节点
type TreeNode struct {
	key    int
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
}

// Root 返回二叉搜索树的根
func (t *SearchTree) Root() *TreeNode {
	return t.TreeNode
}

// GetKey 返回节点的key
func (t *TreeNode) GetKey() int {
	return t.key
}

// InorderWalk 使用终须遍历二叉树
func (t *TreeNode) InorderWalk() (nums []int) {
	if t != nil {
		nums = append(nums, t.left.InorderWalk()...)
		nums = append(nums, t.key)
		nums = append(nums, t.right.InorderWalk()...)
	}
	return
}

// RecursiveInsert 新建一个节点值为key的节点， 并插入合适的位置， 保持二叉搜索树的特性不变
func (t *TreeNode) RecursiveInsert(key int) *TreeNode {
	if t == nil {
		return newNode(key)
	}
	if key > t.key {
		t.right = t.right.RecursiveInsert(key)
		t.right.parent = t.right
	} else if key < t.key {
		t.left = t.left.RecursiveInsert(key)
		t.left.parent = t.left
	}
	return t
}

// transplant 移植v到u所在的位置
func transplant(u, v *TreeNode) {
	if u.parent != nil {
		if u.parent.left == u {
			u.parent.left = v
		} else if u.parent.right == u {
			u.parent.right = v
		}
	}
	if v != nil {
		v.parent = u.parent
	}
}

// Delete 在树中查找节点值为key的节点, 删除该节点
func (t *TreeNode) Delete(key int) {
	if node := t.Find(key); node != nil {
		if node.left == nil {
			transplant(node, node.right)
		} else if node.right == nil {
			transplant(node, node.left)
		} else {
			suc := node.Successor()
			node.key = suc.key
			transplant(suc, nil)
		}
	}
}

// IterativeInsert 使用循环的方式插入值为key的节点
func (t *TreeNode) IterativeInsert(key int) {
	node := t
	var targetNode *TreeNode

	for node != nil {
		targetNode = node
		if node.key < key {
			node = node.right
		} else if node.key > key {
			node = node.left
		}
	}
	n := newNode(key)
	if targetNode.key < key {
		targetNode.right = n
	}
	if targetNode.key > key {
		targetNode.left = n
	}
	n.parent = targetNode
}

// Successor 返回按照大小排序当前节点后面的一个节点
func (t *TreeNode) Successor() *TreeNode {
	if t.right != nil {
		return t.right.IterativeMinimum()
	}
	node := t
	for node.parent.right == node {
		node = node.parent
	}
	return node.parent
}

// Predecessor 返回按照大小排序的当前节点的前面的一个节点
func (t *TreeNode) Predecessor() *TreeNode {
	if t.left != nil {
		return t.left.IterativeMaximum()
	}
	node := t
	for node.parent.left == node {
		node = node.parent
	}
	return node.parent
}

// Find 搜寻二叉树，找到值和key相等的节点，返回该节点的指针
// 如果没有找到， 那么返回nil
func (t *TreeNode) Find(key int) *TreeNode {
	node := t

	for node != nil && node.key != key {
		if node.key < key {
			node = node.right
		} else {
			node = node.left
		}
	}
	return node
}

// IterativeMinimum 使用循环返回最小值
func (t *TreeNode) IterativeMinimum() *TreeNode {
	node := t
	for node.left != nil {
		node = node.left
	}
	return node
}

// IterativeMaximum 使用循环找到最大值
func (t *TreeNode) IterativeMaximum() *TreeNode {
	node := t
	for node.right != nil {
		node = node.right
	}
	return node
}

// RecursiveMinimum 使用递归找到最小值
func (t *TreeNode) RecursiveMinimum() *TreeNode {
	if t.left != nil {
		return t.left.RecursiveMinimum()
	}
	return t
}

// RecursiveMaximum 使用递归找到最大值
func (t *TreeNode) RecursiveMaximum() *TreeNode {
	if t.right != nil {
		return t.right.RecursiveMaximum()
	}
	return t
}

func newNode(key int) *TreeNode {
	return &TreeNode{
		key:    key,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

// NewTree 从数字数组简历一个搜索二叉树， 返回该二叉树的根
func NewTree(nums []int) (t *SearchTree) {
	if len(nums) <= 0 {
		return nil
	}
	t = &SearchTree{
		TreeNode: newNode(nums[0]),
	}
	for _, n := range nums[1:] {
		t.IterativeInsert(n)
	}
	return
}
