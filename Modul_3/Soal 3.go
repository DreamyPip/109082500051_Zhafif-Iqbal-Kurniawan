package main

import (
	"fmt"
	"math"
)

func dist(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func inside(px, py, r, x, y float64) bool {
	return dist(px, py, x, y) <= r
}

func main() {
	var (
		x1, y1, r1 int
		x2, y2, r2 int
		tx, ty     int
	)

	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&tx, &ty)

	in1 := inside(float64(x1), float64(y1), float64(r1), float64(tx), float64(ty))
	in2 := inside(float64(x2), float64(y2), float64(r2), float64(tx), float64(ty))

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
