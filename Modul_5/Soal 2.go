package main

import "fmt"

func star(n int) {
	if n != 0 {
		fmt.Print("*")
		star(n - 1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		star(i)
		fmt.Println()
	}
}
