package main

import "fmt"

const NMAX = 1000000

var arr [NMAX]int

func selectionSort(n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		temp := arr[i]
		arr[i] = arr[idxMin]
		arr[idxMin] = temp
	}
}

func main() {
	var n, m int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		selectionSort(m)

		for j := 0; j < m; j++ {
			fmt.Print(arr[j])
			if j < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}