package main

import (
	"fmt"
	"math"
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
	fmt.Scan(&num)
	fmt.Println(solve(numbers, num))
}

func solve(a []int, target int) int {
	dif := math.MaxFloat64
	ansInd := 0
	for i := 0; i < len(a); i++ {
		if math.Abs(float64(target-a[i])) < dif {
			dif = math.Abs(float64(target - a[i]))
			ansInd = i
		}
	}
	return a[ansInd]
}
