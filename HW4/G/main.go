package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	bank := make(map[string]int)

	var op, name1, name2 string
	var amount int
	for {
		_, err := fmt.Fscan(in, &op)
		if err != nil {
			break
		}
		switch op {
		case "DEPOSIT":
			fmt.Fscan(in, &name1, &amount)
			Deposit(bank, name1, amount)
		case "WITHDRAW":
			fmt.Fscan(in, &name1, &amount)
			Withdraw(bank, name1, amount)
		case "BALANCE":
			fmt.Fscan(in, &name1)
			fmt.Println(Balance(bank, name1))
		case "INCOME":
			fmt.Fscan(in, &amount)
			Income(bank, amount)
		case "TRANSFER":
			fmt.Fscan(in, &name1, &name2, &amount)
			Transfer(bank, name1, name2, amount)
		}
	}
}

func Deposit(bank map[string]int, name string, amount int) {
	bank[name] += amount
}

func Withdraw(bank map[string]int, name string, amount int) {
	bank[name] -= amount
}

func Balance(bank map[string]int, name string) string {
	if _, ok := bank[name]; ok {
		return strconv.Itoa(bank[name])
	}
	return "ERROR"
}

func Income(bank map[string]int, amount int) {
	for v := range bank {
		if bank[v] > 0 {
			bank[v] += int(float64(bank[v]) * float64(amount) / 100.0)
		}
	}
}

func Transfer(bank map[string]int, name1, name2 string, amount int) {
	bank[name1] -= amount
	bank[name2] += amount
}
