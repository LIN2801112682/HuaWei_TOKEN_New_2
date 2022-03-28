package main

import (
	"dictionary"
	"fmt"
	"index07"
	_ "indexMaintain"
	_ "matchQuery1"
	"matchQuery2"
	"runtime"
)

func traceMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func main() {
	fmt.Println("字典树D：===============================================================")
	fmt.Println("字典树D内存占用大小：")
	//traceMemStats()
	fmt.Println()
	root := dictionary.GenerateDictionaryTree("src/resources/50000Dic.txt", 1, 2, 100) //
	fmt.Println()
	//traceMemStats()
	fmt.Println()

	fmt.Println("索引项集：===============================================================")
	fmt.Println()
	fmt.Println("索引项集内存占用大小：")
	traceMemStats()
	fmt.Println()
	_, indexTreeNode := index07.GenerateIndexTree("src/resources/500000Index.txt", 1, 2, root) //
	fmt.Println()
	traceMemStats()
	fmt.Println()

	/*indexTreeNode.FixInvertedIndexSize()
	sort.SliceStable(index07.Res, func(i, j int) bool {
		if index07.Res[i] < index07.Res[j]  {
			return true
		}
		return false
	})
	fmt.Println(index07.Res)
	fmt.Println(len(index07.Res))
	sum := 0
	for _,val := range index07.Res{
		sum += val
	}
	fmt.Println(index07.Res[0])
	fmt.Println(index07.Res[len(index07.Res)-1])
	fmt.Println(index07.Res[len(index07.Res)/2])
	fmt.Println(sum/len(index07.Res))*/

	/*indexTreeNode.SearchGramsFromIndexTree()
	fmt.Println(len(index07.Grams))
	var numsOfgrams2_12 [4]int
	for _,val := range index07.Grams{
		numsOfgrams2_12[len(val)]++
	}
	fmt.Println(numsOfgrams2_12)*/

	/*fmt.Println("新增索引后的索引项集：===============================================================")
	fmt.Println()
	fmt.Println("索引项集内存占用大小：")
	//traceMemStats()
	fmt.Println()
	indexMaintain.AddIndex("src/resources/add2000.txt", 1, 3, root, indexTree)
	fmt.Println()
	//traceMemStats()
	fmt.Println()*/

	//resInt := matchQuery2.MatchSearch("get english", root, indexTreeNode, 1, 2)
	var searchQuery = [9]string{"get", "get english", "get english images", "get images", "get english images team_hm_header_shad.gif http 1.0", "get images s102325 gif http 1.0", "get english history history_of images cup", "images space.gif", "get http 1.0"}
	for i := 0; i < 9; i++ {
		resInt := matchQuery2.MatchSearch(searchQuery[i], root, indexTreeNode, 1, 2) //get english venues
		//fmt.Println(resInt)
		fmt.Println(len(resInt))
		fmt.Println("==================================================")
	}
	//fmt.Println(len(resInt))
}
