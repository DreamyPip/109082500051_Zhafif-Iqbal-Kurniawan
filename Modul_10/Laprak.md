# <h1 align="center">Laporan Praktikum Modul 10</h1>
<p align="center">[Zhafif Iqbal Kurniawan] - [109082500051]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

const NMAX int = 1000

type tabBerat [NMAX]float64

func cariTerkecil(berat tabBerat, n int) float64 {
	min := berat[0]
	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
	}
	return min
}

func cariTerbesar(berat tabBerat, n int) float64 {
	max := berat[0]
	for i := 1; i < n; i++ {
		if berat[i] > max {
			max = berat[i]
		}
	}
	return max
}

func main() {
	var n int
	var berat tabBerat

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	if n > 0 {
		min := cariTerkecil(berat, n)
		max := cariTerbesar(berat, n)

		fmt.Println(min, max)
	}
}

```

#### soal2.go

```go
package main

import "fmt"

const NMAX int = 1000

type tabBerat [NMAX]float64

func hitungBeratWadah(berat tabBerat, x, y int, wadah *tabBerat, jumlahWadah *int) {
	*jumlahWadah = 0
	for i := 0; i < x; i += y {
		var total float64 = 0
		for j := 0; j < y && i+j < x; j++ {
			total += berat[i+j]
		}
		wadah[*jumlahWadah] = total
		*jumlahWadah++
	}
}

func hitungRataRata(wadah tabBerat, jumlahWadah int) float64 {
	var total float64 = 0
	for i := 0; i < jumlahWadah; i++ {
		total += wadah[i]
	}
	return total / float64(jumlahWadah)
}

func main() {
	var x, y int
	var berat, wadah tabBerat
	var jumlahWadah int

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	if x > 0 && y > 0 {
		hitungBeratWadah(berat, x, y, &wadah, &jumlahWadah)

		for i := 0; i < jumlahWadah; i++ {
			fmt.Print(wadah[i], " ")
		}
		fmt.Println()

		rataRata := hitungRataRata(wadah, jumlahWadah)
		fmt.Println(rataRata)
	}
}

```

#### soal3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	
	if n > 0 {
		*bMin = arrBerat[0]
		*bMax = arrBerat[0]
		for i := 1; i < n; i++ {
			if arrBerat[i] < *bMin {
				*bMin = arrBerat[i]
			}
			if arrBerat[i] > *bMax {
				*bMax = arrBerat[i]
			}
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0
	if n == 0 {
		return 0
	}
	
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var arr arrBalita
	var min, max float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	if n > 0 {
		hitungMinMax(arr, n, &min, &max)
		rata := rerata(arr, n)

		fmt.Printf("Berat balita minimum: %.2f kg\n", min)
		fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
		fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
	}
}

```

[penjelasan]
<h2>Soal 1</h2>
<br>
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan dan Println</p>
<p>const : kata kunci untuk mendeklarasikan nilai tetap (konstanta) yang tidak bisa diubah nilainya</p>
<p>type : kata kunci untuk membuat tipe data bentukan atau alias baru</p>
<p>tabBerat [NMAX]float64 : tipe data array buatan bernama tabBerat yang bisa menyimpan angka desimal hingga batas NMAX (1000 elemen)</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func main() : bisa diartikan sebagai “fungsi utama” tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>n, berat, min, max : variabel untuk menyimpan jumlah data (n), kumpulan data array (berat), serta nilai terkecil (min) dan terbesar (max)</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>float64 : tipe data untuk bilangan desimal</p>
<p>scan : untuk memasukan data di terminal</p>
<p>println : untuk menampilkan hasil ke layar dan secara otomatis membuat baris baru</p>
<p>for : perintah perulangan untuk mengeksekusi blok kode berkali-kali</p>
<p>:= : operator untuk mendeklarasikan sekaligus mengisi nilai variabel baru secara singkat</p>
<p>if : perintah kondisional untuk mengecek sebuah syarat (misalnya apakah data lebih kecil/besar)</p>
<p>< : operator "kurang dari"</p>
<p>> : operator "lebih dari"</p>
<p>return : perintah untuk mengembalikan nilai hasil perhitungan dari dalam fungsi ke pemanggilnya</p>
<p>dalam code di atas terdapat variabel n yang diinputkan oleh user untuk menentukan seberapa banyak data berat yang ingin diproses. Program kemudian menjalankan perulangan for untuk meminta input nilai berat satu per satu dan menyimpannya ke dalam array bernama berat.</p>
<p>Setelah semua data dimasukkan, program mengecek apakah nilai n > 0 untuk memastikan ada data yang bisa diproses. Jika ada, program memanggil dua fungsi terpisah yaitu cariTerkecil dan cariTerbesar.</p>
<p>Di dalam kedua fungsi tersebut, program mengambil elemen pertama dari array sebagai nilai referensi sementara. Program kemudian melakukan perulangan untuk membandingkan referensi tersebut dengan elemen array lainnya. Jika ditemukan nilai yang lebih kecil (pada fungsi cariTerkecil) atau lebih besar (pada fungsi cariTerbesar), maka nilai referensi sementara akan diperbarui.</p>
<p>Setelah seluruh data dalam array selesai dibandingkan, kedua fungsi tersebut mengembalikan hasil akhirnya (return) ke fungsi utama (main), lalu program mencetak nilai terkecil dan terbesar tersebut ke layar.</p>
##### Output
<img width="1920" height="1080" alt="Soal 1" src="https://github.com/user-attachments/assets/643956ad-f599-4ed4-86a6-9829aa88f921" />


