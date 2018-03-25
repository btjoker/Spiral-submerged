package btree

import "fmt"

// TreeNode 二叉树
type TreeNode struct {
	leftNode  *TreeNode
	rightNode *TreeNode
	value     int
}

// TreeRoot 二叉树根节点
type TreeRoot struct {
	treeRoot *TreeNode
}

// GetLeftTreeNode 获取左节点
func (t *TreeNode) GetLeftTreeNode() *TreeNode {
	return t.leftNode
}

// GetRightTreeNode 获取右节点
func (t *TreeNode) GetRightTreeNode() *TreeNode {
	return t.rightNode
}

// SetLeftTreeNode 设置左分叉
func (t *TreeNode) SetLeftTreeNode(node *TreeNode) {
	t.leftNode = node
}

// SetRightTreeNode 设置右分叉
func (t *TreeNode) SetRightTreeNode(node *TreeNode) {
	t.rightNode = node
}

// GetValue 获取节点保存的值
func (t *TreeNode) GetValue() int {
	return t.value
}

// GetTreeRoot 获取根节点
func (r *TreeRoot) GetTreeRoot() *TreeNode {
	return r.treeRoot
}

// SetTreeRoot 设置根节点
func (r *TreeRoot) SetTreeRoot(treeRoot *TreeNode) {
	r.treeRoot = treeRoot
}

// NewTreeNode 创建一个新的节点
func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		value: value,
	}
}

// InTraverseBTree 中序遍历
func InTraverseBTree(root *TreeNode) {
	if root != nil {
		fmt.Println(root.GetValue())
		InTraverseBTree(root.GetLeftTreeNode())
		InTraverseBTree(root.GetRightTreeNode())
	}
}

// PreTraverseBTree 先序遍历
func PreTraverseBTree(root *TreeNode) {
	if root != nil {
		PreTraverseBTree(root.GetLeftTreeNode())
		fmt.Println(root.GetValue())
		PreTraverseBTree(root.GetRightTreeNode())
	}
}

// PostTraverseBTree 后序遍历
func PostTraverseBTree(root *TreeNode) {
	if root != nil {
		PostTraverseBTree(root.GetLeftTreeNode())
		PostTraverseBTree(root.GetRightTreeNode())
		fmt.Println(root.GetValue())
	}
}

// CreateTree 按照数组创建二叉树
// 小的值在根左边, 大的值在根右边
func CreateTree(root *TreeRoot, value int) {
	if root.GetTreeRoot() == nil {
		treeNode := NewTreeNode(value)
		root.SetTreeRoot(treeNode)
	} else {
		tempRoot := root.GetTreeRoot()

		for tempRoot != nil {
			if value > tempRoot.GetValue() {
				if tempRoot.GetRightTreeNode() == nil {
					tempRoot.SetRightTreeNode(NewTreeNode(value))
					break
				}
				tempRoot = tempRoot.GetRightTreeNode()
			} else {
				if tempRoot.GetLeftTreeNode() == nil {
					tempRoot.SetLeftTreeNode(NewTreeNode(value))
					break
				}
				tempRoot = tempRoot.GetLeftTreeNode()
			}
		}
	}
}

// GetHeight 获取二叉树最大深度
func GetHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := GetHeight(root.GetLeftTreeNode())
	right := GetHeight(root.GetRightTreeNode())

	max := left
	if right > max {
		max = right
	}
	return max + 1
}

// GetMax 获取树节点中最大值
func GetMax(root *TreeNode) int {
	if root == nil {
		return -1
	}

	left := GetMax(root.GetLeftTreeNode())
	right := GetMax(root.GetRightTreeNode())

	currentRootValue := root.GetValue()

	max := left

	if right > max {
		max = right
	}
	if currentRootValue > max {
		max = currentRootValue
	}
	return max
}
