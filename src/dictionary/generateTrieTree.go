package dictionary

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func GenerateDictionaryTree(filename string, qmin int, qmax int, T int) *TrieTreeNode {
	tree := NewTrieTree(qmin, qmax)
	data, err := os.Open(filename)
	defer data.Close()
	if err != nil {
		fmt.Println(err)
	}
	buff := bufio.NewReader(data)
	var sum = 0
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		str := (string)(data)
		start2 := time.Now()
		tokenArray := strings.Fields(str)
		for i := 0; i < len(tokenArray)-qmax; i++ {
			var substring = tokenArray[i : i+qmax]
			tree.InsertIntoTrieTree(&substring)
		}
		for i := len(tokenArray) - qmax; i < len(tokenArray)-qmin+1; i++ {
			var substring = tokenArray[i:len(tokenArray)]
			tree.InsertIntoTrieTree(&substring)
		}
		end2 := time.Since(start2).Microseconds()
		sum = int(end2) + sum
	}
	start1 := time.Now()
	tree.PruneTree(T)
	end1 := time.Since(start1).Microseconds()
	sum = int(end1) + sum
	tree.UpdateRootFrequency()
	fmt.Println("构建字典树花费时间（us）：", sum)
	//tree.PrintTree()
	return tree.root
}
