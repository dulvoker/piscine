package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return root
	}
	if root.Data < elem {
		if root.Right == nil {
			return nil
		}
		return BTreeSearchItem(root.Right, elem)
	}
	if root.Left == nil {
		return nil
	}
	return BTreeSearchItem(root.Left, elem)
}