<br>
<br>
[penjelasan]
<h2>Soal 2</h2>
<br>
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan, Print, dan Println</p>
<p>const : kata kunci untuk mendeklarasikan nilai tetap (konstanta) yang tidak bisa diubah nilainya</p>
<p>type : kata kunci untuk membuat tipe data bentukan atau alias baru</p>
<p>tabBerat [NMAX]float64 : tipe data array buatan bernama tabBerat yang bisa menyimpan angka desimal hingga batas NMAX (1000 elemen)</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func main() : bisa diartikan sebagai “fungsi utama” tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>x, y, berat, wadah, jumlahWadah : variabel untuk menyimpan total data (x), isi per kelompok/wadah (y), array data asli (berat), array hasil pengelompokan (wadah), dan total wadah (jumlahWadah)</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>float64 : tipe data untuk bilangan desimal</p>
<p>scan : untuk memasukan data di terminal</p>
<p>print : untuk menampilkan hasil ke layar tanpa membuat baris baru (menyamping)</p>
<p>println : untuk menampilkan hasil ke layar dan secara otomatis membuat baris baru</p>
<p>for : perintah perulangan untuk mengeksekusi blok kode berkali-kali</p>
<p>:= : operator untuk mendeklarasikan sekaligus mengisi nilai variabel baru secara singkat</p>
<p>+= : operator penjumlahan yang langsung memperbarui dan menambahkan nilai ke variabel itu sendiri</p>
<p>++ : operator increment untuk menambah nilai variabel sebanyak 1</p>
<p>if : perintah kondisional untuk mengecek sebuah syarat</p>
<p>< : operator "kurang dari"</p>
<p>&& : operator logika "DAN" (AND), yang berarti kedua syarat di kiri dan kanannya harus terpenuhi</p>
<p>* (pointer) : tanda asterisk pada parameter fungsi (seperti *tabBerat dan *int) digunakan sebagai pointer, agar fungsi tersebut bisa memodifikasi memori variabel aslinya secara langsung</p>
<p>& (ampersand) : tanda dan pada pemanggilan fungsi atau scan digunakan untuk mengirimkan alamat referensi memori dari variabel tersebut</p>
<p>return : perintah untuk mengembalikan nilai hasil perhitungan dari dalam fungsi ke pemanggilnya</p>
<p>dalam code di atas terdapat variabel x (jumlah total data) dan y (kapasitas per wadah) yang diinputkan oleh user. Program kemudian menjalankan perulangan untuk meminta input nilai berat satu per satu dan menyimpannya di array berat.</p>
<p>Setelah memastikan nilai x dan y lebih dari 0, program memanggil fungsi hitungBeratWadah. Di fungsi ini, program melakukan perulangan untuk mengelompokkan data berat secara bertahap sebanyak y elemen. Total penjumlahan dari setiap kelompok tersebut disimpan ke dalam array baru bernama wadah. Karena menggunakan pointer (*), semua perubahan pada array wadah dan variabel jumlahWadah otomatis tersimpan dan berlaku di fungsi utama.</p>
<p>Setelah perhitungan wadah selesai, program utama (main) akan mencetak total berat dari masing-masing wadah tersebut secara menyamping dengan pemisah spasi.</p>
<p>Terakhir, program memanggil fungsi hitungRataRata. Fungsi ini akan menjumlahkan semua berat di dalam array wadah, lalu membaginya dengan jumlahWadah untuk mendapatkan nilai rata-rata keseluruhan. Nilai rata-rata tersebut dikembalikan ke fungsi utama dan dicetak ke layar.</p>
##### Output
<img width="1920" height="1080" alt="Soal 2" src="https://github.com/user-attachments/assets/4b40742c-8588-4326-b1bf-fc6d9cd549f7" />

