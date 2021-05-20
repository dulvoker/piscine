package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root.Data == elem {
		return root
	}
	if root.Data > elem {
		BTreeSearchItem(root.Left, elem)
	} else {
		BTreeSearchItem(root.Right, elem)
	}
	return root
}
