package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root != nil {
		return (BTreeLevelCount(root.Left) + BTreeLevelCount(root.Right) + 1)
	} else {
		return 0
	}
}
