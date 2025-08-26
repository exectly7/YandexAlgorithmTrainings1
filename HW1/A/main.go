package main

import (
	"fmt"
)

func main() {
	var tRoom, tCond int
	var mode string
	fmt.Scan(&tRoom, &tCond, &mode)

	switch mode {
	case "freeze":
		if tCond < tRoom {
			fmt.Println(tCond)
			return
		}
		fmt.Println(tRoom)
		return
	case "heat":
		if tCond > tRoom {
			fmt.Println(tCond)
			return
		}
		fmt.Println(tRoom)
		return
	case "auto":
		fmt.Println(tCond)
		return
	case "fan":
		fmt.Println(tRoom)
		return
	default:
		return
	}
}
