package main

import (
	"fmt"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	field := make([][]int, n)
	for i := 0; i < n; i++ {
		field[i] = make([]int, m)
	}
	var x, y int
	for i := 0; i < k; i++ {
		fmt.Scan(&x, &y)
		field[x-1][y-1] = -1
	}
	res := solve(field)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if res[i][j] == -1 {
				fmt.Print("*")
			} else {
				fmt.Print(res[i][j])
			}
			if j != m-1 {
				fmt.Print(" ")
			}
		}
		if i != n-1 {
			fmt.Print("\n")
		}
	}
}

func solve(field [][]int) [][]int {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if i != 0 && field[i-1][j] == -1 && field[i][j] != -1 {
				field[i][j]++
			}
			if j != 0 && field[i][j-1] == -1 && field[i][j] != -1 {
				field[i][j]++
			}
			if i != len(field)-1 && field[i+1][j] == -1 && field[i][j] != -1 {
				field[i][j]++
			}
			if j != len(field[i])-1 && field[i][j+1] == -1 && field[i][j] != -1 {
				field[i][j]++
			}
			if i != 0 && j != 0 && field[i-1][j-1] == -1 && field[i][j] != -1 {
				field[i][j]++
			}
			if i != 0 && j != len(field[i])-1 && field[i-1][j+1] == -1 && field[i][j] != -1 {
				field[i][j]++
			}
			if i != len(field)-1 && j != 0 && field[i+1][j-1] == -1 && field[i][j] != -1 {
				field[i][j]++
			}
			if i != len(field)-1 && j != len(field[i])-1 && field[i+1][j+1] == -1 && field[i][j] != -1 {
				field[i][j]++
			}
		}
	}
	return field
}
