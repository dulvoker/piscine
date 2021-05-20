package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root.Data == elem {
		return root
	}
	if root == nil {
		return root.Parent
	}
	if root.Data < elem {
		return BTreeSearchItem(root.Right, elem)
	}
	return BTreeSearchItem(root.Left, elem)
}
