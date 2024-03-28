package main

import "fmt"

var m = map[int]int{
	0: 0,
	1: 1,
	2: 1,
}

func fib(n int) int {
	if _, ok := m[n]; ok {
		return m[n]
	}

	m[n] = fib(n-1) + fib(n-2)
	return m[n]
}

func main() {
	fmt.Printf("res: %d", fib(25))
}
