package main

import (
	"fmt"
)

type Power2 int

func main() {
	const (
		p4_0 Power2 = 1 << iota
		_
		p4_1
		_
		p4_2
	)
	fmt.Println("4^0: ", p4_0)
	fmt.Println("4^1: ", p4_1)
	fmt.Println("4^2: ", p4_2)
}