<br>
<br>

[penjelasan]
<h2>Soal 3</h2>
<br>
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan, Print, dan Printf</p>
<p>type : kata kunci untuk membuat tipe data bentukan atau alias baru</p>
<p>arrBalita [100]float64 : tipe data array buatan bernama arrBalita yang bisa menyimpan angka desimal hingga batas 100 elemen</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func main() : bisa diartikan sebagai “fungsi utama” tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>n, arr, min, max, total : variabel untuk menyimpan banyak data (n), kumpulan data berat (arr), nilai terkecil (min), terbesar (max), dan total penjumlahan (total)</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>float64 : tipe data untuk bilangan desimal</p>
<p>print : untuk menampilkan teks ke layar tanpa baris baru</p>
<p>printf : untuk menampilkan teks dan hasil dengan format tertentu (contohnya %.2f untuk membatasi angka desimal menjadi 2 angka di belakang koma)</p>
<p>scan : untuk memasukan data di terminal</p>
<p>for : perintah perulangan untuk mengeksekusi blok kode berkali-kali</p>
<p>:= : operator untuk mendeklarasikan sekaligus mengisi nilai variabel baru secara singkat</p>
<p>+= : operator penjumlahan yang langsung memperbarui dan menambahkan nilai ke variabel itu sendiri</p>
<p>++ : operator increment untuk menambah nilai variabel sebanyak 1</p>
<p>if : perintah kondisional untuk mengecek sebuah syarat</p>
<p>< : operator "kurang dari"</p>
<p>> : operator "lebih dari"</p>
<p>== : operator "sama dengan" (digunakan untuk mengecek apakah n bernilai 0)</p>
<p>* (pointer) : tanda asterisk pada parameter fungsi (seperti *float64) digunakan sebagai pointer, agar fungsi tersebut bisa memodifikasi memori variabel aslinya secara langsung</p>
<p>& (ampersand) : tanda dan pada pemanggilan fungsi atau scan digunakan untuk mengirimkan alamat referensi memori dari variabel tersebut</p>
<p>return : perintah untuk mengembalikan nilai hasil perhitungan dari dalam fungsi ke pemanggilnya</p>
<p>dalam code di atas terdapat variabel n yang diinputkan oleh user untuk menentukan berapa banyak data berat balita yang akan diproses. Program kemudian menjalankan perulangan for untuk meminta user memasukkan berat balita satu per satu dan menyimpannya di dalam array bernama arr.</p>
<p>Setelah data terkumpul dan dipastikan n > 0, program memanggil fungsi hitungMinMax. Karena fungsi ini menggunakan referensi pointer (&min, &max), fungsi ini dapat langsung memodifikasi dan mencari dua nilai sekaligus (nilai terkecil dan terbesar) dari dalam array, lalu menyimpannya ke variabel asli di fungsi utama.</p>
<p>Selanjutnya, program memanggil fungsi rerata. Fungsi ini menjumlahkan seluruh isi array berat balita, lalu membaginya dengan jumlah data (n) untuk mendapatkan nilai rata-rata, dan mengembalikan (return) nilai tersebut.</p>
<p>Terakhir, program utama mencetak nilai minimum, maksimum, dan rata-rata ke layar. Penggunaan Printf dengan format %.2f berfungsi agar angka desimal yang panjang dipotong dengan rapi menjadi dua angka di belakang koma saja.</p>
##### Output
<img width="1920" height="1080" alt="Soal 3" src="https://github.com/user-attachments/assets/441aeeca-ae67-47c9-b859-bc4c488eb264" />
