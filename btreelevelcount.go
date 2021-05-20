package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := BTreeLevelCount(root.Left) + 1
	sum = sum + BTreeLevelCount(root.Right) + 1
	return sum
}
