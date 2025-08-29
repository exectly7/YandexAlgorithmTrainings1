package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Print(words[word])
		fmt.Print(" ")
		words[word]++
	}
}
