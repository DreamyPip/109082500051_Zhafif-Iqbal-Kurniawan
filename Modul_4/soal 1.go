package main

import (
	"fmt"
)

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}
func mutasi(n, a int, hasil *int) {
	var x, y int
	if n >= a {
		faktorial(n, &x)
		faktorial(n-a, &y)

		*hasil = x / y
	} else {
		*hasil = 0
	}
}

func kombinasi(n, a int, hasil *int) {
	var x, y, z int
	if n >= a {
		faktorial(n, &z)
		faktorial(a, &x)
		faktorial(n-a, &y)

		*hasil = z / (x * y)
	} else {
		*hasil = 0
	}
}

func main() {
	var a, b, c, d, e, f int
	fmt.Scan(&a, &b, &c, &d)
	mutasi(a, c, &e)
	kombinasi(a, c, &f)
	fmt.Println(e, f)

	mutasi(b, d, &e)
	kombinasi(b, d, &f)
	fmt.Println(e, f)
}
