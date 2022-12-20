package ds

type node struct {
	data  int
	left  *node
	right *node
}

func (n *node) addLeftNode(x int) {
	temp := new(node)
	temp.data = x
	n.left = temp
}
func (n *node) addRightNode(x int) {
	temp := new(node)
	temp.data = x
	n.right = temp
}

// here we need to use one of the traversal algorithms
// DFS is preferred for trees
func (n *node) printTree() {

}

// func main() {
// 	var head *node

// }
