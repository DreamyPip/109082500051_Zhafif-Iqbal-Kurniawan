# <h1 align="center">Laporan Praktikum Modul 3</h1>
<p align="center">[Zhafif Iqbal Kurniawan] - [109082500051]</p>

## Unguided
### 1. [Soal]
#### soal1.go

```go
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


```

#### soal2.go

```go
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

```

#### soal3.go

```go
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

```

[penjelasan]
<h2>Soal 1</h2>
<br>
<p>package main : ini adalah paket utama tempat kode dijalankan</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan, Print, dan Println</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func hitungFaktorial() : fungsi buatan untuk menghitung nilai faktorial dari sebuah angka menggunakan perulangan (misal 5! = 1 * 2 * 3 * 4 * 5)</p>
<p>func hitungPermutasi() : fungsi buatan untuk menghitung nilai permutasi (susunan yang memperhatikan urutan) menggunakan rumus matematika n! / (n-r)!</p>
<p>func hitungKombinasi() : fungsi buatan untuk menghitung nilai kombinasi (susunan tanpa memperhatikan urutan) menggunakan rumus matematika n! / (r! * (n-r)!)</p>
<p>func main() : fungsi utama tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>n1, n2, r1, r2 : variabel untuk menyimpan input nilai n (total elemen) dan r (elemen yang diambil) untuk dua perhitungan yang berbeda</p>
<p>hasil, j, x : variabel dalam fungsi faktorial untuk menampung hasil kali (hasil), iterator perulangan (j), dan angka batas faktorial (x)</p>
<p>total, ambil : parameter variabel di dalam fungsi permutasi dan kombinasi sebagai perwakilan nilai n (total) dan r (ambil)</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk membaca banyak data sekaligus dari input terminal</p>
<p>print / println : untuk menampilkan hasil cetak menyamping dan mencetak hasil dengan baris baru</p>
<p>for : perintah perulangan untuk mengeksekusi blok kode secara berulang</p>
<p>if : perintah kondisional untuk mengecek sebuah syarat</p>
<p>>= : operator perbandingan "lebih dari atau sama dengan" (digunakan untuk memastikan jumlah total n tidak lebih kecil dari jumlah r yang diambil)</p>
<p><= : operator perbandingan "kurang dari atau sama dengan" (sebagai batas perulangan faktorial)</p>
<p>++ : operator increment untuk menambah nilai variabel sebanyak 1</p>
<p>* / / : operator matematika dasar untuk perkalian dan pembagian</p>
<p>return : perintah untuk mengembalikan nilai hasil perhitungan ke fungsi yang memanggilnya</p>
<p>dalam code di atas, program memiliki tiga fungsi pembantu. Fungsi paling dasar adalah hitungFaktorial yang mengalikan angka secara berurutan mulai dari 1 sampai angka x. Fungsi ini kemudian dimanfaatkan oleh dua fungsi lainnya, yaitu hitungPermutasi dan hitungKombinasi, untuk menyelesaikan rumus permutasi dan kombinasi dengan memanggil hasil faktorial tersebut.</p>
<p>Di dalam program utama (main), pengguna diminta untuk memasukkan 4 angka sekaligus yang diwakili oleh variabel n1, n2, r1, dan r2. Angka-angka ini adalah data input untuk dua kasus (kasus pertama menggunakan n1 dan r1, kasus kedua menggunakan n2 dan r2).</p>
<p>Setelah input dimasukkan, program akan mengevaluasi kasus pertama dengan mengecek syarat if n1 >= r1. Pengecekan ini sangat penting karena secara logika matematika, jumlah barang yang diambil (r) tidak boleh melebihi total barang yang tersedia (n). Jika syarat ini terpenuhi, program memanggil fungsi permutasi dan kombinasi lalu mencetak hasilnya secara menyamping dengan dipisah spasi.</p>
<p>Langkah serupa juga diulangi untuk kasus kedua. Program mengecek syarat if n2 >= r2. Jika valid, program kembali menghitung dan mencetak nilai permutasi serta kombinasinya ke layar.</p>

