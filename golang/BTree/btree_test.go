package btree

import (
	"testing"
)

func NewTestTreeNode() *TreeNode {
	list := []int{2, 3, 1, 4, 5}
	root := new(TreeRoot)
	for _, v := range list {
		CreateTree(root, v)
	}
	return root.GetTreeRoot()
}

func TestGetMax(t *testing.T) {
	root := NewTestTreeNode()

	t.Run("2, 3, 1, 4, 5", func(t *testing.T) {
		if got := GetMax(root); got != 5 {
			t.Errorf("GetMax() = %v, want %v", got, 5)
		}
	})
}

func TestGetHeight(t *testing.T) {
	root := NewTestTreeNode()

	t.Run("2, 3, 1, 4, 5", func(t *testing.T) {
		if got := GetHeight(root); got != 4 {
			t.Errorf("GetHeight() = %v, want %v", got, 4)
		}
	})
}

func TestTreeNode_GetValue(t *testing.T) {
	root := NewTreeNode(110)

	t.Run("110", func(t *testing.T) {
		if got := root.GetValue(); got != 110 {
			t.Errorf("TreeNode.GetValue() = %v, want %v", got, 110)
		}
	})
}
