package new_precise_query

import (
	"build_VToken_index"
	"build_dictionary"
	"fmt"
	"github.com/imdario/mergo"
	"sort"
	_ "strings"
	"time"
)

func MatchSearch(searchStr string, root *build_dictionary.TrieTreeNode, indexRoot *build_VToken_index.IndexTreeNode, qmin int, qmax int) []build_VToken_index.SeriesId {
	var vgMap map[int][]string
	vgMap = make(map[int][]string)
	build_VToken_index.VGCons(root, qmin, qmax, searchStr, vgMap)
	fmt.Println(vgMap)
	var keys = []int{}
	for key := range vgMap {
		keys = append(keys, key)
	}
	sort.Sort(sort.IntSlice(keys)) //对map中的key进行排序（map遍历是无序的）
	var resArr []build_VToken_index.SeriesId
	preSeaPosition := 0
	var preInverPositionDis []PosList
	var nowInverPositionDis []PosList
	//fields := strings.Fields(searchStr)
	start2 := time.Now()
	for m := 0; m < len(keys); m++ {
		i := keys[m]
		tokenArr := vgMap[i]
		if tokenArr != nil {
			nowSeaPosition := i
			invertIndex = nil
			invertIndex2 = nil
			searchIndexTreeFromLeaves(tokenArr, indexRoot, 0)
			searchListsTreeFromLeaves(indexNode)
			mergo.Merge(&invertIndex, invertIndex2)
			fmt.Println(len(invertIndex))
			if invertIndex == nil {
				return nil
			}
			if i == 0 {
				for sid := range invertIndex {
					preInverPositionDis = append(preInverPositionDis, NewPosList(sid, make([]int, len(invertIndex[sid]), len(invertIndex[sid]))))
					nowInverPositionDis = append(nowInverPositionDis, NewPosList(sid, invertIndex[sid]))
					resArr = append(resArr, sid)
				}
			} else {
				for j := 0; j < len(resArr); j++ { //遍历之前合并好的resArr
					findFlag := false
					sid := resArr[j]
					if _, ok := invertIndex[sid]; ok {
						nowInverPositionDis[j] = NewPosList(sid, invertIndex[sid])
						for z1 := 0; z1 < len(preInverPositionDis[j].PosArray); z1++ {
							z1Pos := preInverPositionDis[j].PosArray[z1]
							for z2 := 0; z2 < len(nowInverPositionDis[j].PosArray); z2++ {
								z2Pos := nowInverPositionDis[j].PosArray[z2]
								if nowSeaPosition-preSeaPosition == z2Pos-z1Pos {
									findFlag = true
									break
								}
							}
							if findFlag == true {
								break
							}
						}
					}
					if findFlag == false { //没找到并且候选集的sid比resArr大，删除resArr[j]
						resArr = append(resArr[:j], resArr[j+1:]...)
						preInverPositionDis = append(preInverPositionDis[:j], preInverPositionDis[j+1:]...)
						nowInverPositionDis = append(nowInverPositionDis[:j], nowInverPositionDis[j+1:]...)
						j-- //删除后重新指向，防止丢失元素判断
					}
				}
			}
			preSeaPosition = nowSeaPosition
			copy(preInverPositionDis, nowInverPositionDis)
		}
	}
	elapsed2 := time.Since(start2).Microseconds()
	fmt.Println("精确查询花费时间（us）：", elapsed2)
	sort.SliceStable(resArr, func(i, j int) bool {
		if resArr[i].Id < resArr[j].Id && resArr[i].Time < resArr[j].Time {
			return true
		}
		return false
	})
	return resArr
}

var invertIndex build_VToken_index.Inverted_index
var indexNode *build_VToken_index.IndexTreeNode

//查询当前串对应的倒排表（叶子节点）
func searchIndexTreeFromLeaves(tokenArr []string, indexRoot *build_VToken_index.IndexTreeNode, i int) {
	if indexRoot == nil {
		return
	}
	for j := 0; j < len(indexRoot.Children); j++ {
		if i < len(tokenArr)-1 && tokenArr[i] == indexRoot.Children[j].Data {
			searchIndexTreeFromLeaves(tokenArr, indexRoot.Children[j], i+1)
		}
		if i == len(tokenArr)-1 && tokenArr[i] == indexRoot.Children[j].Data { //找到那一层的倒排表
			invertIndex = indexRoot.Children[j].InvertedIndex
			indexNode = indexRoot.Children[j]
		}
	}
}

var invertIndex2 build_VToken_index.Inverted_index

func searchListsTreeFromLeaves(indexNode *build_VToken_index.IndexTreeNode) {
	if indexNode != nil {
		for l := 0; l < len(indexNode.Children); l++ {
			if len(indexNode.Children[l].InvertedIndex) > 0 {
				mergo.Merge(&invertIndex2, indexNode.Children[l].InvertedIndex)
			}
			searchListsTreeFromLeaves(indexNode.Children[l])
		}
	}
}
