package main

import "fmt"

const NMAX = 1000000

var arrGanjil [NMAX]int
var arrGenap [NMAX]int

func sortGanjil(n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arrGanjil[j] < arrGanjil[idxMin] {
				idxMin = j
			}
		}
		temp := arrGanjil[i]
		arrGanjil[i] = arrGanjil[idxMin]
		arrGanjil[idxMin] = temp
	}
}

func sortGenap(n int) {
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if arrGenap[j] > arrGenap[idxMax] {
				idxMax = j
			}
		}
		temp := arrGenap[i]
		arrGenap[i] = arrGenap[idxMax]
		arrGenap[idxMax] = temp
	}
}

func main() {
	var n, m, num int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		var nGanjil int = 0
		var nGenap int = 0

		for j := 0; j < m; j++ {
			fmt.Scan(&num)
			if num%2 != 0 {
				arrGanjil[nGanjil] = num
				nGanjil++
			} else {
				arrGenap[nGenap] = num
				nGenap++
			}
		}

		sortGanjil(nGanjil)
		sortGenap(nGenap)

		var pertama int = 1

		for j := 0; j < nGanjil; j++ {
			if pertama == 1 {
				fmt.Print(arrGanjil[j])
				pertama = 0
			} else {
				fmt.Print(" ", arrGanjil[j])
			}
		}

		for j := 0; j < nGenap; j++ {
			if pertama == 1 {
				fmt.Print(arrGenap[j])
				pertama = 0
			} else {
				fmt.Print(" ", arrGenap[j])
			}
		}
		fmt.Println()
	}
}	