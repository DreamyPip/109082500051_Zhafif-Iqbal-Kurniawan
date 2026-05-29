package main

import "fmt"

const NMAX int = 1000

type tabBerat [NMAX]float64

func hitungBeratWadah(berat tabBerat, x, y int, wadah *tabBerat, jumlahWadah *int) {
	*jumlahWadah = 0
	for i := 0; i < x; i += y {
		var total float64 = 0
		for j := 0; j < y && i+j < x; j++ {
			total += berat[i+j]
		}
		wadah[*jumlahWadah] = total
		*jumlahWadah++
	}
}

func hitungRataRata(wadah tabBerat, jumlahWadah int) float64 {
	var total float64 = 0
	for i := 0; i < jumlahWadah; i++ {
		total += wadah[i]
	}
	return total / float64(jumlahWadah)
}

func main() {
	var x, y int
	var berat, wadah tabBerat
	var jumlahWadah int

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	if x > 0 && y > 0 {
		hitungBeratWadah(berat, x, y, &wadah, &jumlahWadah)

		for i := 0; i < jumlahWadah; i++ {
			fmt.Print(wadah[i], " ")
		}
		fmt.Println()

		rataRata := hitungRataRata(wadah, jumlahWadah)
		fmt.Println(rataRata)
	}
}