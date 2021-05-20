package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return root.Parent
	}
	if root.Data == elem {
		return root
	}
	if root.Data < elem {
		if root.Right == nil {
			return root
		}
		return BTreeSearchItem(root.Right, elem)
	}
	if root.Left == nil {
		return root
	}
	return BTreeSearchItem(root.Left, elem)
}
