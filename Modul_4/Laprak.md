# <h1 align="center">Laporan Praktikum Modul 4</h1>
<p align="center">[Zhafif Iqbal Kurniawan] - [109082500051]</p>

## Unguided
### 1. [Soal]
#### soal1.go

```go
package main

import (
	"fmt"
)

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}
func mutasi(n, a int, hasil *int) {
	var x, y int
	if n >= a {
		faktorial(n, &x)
		faktorial(n-a, &y)

		*hasil = x / y
	} else {
		*hasil = 0
	}
}

func kombinasi(n, a int, hasil *int) {
	var x, y, z int
	if n >= a {
		faktorial(n, &z)
		faktorial(a, &x)
		faktorial(n-a, &y)

		*hasil = z / (x * y)
	} else {
		*hasil = 0
	}
}

func main() {
	var a, b, c, d, e, f int
	fmt.Scan(&a, &b, &c, &d)
	mutasi(a, c, &e)
	kombinasi(a, c, &f)
	fmt.Println(e, f)

	mutasi(b, d, &e)
	kombinasi(b, d, &f)
	fmt.Println(e, f)
}



```

#### soal2.go

```go
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


```

#### soal3.go

```go
package main

import "fmt"

func baris(n int) {
	fmt.Printf("%d ", n)
	for n != 1 && n > 0 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		fmt.Printf("%d ", n)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	baris(n)
}


```

[penjelasan]
<h2>Soal 1</h2>
<br>
<p>package main : ini adalah paket utama tempat kode dijalankan</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan dan Println</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func faktorial() : fungsi buatan untuk menghitung nilai faktorial. Hasilnya tidak di-return, melainkan langsung disimpan ke memori variabel lewat pointer *hasil</p>
<p>func mutasi() : fungsi buatan untuk menghitung nilai permutasi n! / (n-a)!. Menggunakan pointer untuk menampung dan menyimpan hasil akhir</p>
<p>func kombinasi() : fungsi buatan untuk menghitung nilai kombinasi n! / (a! * (n-a)!). Juga menggunakan pointer untuk hasil akhirnya</p>
<p>func main() : fungsi utama tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>a, b, c, d : variabel untuk menyimpan 4 angka inputan dari pengguna yang mewakili 2 pasang kasus (pasangan n1, r1 dan n2, r2)</p>
<p>e, f : variabel penampung di fungsi utama yang memorinya akan "dipinjamkan" ke fungsi perhitungan agar bisa diisi dengan hasil permutasi dan kombinasi</p>
<p>x, y, z : variabel sementara di dalam fungsi perhitungan untuk menampung nilai faktorial individu sebelum dibagi</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan / println : untuk membaca input dari terminal dan mencetak hasil ke layar dengan baris baru</p>
<p>* (pointer) : tanda asterisk pada parameter fungsi (seperti *hasil dan *int) menandakan bahwa fungsi ini menerima "kunci" akses langsung ke memori suatu variabel untuk memodifikasi nilainya</p>
<p>& (ampersand) : tanda dan yang disertakan saat memanggil fungsi (contoh: &e, &x) digunakan untuk mengirimkan alamat referensi memori dari variabel tersebut</p>
<p>for : perintah perulangan untuk menjalankan perhitungan faktorial</p>
<p>if / else : perintah kondisional untuk memastikan jumlah objek yang tersedia (n) lebih besar atau sama dengan jumlah objek yang diambil (a)</p>
<p>>= : operator perbandingan "lebih dari atau sama dengan"</p>
<p>*= : operator perkalian yang langsung memperbarui nilai variabel (*hasil *= i sama artinya dengan *hasil = *hasil * i)</p>
<p>* / / : operator matematika dasar untuk perkalian dan pembagian</p>
<p>dalam code di atas, program tidak lagi menggunakan "return" untuk memindahkan data antar fungsi. Fungsi utama (main) mula-mula meminta 4 angka yang disimpan di variabel a, b, c, dan d.</p>
<p>Selanjutnya, program memanggil fungsi mutasi(a, c, &e). Artinya: tolong hitung permutasi dari a dan c, lalu simpan hasil akhirnya langsung ke dalam alamat memori variabel e (&e). Di dalam fungsi mutasi, program mengecek apakah a >= c. Jika iya, program kembali memanggil fungsi faktorial menggunakan teknik pointer yang sama (&x dan &y) untuk mendapatkan nilai n! dan (n-a)!. Setelah faktorialnya didapat, fungsi mutasi membaginya dan menyimpan nilainya ke *hasil (yang merujuk pada variabel e).</p>
<p>Proses meminjam alamat memori ini juga dilakukan saat memanggil fungsi kombinasi(a, c, &f). Setelah kedua nilai fungsi didapatkan (tersimpan di e dan f), program utama langsung mencetaknya.</p>
<p>Langkah yang sama persis kemudian diulang untuk pasangan angka kedua (b dan d). Program memanggil lagi fungsi mutasi dan kombinasi, "menimpa" memori variabel e dan f yang lama dengan hasil perhitungan yang baru, lalu kembali mencetaknya ke layar.</p>

