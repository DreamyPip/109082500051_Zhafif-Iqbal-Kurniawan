package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y float64
}

type lingkaran struct {
	pusat  titik
	radius float64
}

func jarak(p, q titik) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

func didalam(c lingkaran, p titik) bool {
	return jarak(p, c.pusat) <= c.radius
}

func main() {
	var daftarL [2]lingkaran
	var t titik

	for i := 0; i < 2; i++ {
		fmt.Scan(&daftarL[i].pusat.x, &daftarL[i].pusat.y, &daftarL[i].radius)
	}

	fmt.Scan(&t.x, &t.y)

	cek1 := didalam(daftarL[0], t)
	cek2 := didalam(daftarL[1], t)

	if cek1 && cek2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if cek1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if cek2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}