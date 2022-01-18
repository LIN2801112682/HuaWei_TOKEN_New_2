package index_maintenance

import (
	"bufio"
	"build_VToken_index"
	"build_dictionary"
	"fmt"
	"io"
	"os"
	"time"
)

//根据一批日志数据通过字典树划分VG，增加到索引项集中
func AddIndex(filename string, qmin int, qmax int, root *build_dictionary.TrieTreeNode, indexTree *build_VToken_index.IndexTree) *build_VToken_index.IndexTree {
	data, err := os.Open(filename)
	defer data.Close()
	if err != nil {
		fmt.Print(err)
	}
	buff := bufio.NewReader(data)
	sid := indexTree.Cout
	var sum = 0
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		var vgMap map[int][]string
		vgMap = make(map[int][]string)
		sid++
		str := string(data)
		start2 := time.Now()
		build_VToken_index.VGCons(root, qmin, qmax, str, vgMap)
		for vgKey := range vgMap {
			tokenArr := vgMap[vgKey]
			build_VToken_index.InsertIntoIndexTree(indexTree, &tokenArr, sid, vgKey)
		}
		end2 := time.Since(start2).Microseconds()
		sum = int(end2) + sum
	}
	indexTree.Cout = sid
	indexTree.Root.Frequency = 1
	build_VToken_index.UpdateIndexRootFrequency(indexTree)
	fmt.Println("新增索引项集花费时间(us)：", sum)
	//PrintIndexTree(indexTree)
	return indexTree
}
