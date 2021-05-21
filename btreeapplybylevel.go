package piscine

func fapply(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else if level > 1 {
		fapply(root.Left, level-1, f)
		fapply(root.Right, level-1, f)
	}
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		a := BTreeLevelCount(root)
		for b := 1; b <= a; b++ {
			if root == nil {
				return
			}
			if b == 1 {
				f(root.Data)
			} else if b > 1 {
				fapply(root.Left, b-1, f)
				fapply(root.Right, b-1, f)
			}
		}
	}
}
