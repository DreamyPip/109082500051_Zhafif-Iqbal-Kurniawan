package main

import "fmt"

func pangkat(n, x int) int {
	if x == 0 {
		return 1
	}
	return n * pangkat(n, x-1)
}

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	fmt.Print(pangkat(n, x))
}
