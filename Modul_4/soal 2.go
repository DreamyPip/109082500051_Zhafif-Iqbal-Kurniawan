package main

import "fmt"

func hitungSkor(s, w *int) {
	*s = 0
	*w = 0
	var t int
	for i := 1; i <= 8; i++ {
		fmt.Scan(&t)
		if t < 301 {
			*s++
			*w += t
		}
	}
}

func main() {
	var nama, juara string
	var s_max, w_min int
	awal := true

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		var s, w int
		hitungSkor(&s, &w)

		if awal || s > s_max || (s == s_max && w < w_min) {
			s_max = s
			w_min = w
			juara = nama
			awal = false
		}
	}

	fmt.Println(juara, s_max, w_min)
}
