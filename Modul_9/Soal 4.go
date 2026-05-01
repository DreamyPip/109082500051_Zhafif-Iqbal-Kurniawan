package main

import "fmt"

const MAX int = 127

type daftarKarakter [MAX]rune

func isi(t *daftarKarakter, jumlah *int) {
	var huruf rune
	*jumlah = 0
	for *jumlah < MAX {
		fmt.Scanf("%c", &huruf)
		if huruf == '.' {
			break
		}
		if huruf != ' ' && huruf != '\n' && huruf != '\r' {
			t[*jumlah] = huruf
			*jumlah++
		}
	}
}

func tampil(t daftarKarakter, jumlah int) {
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balik(t *daftarKarakter, jumlah int) {
	for i := 0; i < jumlah/2; i++ {
		temp := t[i]
		t[i] = t[jumlah-1-i]
		t[jumlah-1-i] = temp
	}
}

func cekPalindrom(t daftarKarakter, jumlah int) bool {
	for i := 0; i < jumlah/2; i++ {
		if t[i] != t[jumlah-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var data daftarKarakter
	var n int

	fmt.Print("Teks : ")
	isi(&data, &n)

	isPalindrom := cekPalindrom(data, n)

	fmt.Print("Reverse teks : ")
	balik(&data, n)
	tampil(data, n)

	fmt.Printf("Palindrom : %v\n", isPalindrom)
}