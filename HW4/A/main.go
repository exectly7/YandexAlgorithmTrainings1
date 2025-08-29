package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var num int
	fmt.Fscan(in, &num)

	sinons := make(map[string]string)

	for i := 0; i < num; i++ {
		var in1, in2 string
		fmt.Fscan(in, &in1, &in2)
		sinons[in1] = in2
		sinons[in2] = in1
	}

	var word string
	fmt.Fscan(in, &word)
	fmt.Println(sinons[word])
}
