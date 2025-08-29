package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	n := 0
	fmt.Fscan(in, &n)
	var w, h int
	blocks := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &w, &h)
		if _, ok := blocks[w]; !ok {
			blocks[w] = h
		} else {
			blocks[w] = max(blocks[w], h)
		}
	}
	c := 0
	for i := range blocks {
		c += blocks[i]
	}
	fmt.Println(c)
}
