package main

import "fmt"

func ganjil(n int, x int) {
	if x <= n {
		fmt.Print(x, " ")
		ganjil(n, x+2)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	ganjil(n, 1)
}
