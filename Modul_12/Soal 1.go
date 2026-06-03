package main

import "fmt"

const NCALON int = 20

type arrKandidat [NCALON]int

func seqSearch(T arrKandidat, n int, X int) int {
	var found int = -1
	var j int = 0
	for j < n && found == -1 {
		if T[j] == X {
			found = j
		}
		j = j + 1
	}
	return found
}

func main() {
	var daftar arrKandidat
	var rekap [NCALON]int
	var suara, suaraMasuk, suaraSah int

	for i := 0; i < NCALON; i++ {
		daftar[i] = i + 1
	}

	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}

		suaraMasuk++
		idx := seqSearch(daftar, NCALON, suara)

		if idx != -1 {
			suaraSah++
			rekap[idx]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	for i := 0; i < NCALON; i++ {
		if rekap[i] > 0 {
			fmt.Printf("%d: %d\n", daftar[i], rekap[i])
		}
	}
}