##### Output
<img width="1920" height="1080" alt="Soal 1" src="https://github.com/user-attachments/assets/8bfc3095-5ee0-4507-829b-e9ddc48e37fc" />


[penjelasan]
<h2>Soal 2</h2>
<br>

<p>package main : ini adalah paket utama tempat kode dijalankan</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan dan Println</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func kuadrat() : fungsi buatan untuk menghitung hasil pangkat dua dari sebuah angka (angka * angka)</p>
<p>func kurang() : fungsi buatan untuk mengurangi nilai sebuah angka dengan 2 (angka - 2)</p>
<p>func tambah() : fungsi buatan untuk menambahkan nilai sebuah angka dengan 1 (angka + 1)</p>
<p>func main() : fungsi utama tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>x1, x2, x3 : variabel untuk menampung tiga angka inputan dari pengguna</p>
<p>hasilKuadrat, hasilKurang, hasilTambah : variabel di dalam masing-masing fungsi untuk menyimpan hasil perhitungan sementara</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk membaca input angka dari terminal</p>
<p>println : untuk menampilkan hasil perhitungan akhir ke layar sekaligus membuat baris baru</p>
<p>:= : operator untuk mendeklarasikan sekaligus mengisi nilai variabel baru secara singkat</p>
<p>* / - / + : operator matematika dasar untuk perkalian, pengurangan, dan penjumlahan</p>
<p>return : perintah untuk mengembalikan nilai hasil perhitungan dari dalam fungsi ke pemanggilnya</p>
<p>dalam code di atas, program memodelkan konsep komposisi fungsi matematika menggunakan tiga sub-program sederhana: kuadrat, kurang, dan tambah. Masing-masing fungsi hanya menerima satu angka, melakukan satu operasi matematika spesifik, dan mengembalikan hasilnya.</p>
<p>Di dalam program utama (main), pengguna diminta memasukkan tiga buah angka sekaligus yang akan disimpan ke dalam variabel x1, x2, dan x3. Program kemudian memproses angka-angka tersebut dengan urutan pemanggilan fungsi yang ditumpuk-tumpuk (nested calls).</p>
<p>Cara kerja pemanggilan fungsi yang bertumpuk ini selalu dieksekusi dari kurung yang paling dalam ke kurung yang paling luar. Pada baris pencetakan pertama untuk angka x1 yaitu kuadrat(kurang(tambah(x1))), program mengeksekusi fungsi tambah() terlebih dahulu. Hasil dari fungsi tambah() tidak disimpan ke variabel baru, melainkan langsung dilempar menjadi input untuk fungsi kurang(). Setelah fungsi kurang() selesai, hasilnya kembali dilempar menjadi input untuk fungsi kuadrat(). Barulah hasil akhirnya dicetak ke layar.</p>
<p>Konsep eksekusi "dari dalam ke luar" ini juga berlaku persis pada baris pencetakan kedua (angka x2) dan ketiga (angka x3). Karena susunan pemanggilan fungsinya sengaja diubah-ubah pada setiap barisnya, maka urutan perhitungannya pun menjadi berbeda dan menghasilkan output akhir yang berbeda pula untuk setiap kasus.</p>

##### Output
<img width="1920" height="1080" alt="Soal 2" src="https://github.com/user-attachments/assets/66353d80-32a6-4431-a44c-0bb540ba4604" />



