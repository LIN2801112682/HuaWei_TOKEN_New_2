package go_dic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
		tokenArray := strings.Fields(str)
		for i := 0; i < len(tokenArray)-qmax; i++ {
			var substring = tokenArray[i : i+qmax]
			InsertIntoTrieTree(tree, &substring)
		}
		for i := len(tokenArray) - qmax; i < len(tokenArray)-qmin+1; i++ {
			var substring = tokenArray[i:len(tokenArray)]
			InsertIntoTrieTree(tree, &substring)
		}
	}
	PruneTree(tree, T)
	UpdateRootFrequency(tree)
	elapsed1 := time.Since(start1)
	fmt.Println("构建字典树花费时间（ms）：", elapsed1)
	PrintTree(tree)
	return tree.root
}
