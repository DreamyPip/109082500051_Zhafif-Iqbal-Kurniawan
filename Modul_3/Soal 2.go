package main

import "fmt"

func kuadrat(angka int) int {
	hasilKuadrat := angka * angka
	return hasilKuadrat
}

func kurang(angka int) int {
	hasilKurang := angka - 2
	return hasilKurang
}

func tambah(angka int) int {
	hasilTambah := angka + 1
	return hasilTambah
}

func main() {
	var x1, x2, x3 int
	fmt.Scan(&x1, &x2, &x3)

	fmt.Println(kuadrat(kurang(tambah(x1))))
	fmt.Println(kurang(tambah(kuadrat(x2))))
	fmt.Println(tambah(kuadrat(kurang(x3))))
}
