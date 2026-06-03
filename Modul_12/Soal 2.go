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
	var rekap [NCALON + 1]int
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
			rekap[daftar[idx]]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	var max1, max2 int = -1, -1
	var id1, id2 int = 0, 0

	for i := 1; i <= NCALON; i++ {
		if rekap[i] > max1 {
			max2 = max1
			id2 = id1
			max1 = rekap[i]
			id1 = i
		} else if rekap[i] > max2 {
			max2 = rekap[i]
			id2 = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", id1)
	fmt.Printf("Wakil ketua: %d\n", id2)
}