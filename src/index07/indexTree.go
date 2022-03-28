package index07

type IndexTree struct {
	qmin int
	qmax int
	cout int
	root *IndexTreeNode
}

func (i *IndexTree) Qmin() int {
	return i.qmin
}

func (i *IndexTree) SetQmin(qmin int) {
	i.qmin = qmin
}

func (i *IndexTree) Qmax() int {
	return i.qmax
}

func (i *IndexTree) SetQmax(qmax int) {
	i.qmax = qmax
}

func (i *IndexTree) Cout() int {
	return i.cout
}

func (i *IndexTree) SetCout(cout int) {
	i.cout = cout
}

func (i *IndexTree) Root() *IndexTreeNode {
	return i.root
}

func (i *IndexTree) SetRoot(root *IndexTreeNode) {
	i.root = root
}

//初始化trieTree
func NewIndexTree(qmin int, qmax int) *IndexTree {
	return &IndexTree{
		qmin: qmin,
		qmax: qmax,
		cout: 0,
		root: NewIndexTreeNode(""),
	}
}

//将gram插入trieTree上
//trieTree:待插入的树
//token:待插入数组字符串
//sid:字符串所属sid
//position:字符串在sid中的位置
func (tree *IndexTree) InsertIntoIndexTree(token *[]string, sid SeriesId, position int) {
	//初始化node、qmin
	node := tree.root
	//qmin := tree.qmin
	// 孩子节点在childrenlist中的位置
	var childindex = 0
	for i, str := range *token {
		childindex = GetIndexNode(node.children, (*token)[i])
		if childindex == -1 {
			// childrenlist里没有该节点
			currentnode := NewIndexTreeNode(str)
			node.children = append(node.children, currentnode)
			node = currentnode
		} else {
			//childrenlist里有该节点
			//childrenindex为该节点在数组中的位置
			node = node.children[childindex]
			node.frequency++
		}
		//从root的孩子节点开始判断，少一层故大于等于 qmin-1 不是qmin
		//if i >= qmin-1 {
		//	node.isleaf = true
		//}
		if i == len(*token)-1 {
			//InsertInvertedIndex(node, sid, position)
			//叶子节点，需要挂倒排链表
			node.isleaf = true
			if _, ok := node.invertedIndex[sid]; !ok { //key中没有sid 创建sid对应的倒排
				node.InsertSidAndPosArrToInvertedIndexMap(sid, position)
			} else { //寻找相同sid下增加posArray即可
				node.InsertPosArrToInvertedIndexMap(sid, position)
			}
		}
	}
}

func (tree *IndexTree) PrintIndexTree() {
	tree.root.PrintIndexTreeNode(0)
}

//更新root节点的频率
func (tree *IndexTree) UpdateIndexRootFrequency() {
	for _, child := range tree.root.children {
		tree.root.frequency += child.frequency
	}
	tree.root.frequency--
}

var Res []int

func (root *IndexTreeNode) FixInvertedIndexSize() {
	for i := 0; i < len(root.children); i++ {
		if root.children[i].isleaf == true {
			Res = append(Res, len(root.children[i].invertedIndex))
		}
		root.children[i].FixInvertedIndexSize()
	}
}

var Grams [][]string
var temp []string

func (root *IndexTreeNode) SearchGramsFromIndexTree() {
	if root == nil {
		return
	}
	for i := 0; i < len(root.children); i++ {
		temp = append(temp, root.children[i].data)
		if root.children[i].isleaf == true {
			Grams = append(Grams, temp)
		}
		root.children[i].SearchGramsFromIndexTree()
		if len(temp) > 0 {
			temp = temp[:len(temp)-1]
		}
	}
}
