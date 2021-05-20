package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root.Data == elem {
		return root
	}
	if root == nil {
		return root.Parent
	}
	if root.Data < elem {
		root.Right = BTreeSearchItem(root.Right, elem)
	}
	root.Left = BTreeSearchItem(root.Left, elem)
	return root
}
