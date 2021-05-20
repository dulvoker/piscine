package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root != nil {
		rightside := BTreeLevelCount(root.Right) + 1
		leftside := BTreeLevelCount(root.Left) + 1
		if rightside > leftside {
			return rightside
		}
		return leftside
	} else {
		return 0
	}
}
