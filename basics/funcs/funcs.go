package main

import "fmt"

func main() {
	fmt.Println(exp(5, 3))
	fmt.Println(exp(2, 0))
	fmt.Println(exp(2, 7))
}

func exp(e uint64, x uint64) uint64 {
	for ; x > 0; x-- {
		e = e * x
	}

	return e
}
