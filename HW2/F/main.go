package main

import (
	"fmt"
	"slices"
)

func main() {
	var num, amount int
	fmt.Scan(&amount)
	var numbers []int
	for ; amount > 0; amount-- {
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		numbers = append(numbers, num)
	}
	num, res := solve(numbers)
	fmt.Println(num)
	for i := 0; i < num; i++ {
		fmt.Println(res[i], " ")
	}
}

func solve(a []int) (int, []int) {
	var curM []int
	j := len(a) - 1
	amount := 0
	var res []int
	for i := 0; i < j; i++ {
		curM = append(curM, a[i])
		if a[i] != a[j] {
			amount += len(curM)
			res = append(res, curM...)
			j += len(curM) - 1
			curM = make([]int, 0)
		} else {
			j--
		}
	}
	slices.Reverse(res)
	return amount, res
}
