package tree

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) AddLeftNode(x int) {
	temp := new(TreeNode)
	temp.Data = x
	n.Left = temp
}
func (n *TreeNode) AddRightNode(x int) {
	temp := new(TreeNode)
	temp.Data = x
	n.Right = temp
}

// here we need to use one of the traversal algorithms
// DFS is preferred for trees
func (n *TreeNode) PrintTree() {

}
