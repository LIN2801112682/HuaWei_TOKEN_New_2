package build_VToken_index

type IndexTree struct {
	qmin int
	qmax int
	Cout int
	Root *IndexTreeNode
}

//初始化trieTree
func NewIndexTree(qmin int, qmax int) *IndexTree {
	return &IndexTree{
		qmin: qmin,
		qmax: qmax,
		Cout: 0,
		Root: NewIndexTreeNode(""),
	}
}

//将gram插入trieTree上
//trieTree:待插入的树
//token:待插入数组字符串
//sid:字符串所属sid
//position:字符串在sid中的位置
func InsertIntoIndexTree(tree *IndexTree, token *[]string, sid SeriesId, position int) {
	//初始化node、qmin
	node := tree.Root
	//qmin := tree.qmin
	// 孩子节点在childrenlist中的位置
	var childindex = 0
	for i, str := range *token {
		childindex = getIndexNode(node.Children, (*token)[i])
		if childindex == -1 {
			// childrenlist里没有该节点
			currentnode := NewIndexTreeNode(str)
			IndexNodeArrayInsertStrategy(&node.Children, currentnode)
			node = currentnode
		} else {
			//childrenlist里有该节点
			//childrenindex为该节点在数组中的位置
			node = node.Children[childindex]
			node.Frequency++
		}
		//从root的孩子节点开始判断，少一层故大于等于 qmin-1 不是qmin
		//if i >= qmin-1 {
		//	node.isleaf = true
		//}
		if i == len(*token)-1 {
			//InsertInvertedIndex(node, sid, position)
			//叶子节点，需要挂倒排链表
			node.isleaf = true
			if _, ok := node.InvertedIndex[sid]; !ok { //key中没有sid 创建sid对应的倒排
				InsertSidAndPosArrToInvertedIndexMap(node, sid, position)
			} else { //寻找相同sid下增加posArray即可
				InsertPosArrToInvertedIndexMap(node, sid, position)
			}
		}
	}
}

func InsertPosArrToInvertedIndexMap(node *IndexTreeNode, sid SeriesId, position int) {
	//倒排列表数组中找到sid的invertedIndex，把position加入到invertedIndex中的posArray中去
	node.InvertedIndex[sid] = append(node.InvertedIndex[sid], position)
}

//插入倒排
func InsertSidAndPosArrToInvertedIndexMap(node *IndexTreeNode, sid SeriesId, position int) {
	posArray := []int{}
	posArray = append(posArray, position)
	node.InvertedIndex[sid] = posArray
}

func PrintIndexTree(tree *IndexTree) {
	PrintIndexTreeNode(tree.Root, 0)
}

//更新root节点的频率
func UpdateIndexRootFrequency(tree *IndexTree) {
	for _, child := range tree.Root.Children {
		tree.Root.Frequency += child.Frequency
	}
	tree.Root.Frequency--
}
