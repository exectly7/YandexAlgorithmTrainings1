package main

import (
	"fmt"
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
	fmt.Println(solve(numbers))
}

func solve(a []int) int {
	res := 1
	mx := 0
	mxi := 0
	for i, v := range a {
		if v > mx {
			mx = v
			mxi = i
		}
	}
	vmx := 0
	if mxi == len(a)-1 {
		return 0
	}
	for i := mxi + 1; i < len(a)-1; i++ {
		if a[i]%10 == 5 && a[i+1] < a[i] && a[i] > vmx {
			vmx = a[i]
		}
	}
	if vmx == 0 {
		return 0
	}
	for _, v := range a {
		if v > vmx {
			res++
		}
	}
	return res
}
