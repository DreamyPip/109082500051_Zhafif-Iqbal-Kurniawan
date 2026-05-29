package main

import "fmt"

const NMAX int = 1000

type tabBerat [NMAX]float64

func cariTerkecil(berat tabBerat, n int) float64 {
	min := berat[0]
	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
	}
	return min
}

func cariTerbesar(berat tabBerat, n int) float64 {
	max := berat[0]
	for i := 1; i < n; i++ {
		if berat[i] > max {
			max = berat[i]
		}
	}
	return max
}

func main() {
	var n int
	var berat tabBerat

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	if n > 0 {
		min := cariTerkecil(berat, n)
		max := cariTerbesar(berat, n)

		fmt.Println(min, max)
	}
}