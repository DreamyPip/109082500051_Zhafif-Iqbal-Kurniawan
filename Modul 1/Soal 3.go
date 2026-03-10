package main

import "fmt"

func main() {
	var g, kg, gr, tambahan int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&g)
	kg = g / 1000
	gr = g % 1000
	hasil := kg * 10000
	fmt.Println("Detai berat: Rp.", kg, "kg +", gr, "gr")
	if gr >= 500 && kg <= 10 {
		tambahan = gr * 5
		fmt.Println("Detail biaya: Rp.", hasil, " + Rp.", tambahan)
		hasil = hasil + tambahan
		fmt.Print("Total biaya: ", hasil)
	} else if gr < 500 && kg <= 10 {
		tambahan = gr * 15
		fmt.Println("Detail biaya: Rp.", hasil, " + Rp.", tambahan)
		hasil = hasil + tambahan
		fmt.Print("Total biaya: ", hasil)
	} else if kg > 10 {
		fmt.Println("Detail biaya: Rp.", hasil, " + Rp.", tambahan)
		fmt.Print("Total biaya: ", hasil)
	}
}
