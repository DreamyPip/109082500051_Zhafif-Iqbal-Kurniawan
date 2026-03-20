package main

import "fmt"

func hitungFaktorial(x int) int {
	var hasil int = 1
	var j int
	for j = 1; j <= x; j++ {
		hasil = hasil * j
	}
	return hasil
}

func hitungPermutasi(total, ambil int) int {
	return hitungFaktorial(total) / hitungFaktorial(total-ambil)
}

func hitungKombinasi(total, ambil int) int {
	return hitungFaktorial(total) / (hitungFaktorial(ambil) * hitungFaktorial(total-ambil))
}

func main() {
	var n1, n2, r1, r2 int
	fmt.Scan(&n1, &n2, &r1, &r2)

	if n1 >= r1 {
		fmt.Print(hitungPermutasi(n1, r1), " ")
		fmt.Println(hitungKombinasi(n1, r1))
	}

	if n2 >= r2 {
		fmt.Print(hitungPermutasi(n2, r2), " ")
		fmt.Println(hitungKombinasi(n2, r2))
	}
}