##### Output
<img width="1920" height="1080" alt="soal 1" src="https://github.com/user-attachments/assets/ba7c6bb0-12b5-4a98-8b36-4945e20890f3" />



[penjelasan]
<h2>Soal 2</h2>
<br>
<p>package main : ini adalah paket utama tempat kode dijalankan</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan dan Println</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func hitungSkor() : fungsi buatan untuk menghitung total soal yang berhasil dijawab (s) dan total akumulasi waktu (w) menggunakan konsep pointer</p>
<p>func main() : fungsi utama tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>nama, juara : variabel bertipe string (teks) untuk menyimpan nama peserta yang sedang diproses dan nama pemenang sementara</p>
<p>s_max, w_min, s, w, t : variabel angka untuk menyimpan skor tertinggi, rekor waktu tercepat, skor peserta saat ini, waktu total peserta saat ini, dan waktu per soal</p>
<p>awal : variabel boolean (true/false) sebagai penanda/flag apakah data yang sedang dievaluasi adalah data peserta yang pertama dimasukkan</p>
<p>int / string : tipe data untuk bilangan bulat dan teks karakter</p>
<p>scan / println : untuk membaca input teks/angka dari terminal dan mencetak hasil akhir ke layar</p>
<p>* (pointer) / & (ampersand) : digunakan agar fungsi hitungSkor dapat langsung memodifikasi nilai variabel s dan w di dalam memori fungsi utama (main)</p>
<p>for : perintah perulangan. Ada perulangan terbatas (sebanyak 8 kali iterasi di hitungSkor) dan perulangan tanpa batas (for { ... } di fungsi utama)</p>
<p>if : perintah kondisional untuk mengecek dan memvalidasi syarat tertentu</p>
<p>break : perintah untuk menghentikan perulangan secara paksa (aktif ketika nama diisi "Selesai")</p>
<p>== / < / > : operator perbandingan matematika (sama dengan, kurang dari, lebih dari)</p>
<p>|| : operator logika "ATAU" (OR), bernilai benar jika minimal salah satu syarat di kiri atau kanannya terpenuhi</p>
<p>&& : operator logika "DAN" (AND), bernilai benar hanya jika kedua syarat di kiri dan kanannya terpenuhi sekaligus</p>
<p>++ / += : operator increment (tambah 1) dan operator penambahan nilai langsung ke variabel aslinya</p>
<p>dalam code di atas, program dibuat untuk mencari juara perlombaan di mana setiap peserta akan diberikan 8 buah soal untuk diselesaikan.</p>
<p>Program utama berjalan menggunakan perulangan tanpa batas untuk mendata peserta. Pertama-tama, ia meminta input nama peserta. Jika nama yang diketik adalah "Selesai", program langsung berhenti mendata (break). Namun jika bukan "Selesai", program akan memanggil fungsi hitungSkor(&s, &w) untuk menghitung poin peserta tersebut.</p>
<p>Di dalam fungsi hitungSkor, nilai s (skor) dan w (waktu) di-reset menjadi 0 terlebih dahulu. Kemudian, program menjalankan perulangan 8 kali untuk meminta input waktu pengerjaan (t) pada setiap soal. Jika angka waktu tersebut kurang dari 301 (artinya soal berhasil dijawab dalam batas waktu yang sah), maka skor peserta bertambah 1 (*s++) dan waktunya diakumulasikan ke total waktu (*w += t).</p>
<p>Setelah skor selesai dihitung, program kembali ke fungsi utama untuk menyeleksi apakah peserta ini berhak menggantikan posisi juara sementara. Syarat menjadi juara (tercatat di dalam instruksi if yang cukup panjang) adalah: dia adalah peserta paling pertama yang didata (awal == true), ATAU skornya lebih besar dari rekor juara sebelumnya (s > s_max), ATAU skornya seri/sama namun waktu pengerjaannya lebih cepat (s == s_max && w < w_min).</p>
<p>Jika minimal satu dari ketiga kondisi tersebut terpenuhi, maka posisi juara direbut. Data juara, s_max, dan w_min diperbarui dengan data peserta saat ini, dan status "awal" diubah menjadi false. Setelah semua peserta beres didata (user mengetik "Selesai"), perulangan utama berhenti dan program akan mencetak nama sang juara akhir, diikuti dengan total skor dan total waktunya.</p>

