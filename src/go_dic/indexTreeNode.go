package go_dic
import (
	"fmt"
)

type indexTreeNode struct {
	data string
	frequency int
	children []*indexTreeNode
	isleaf bool
	invertedIndexList []*inverted_index
}

func NewIndexTreeNode (data string) *indexTreeNode {
	return &indexTreeNode{
		data: data,
		frequency: 1,
		isleaf: false,
		children: make([]*indexTreeNode, 0),
		invertedIndexList: make([]*inverted_index, 0),
	}
}

//插入倒排
func InsertInvertedIndex(node *indexTreeNode, sid int , position int){
	// 倒排列表数组中创建新inverted_index，并加入到invertedIndexList中
	newInverted := NewInverted_index(sid , position)
	invertedIndexArrayInsertStrategy(&node.invertedIndexList , newInverted)
}


//插入数组策略
func IndexNodeArrayInsertStrategy(array *[]*indexTreeNode, node *indexTreeNode){
	*array =  append(*array , node)
}

//插入倒排链表策略
func invertedIndexArrayInsertStrategy(array *[]*inverted_index, invertedindex *inverted_index){
	*array =  append( *array , invertedindex)
}

//判断children有无此节点
func getIndexNode(children []*indexTreeNode, char string) int{
	for i , child := range children{
		if child.data == char{
			return i
		}
	}
	return -1
}

//输出以node为根的子树
func PrintIndexTreeNode(node *indexTreeNode, level int) {
	fmt.Println()
	for i := 0 ; i < level ; i++{
		fmt.Print("      ")
	}
	fmt.Print(node.data , " - " , node.frequency , " - " , node.isleaf)
	for _ , invertedIndex := range node.invertedIndexList{
		fmt.Print("  /  sid : " , invertedIndex.sid , " position : " , invertedIndex.position)
	}
	for _ , child := range node.children{
		PrintIndexTreeNode(child , level + 1)
	}
}