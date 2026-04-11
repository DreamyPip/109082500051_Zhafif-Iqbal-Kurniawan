package main

import "fmt"

func bagi(n int, i int) {
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	if n >= i {
		bagi(n, i+1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	bagi(n, 1)
}
