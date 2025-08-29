package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	words := make(map[string]int)
	var word string
	for {
		_, err := fmt.Fscan(in, &word)
		if err != nil {
			break
		}
		words[word]++
	}
	var candidates []string
	max := 0
	for k, v := range words {
		if v > max {
			candidates = make([]string, 0)
			candidates = append(candidates, k)
			max = v
		}
		if v == max {
			candidates = append(candidates, k)
		}
	}
	sort.Strings(candidates)
	fmt.Println(candidates[0])
}
