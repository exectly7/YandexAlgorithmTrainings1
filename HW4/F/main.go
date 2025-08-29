package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	buys := make(map[string]map[string]int)
	var buyer, item string
	var amount int
	for {
		_, err := fmt.Fscan(in, &buyer, &item, &amount)
		if err != nil {
			break
		}
		if _, ok := buys[buyer]; !ok {
			buys[buyer] = make(map[string]int)
			buys[buyer][item] += amount
		} else {
			buys[buyer][item] += amount
		}
	}
	outerKeys := make([]string, 0, len(buys))
	for k := range buys {
		outerKeys = append(outerKeys, k)
	}
	sort.Strings(outerKeys)

	for _, outer := range outerKeys {
		fmt.Println(outer + ":")
		innerKeys := make([]string, 0, len(buys[outer]))
		for k := range buys[outer] {
			innerKeys = append(innerKeys, k)
		}
		sort.Strings(innerKeys)
		for _, inner := range innerKeys {
			fmt.Printf("%s %d\n", inner, buys[outer][inner])
		}
	}
}
