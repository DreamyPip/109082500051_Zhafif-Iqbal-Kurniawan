package main

import (
	"fmt"
	"math"
)

const KAPASITAS int = 100

func main() {
	var n, x, hapusIndex, cariBilangan int
	var data [KAPASITAS]int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	for i := 1; i < n; i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	for i := 0; i < n; i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	fmt.Scan(&x)
	if x > 0 {
		for i := 0; i < n; i++ {
			if i%x == 0 {
				fmt.Print(data[i], " ")
			}
		}
		fmt.Println()
	}

	fmt.Scan(&hapusIndex)
	if hapusIndex >= 0 && hapusIndex < n {
		for i := hapusIndex; i < n-1; i++ {
			data[i] = data[i+1]
		}
		n--
	}
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	var total float64
	for i := 0; i < n; i++ {
		total += float64(data[i])
	}
	rerata := total / float64(n)
	fmt.Printf("%.2f\n", rerata)

	var jumlahKuadratSelisih float64
	for i := 0; i < n; i++ {
		selisih := float64(data[i]) - rerata
		jumlahKuadratSelisih += selisih * selisih
	}
	stdev := math.Sqrt(jumlahKuadratSelisih / float64(n))
	fmt.Printf("%.2f\n", stdev)

	fmt.Scan(&cariBilangan)
	frekuensi := 0
	for i := 0; i < n; i++ {
		if data[i] == cariBilangan {
			frekuensi++
		}
	}
	fmt.Println(frekuensi)
}