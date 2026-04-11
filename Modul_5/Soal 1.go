package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", fibonaci(i))
	}
}

func fibonaci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonaci(n-1) + fibonaci(n-2)
}
