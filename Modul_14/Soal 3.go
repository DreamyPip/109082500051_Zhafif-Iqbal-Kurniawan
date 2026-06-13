package main

import "fmt"

const NMAX = 1000000

var arr [NMAX]int

func insertionSort(n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var n int = 0
	var num int

	fmt.Scan(&num)
	for num != -5313 {
		if num == 0 {
			if n > 0 {
				insertionSort(n)
				if n%2 != 0 {
					fmt.Println(arr[n/2])
				} else {
					fmt.Println((arr[n/2-1] + arr[n/2]) / 2)
				}
			}
		} else {
			arr[n] = num
			n++
		}
		fmt.Scan(&num)
	}
}