package go_dic

import (
	"fmt"
	"strings"
	"time"
)

func MatchSearch(searchStr string, root *trieTreeNode, indexRoot *indexTreeNode, qmin int, qmax int) []int {
	start2 := time.Now()
	var vgMap map[int][]string
	vgMap = make(map[int][]string)
	VGCons(root, qmin, qmax, searchStr, vgMap)
	//fmt.Println(vgMap)
	var resArr []int
	seaPositionDis := 0
	var inverPositionDis []int
	fields := strings.Fields(searchStr)
	for i := 0; i < len(fields); i++ { // 0 1 3   len(searchStr)
		tokenArr := vgMap[i]
		if tokenArr != nil {
			seaPositionDis = i - seaPositionDis
			invertIndex = nil
			searchIndexTree(tokenArr, indexRoot, 0)
			if invertIndex == nil {
				return nil
			}
			if i == 0 {
				for j := 0; j < len(invertIndex); j++ {
					sid := invertIndex[j].sid
					inverPositionDis = append(inverPositionDis, invertIndex[j].position)
					resArr = append(resArr, sid)
				}
			} else {
				//var lenRes = len(resArr)
				for j := 0; j < len(resArr); j++ { //遍历之前合并好的resArr
					sidResArr := resArr[j]
					var k int
					for k = 0; k < len(invertIndex); k++ {
						sid := invertIndex[k].sid
						if sidResArr == sid {
							inverPositionDis[j] = invertIndex[k].position - inverPositionDis[j]
							if inverPositionDis[j] == seaPositionDis {
								break
							}
						}
					}
					if k == len(invertIndex) { //新的倒排表id不在之前合并好的结果集resArr 把此id从resArr删除
						resArr = append(resArr[:j], resArr[j+1:]...)
						inverPositionDis = append(inverPositionDis[:j], inverPositionDis[j+1:]...)
						j--
					}
				}
			}
		}
	}
	elapsed2 := time.Since(start2)
	fmt.Println("精确查询花费时间（ms）：", elapsed2)
	return resArr
}

var invertIndex []inverted_index

func searchIndexTree(tokenArr []string, indexRoot *indexTreeNode, i int) {
	if indexRoot == nil {
		return
	}
	for j := 0; j < len(indexRoot.children); j++ {
		if i == len(tokenArr)-1 && tokenArr[i] == indexRoot.children[j].data {
			for k := 0; k < len(indexRoot.children[j].invertedIndexList); k++ {
				invertIndex = append(invertIndex, *indexRoot.children[j].invertedIndexList[k])
			}
		} else if tokenArr[i] == indexRoot.children[j].data {
			searchIndexTree(tokenArr, indexRoot.children[j], i+1)
		}
	}
}
