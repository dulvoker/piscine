package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	node := &TreeNode{Data: data}
	if root == nil {
		root = node
	} else if root.Data > node.Data {
		root.Left.Parent = root
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right.Parent = root
		root.Right = BTreeInsertData(root.Right, data)
	}
	return root
}
