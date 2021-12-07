package main

import (
	"fmt"
	"go_dic"
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
	traceMemStats()
	fmt.Println()
	root := go_dic.GererateTree("src/resources/50000Dic.txt", 2, 3, 10) //
	fmt.Println()
	traceMemStats()

	fmt.Println()
	fmt.Println("索引项集：===============================================================")
	fmt.Println()
	fmt.Println("索引项集内存占用大小：")
	traceMemStats()
	fmt.Println()
	go_dic.GererateIndex("src/resources/100Dic.txt", 2, 3, root)
	fmt.Println()
	traceMemStats()
	fmt.Println()
}