[penjelasan]
<h2>Soal 3</h2>
<br>
Penjelasan Program (Pengecekan Titik Koordinat Terhadap Lingkaran)
<p>package main : ini adalah paket utama tempat kode dijalankan</p>
<p>import "fmt", "math": Perintah ini mengimpor dua paket, yaitu fmt untuk operasi input/output terminal dan math untuk melakukan operasi matematika tingkat lanjut seperti perhitungan akar kuadrat</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func dist() : fungsi buatan untuk menghitung jarak garis lurus antara dua buah titik koordinat koordinat (x1, y1) dan (x2, y2) menggunakan rumus jarak Euclidean</p>
<p>func inside() : fungsi buatan untuk mengecek apakah suatu titik target (x, y) berada di dalam lingkaran yang berpusat di (px, py) dengan jari-jari r</p>
<p>func main() : fungsi utama tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>x1, y1, r1 : variabel untuk menyimpan titik pusat koordinat x, koordinat y, serta jari-jari (radius) dari lingkaran pertama</p>
<p>x2, y2, r2 : variabel untuk menyimpan titik pusat koordinat x, koordinat y, serta jari-jari dari lingkaran kedua</p>
<p>tx, ty : variabel untuk menyimpan koordinat titik target yang akan dicek posisinya</p>
<p>in1, in2 : variabel yang menyimpan hasil pengecekan (apakah titik berada di dalam lingkaran 1 dan/atau lingkaran 2)</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>float64 : tipe data untuk bilangan desimal berpresisi tinggi (wajib digunakan saat menggunakan operasi dari paket math)</p>
<p>bool : tipe data boolean yang hanya memiliki dua kemungkinan nilai, yaitu true (benar) atau false (salah)</p>
<p>math.Sqrt : fungsi bawaan dari paket math untuk menghitung akar kuadrat (square root)</p>
<p>float64() : perintah untuk mengubah (casting) tipe data suatu variabel dari bilangan bulat (int) menjadi desimal (float64) agar bisa dikalkulasi oleh fungsi jarak</p>
<p>scan / println : untuk membaca input dari terminal dan mencetak hasil ke layar dengan baris baru</p>
<p>if / else if / else : perintah kondisional untuk mengecek syarat beruntun dan menentukan kesimpulan akhir</p>
<p>:= : operator untuk mendeklarasikan sekaligus mengisi nilai variabel baru secara singkat</p>
<p>&& : operator logika "DAN" (AND) yang mengharuskan kedua kondisi bernilai benar (true)</p>
<p><= : operator perbandingan "kurang dari atau sama dengan" (jika jarak titik ke pusat kurang dari/sama dengan jari-jari, berarti titik ada di dalam lingkaran)</p>
<p>return : perintah untuk mengembalikan nilai perhitungan (jarak) atau nilai logika (true/false) ke pemanggil fungsi</p>
<p>dalam code di atas, program meminta pengguna untuk memasukkan 8 angka sekaligus. Tiga angka pertama mendefinisikan lingkaran kesatu (x1, y1, r1), tiga angka kedua mendefinisikan lingkaran kedua (x2, y2, r2), dan dua angka terakhir adalah titik target (tx, ty).</p>
<p>Setelah input diterima, program mengevaluasi posisi titik target terhadap lingkaran pertama dengan memanggil fungsi inside. Nilai koordinat yang tadinya berupa int dikonversi terlebih dahulu menjadi float64 karena perhitungan jarak (dist) membutuhkan angka desimal untuk diakar kuadrat. Fungsi inside akan mengukur jarak titik tersebut ke titik pusat lingkaran; jika jaraknya tidak melebihi jari-jari lingkaran (<= r), maka fungsi mengembalikan nilai true (benar). Hasil pengecekan ini disimpan dalam variabel in1.</p>
<p>Hal yang persis sama dilakukan untuk lingkaran kedua, dan hasil evaluasinya (true/false) disimpan ke dalam variabel in2.</p>
<p>Tahap terakhir adalah menarik kesimpulan. Program menggunakan logika percabangan (if-else). Jika titik target berada di dalam kedua lingkaran (in1 && in2 sama-sama true), maka dicetak "Titik di dalam lingkaran 1 dan 2". Jika hanya in1 yang true, dicetak "Titik di dalam lingkaran 1". Jika hanya in2 yang true, dicetak "Titik di dalam lingkaran 2". Jika semua salah, berarti titik berada jauh di luar kedua lingkaran tersebut.</p>

##### Output
<img width="1920" height="1080" alt="Soal 3" src="https://github.com/user-attachments/assets/4e03c617-747f-4d2f-9387-f5dc5cf56e57" />

