package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scan(&input)
	newNumber := toNumber(input)
	for i := 0; i < 3; i++ {
		fmt.Scan(&input)
		number := toNumber(input)
		if number != newNumber {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}

func toNumber(s string) (res int) {
	for _, v := range s {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			continue
		}
		res *= 10
		res += n
	}
	if res < 10_000_000 {
		return 495*10_000_000 + res
	} else if res > 10_000_000_000 {
		return res - res/10_000_000_000*10_000_000_000
	} else {
		return
	}
}
