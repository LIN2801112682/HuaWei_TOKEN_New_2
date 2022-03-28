package dictionary

type TrieTree struct {
	qmin int
	qmax int
	root *TrieTreeNode
}

func (t *TrieTree) Qmin() int {
	return t.qmin
}

func (t *TrieTree) SetQmin(qmin int) {
	t.qmin = qmin
}

func (t *TrieTree) Qmax() int {
	return t.qmax
}

func (t *TrieTree) SetQmax(qmax int) {
	t.qmax = qmax
}

func (t *TrieTree) Root() *TrieTreeNode {
	return t.root
}

func (t *TrieTree) SetRoot(root *TrieTreeNode) {
	t.root = root
}

//初始化trieTree
func NewTrieTree(qmin int, qmax int) *TrieTree {
	return &TrieTree{
		qmin: qmin,
		qmax: qmax,
		root: NewTrieTreeNode(""),
	}
}

//将gram插入trieTree上
//trieTree:待插入的树
//substring:待插入数组字符串
func (tree *TrieTree) InsertIntoTrieTree(substring *[]string) {
	//初始化node、qmin
	node := tree.root
	qmin := tree.qmin
	// 孩子节点在childrenlist中的位置
	var childindex = 0
	for i, str := range *substring {
		childindex = getNode(node.children, (*substring)[i])
		if childindex == -1 {
			// childrenlist里没有该节点
			currentnode := NewTrieTreeNode(str)
			node.children = append(node.children, currentnode)
			node = currentnode
		} else {
			//childrenlist里有该节点
			//childrenindex为该节点在数组中的位置
			node = node.children[childindex]
			node.frequency++
		}
		if i >= qmin-1 {
			node.isleaf = true
		}
	}
}

//剪枝
//trieTree:待修剪的树
//T:阈值
func (tree *TrieTree) PruneTree(T int) {
	tree.root.PruneNode(T)
}

func (tree *TrieTree) PrintTree() {
	tree.root.PrintTreeNode(0)
}

//更新root节点的频率
func (tree *TrieTree) UpdateRootFrequency() {
	for _, child := range tree.root.children {
		tree.root.frequency += child.frequency
	}
	tree.root.frequency--
}
