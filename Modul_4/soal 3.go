package main

import "fmt"

func baris(n int) {
	fmt.Printf("%d ", n)
	for n != 1 && n > 0 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		fmt.Printf("%d ", n)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	baris(n)
}