##### Output
<img width="1920" height="1080" alt="soal 2" src="https://github.com/user-attachments/assets/e7e5cfc7-d91d-4b08-8fdf-4c2877ed06ae" />




[penjelasan]
<h2>Soal 3</h2>
<br>
<p>package main : ini adalah paket utama tempat kode dijalankan</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan dan Printf</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func baris() : fungsi buatan untuk menghitung dan mencetak urutan barisan bilangan (Collatz sequence) berdasarkan angka awal yang diberikan</p>
<p>func main() : fungsi utama tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>n : variabel untuk menyimpan angka awal yang diinputkan oleh pengguna</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk membaca input angka dari terminal</p>
<p>printf : untuk mencetak angka ke layar dengan format tertentu (%d untuk bilangan bulat, diikuti dengan spasi agar angka tidak menempel)</p>
<p>for : perintah perulangan untuk mengeksekusi blok kode secara berulang selama syaratnya masih terpenuhi</p>
<p>if / else : perintah kondisional untuk mengecek syarat ganjil/genap dan menentukan rumus mana yang akan dipakai</p>
<p>!= : operator "tidak sama dengan" (digunakan untuk memastikan perulangan berjalan selama n tidak sama dengan 1)</p>
<p>> : operator "lebih dari" (untuk memastikan n adalah bilangan bulat positif)</p>
<p>== : operator "sama dengan" (digunakan pada n%2 == 0 untuk mengecek apakah n bernilai genap)</p>
<p>&& : operator logika "DAN" (AND), memastikan kedua syarat pada perulangan (n bukan 1 DAN n lebih dari 0) terpenuhi</p>
<p>% : operator modulus (sisa bagi) untuk mendeteksi bilangan ganjil atau genap</p>
<p>+ / * / / : operator matematika dasar untuk penjumlahan, perkalian, dan pembagian</p>
<p>dalam code di atas, program utama mula-mula meminta pengguna memasukkan satu buah angka bulat positif (n). Setelah angka tersebut dimasukkan, program langsung memanggil fungsi baris(n) dengan membawa angka tersebut sebagai bahan baku.</p>
<p>Di dalam fungsi baris, program pertama-tama mencetak angka asli tersebut ke layar. Kemudian, program masuk ke dalam perulangan bersyarat (while loop ala Go). Perulangan ini akan terus berjalan selama nilai n belum mencapai angka 1 dan n bernilai positif (n != 1 && n > 0).</p>
<p>Di setiap putaran perulangan, program mengecek kondisi angka n saat ini. Jika n adalah bilangan genap (n%2 == 0), maka angka n tersebut dibagi 2 (n = n / 2). Namun, jika n adalah bilangan ganjil (else), maka angka n tersebut dikalikan 3 lalu ditambah 1 (n = 3*n + 1).</p>
<p>Setelah nilai n diperbarui sesuai kondisinya, program mencetak nilai n yang baru tersebut ke layar dengan tambahan spasi di belakangnya. Proses (cek genap/ganjil, ubah nilai, cetak) ini terus menumpuk dan diulang sampai pada akhirnya nilai n anjlok menjadi 1. Saat n mencapai 1, syarat perulangan otomatis tidak terpenuhi lagi, perulangan berhenti, dan fungsi selesai dieksekusi.</p>

##### Output
<img width="1920" height="1080" alt="soal 3" src="https://github.com/user-attachments/assets/a1fae71a-c496-4aa8-9a16-e7413bc23682" />

