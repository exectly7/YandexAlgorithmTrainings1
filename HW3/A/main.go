package main

import (
	"fmt"
)

func main() {
	var num int
	numbers := make(map[int]struct{})
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		numbers[num] = struct{}{}
	}

	fmt.Println(len(numbers))
}
