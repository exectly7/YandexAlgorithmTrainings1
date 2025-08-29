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
	durability := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Fscan(in, &durability[i])
	}
	fmt.Fscan(in, &num)
	var key int
	for i := 0; i < num; i++ {
		fmt.Fscan(in, &key)
		durability[key-1]--
	}
	for v := range durability {
		if durability[v] < 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
