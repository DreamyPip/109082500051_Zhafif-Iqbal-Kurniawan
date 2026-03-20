package main

import "fmt"

func main() {
	var (
		wrna1, wrna2, wrna3, wrna4 string
		temp                       bool
	)
	fmt.Println("Urutkan warna sesuai rules")
	temp = true
	x := 1
	for x <= 5 {
		fmt.Print("Percobaan ", x, " : ")
		fmt.Scan(&wrna1, &wrna2, &wrna3, &wrna4)
		if wrna1 != "merah" || wrna2 != "kuning" || wrna3 != "hijau" || wrna4 != "ungu" {
			temp = false
		}
		x++
	}
	fmt.Println("Berhasil :", temp)
}
