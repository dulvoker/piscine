package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return root.Parent
	}
	return BTreeMax(root.Right)
}
