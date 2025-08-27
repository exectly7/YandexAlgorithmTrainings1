package main

import (
	"fmt"
)

func main() {
	var str string
	set := make(map[string]struct{})
	for {
		_, err := fmt.Scan(&str)
		if err != nil {
			break
		}
		set[str] = struct{}{}
	}

	fmt.Println(len(set))
}
