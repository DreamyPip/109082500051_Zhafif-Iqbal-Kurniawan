package main

import "fmt"

func main() {
	var timA, timB string
	var poinA, poinB int

	fmt.Print("Klub A : ")
	fmt.Scan(&timA)
	fmt.Print("Klub B : ")
	fmt.Scan(&timB)

	daftarPemenang := []string{}
	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&poinA, &poinB)

		if poinA < 0 || poinB < 0 {
			break
		}

		if poinA > poinB {
			daftarPemenang = append(daftarPemenang, timA)
		} else if poinA < poinB {
			daftarPemenang = append(daftarPemenang, timB)
		} else {
			daftarPemenang = append(daftarPemenang, "Draw")
		}
	}

	for j := 0; j < len(daftarPemenang); j++ {
		fmt.Printf("Hasil %d : %s\n", j+1, daftarPemenang[j])
	}
	fmt.Println("Pertandingan selesai")
}