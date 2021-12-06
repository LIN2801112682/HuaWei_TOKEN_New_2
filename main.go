package main

import (
	"fmt"
	"go_dic"
)

func main(){
	fmt.Println("字典树D：===============================================================")
	root := go_dic.GererateTree("src/resources/4Dic.txt", 2, 3, 6)
	fmt.Println()
	fmt.Println("索引项集：===============================================================")
	go_dic.GererateIndex("src/resources/1Index.txt",2,3,root)
}