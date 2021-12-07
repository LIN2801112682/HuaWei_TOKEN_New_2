package go_dic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func GererateTree(filename string, qmin int, qmax int, T int) *trieTreeNode {
	start1 := time.Now()
	tree := NewTrieTree(qmin, qmax)
	data, err := os.Open(filename)
	defer data.Close()
	if err != nil {
		fmt.Println(err)
	}
	buff := bufio.NewReader(data)
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		str := (string)(data)
		for i := 0; i < len(str)-qmax; i++ {
			substring := str[i : i+qmax]
			//字符串变字符串数组
			gram := make([]string, qmax)
			for j := 0; j < qmax; j++ {
				gram[j] = substring[j : j+1]
			}
			InsertIntoTrieTree(tree, &gram)
		}
		for i := len(str) - qmax; i < len(str)-qmin+1; i++ {
			substring := str[i:len(str)]
			gram := make([]string, len(str)-i)
			for j := 0; j < len(str)-i; j++ {
				gram[j] = substring[j : j+1]
			}
			InsertIntoTrieTree(tree, &gram)
		}
	}
	PruneTree(tree, T)
	UpdateRootFrequency(tree)
	elapsed1 := time.Since(start1)
	fmt.Println("构建字典树花费时间（ms）：", elapsed1)
	PrintTree(tree)
	return tree.root
}
