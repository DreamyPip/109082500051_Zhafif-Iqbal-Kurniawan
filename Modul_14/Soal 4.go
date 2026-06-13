package main

import "fmt"

const NMAX = 1000

type tabInt [NMAX]int

func insertionSort(arr *tabInt, n int) {
	for i := 1; i < n; i++ {
		key := (*arr)[i]
		j := i - 1
		for j >= 0 && (*arr)[j] > key {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = key
	}
}

func cekJarak(arr tabInt, n int) {
	if n > 1 {
		diff := arr[1] - arr[0]
		status := 1

		for i := 1; i < n-1; i++ {
			if arr[i+1]-arr[i] != diff {
				status = 0
			}
		}

		if status == 1 {
			fmt.Printf("Data berjarak %d\n", diff)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}

func main() {
	var arr tabInt
	var n int
	var num int

	fmt.Scan(&num)
	for num >= 0 {
		arr[n] = num
		n++
		fmt.Scan(&num)
	}

	insertionSort(&arr, n)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	cekJarak(arr, n)
}
