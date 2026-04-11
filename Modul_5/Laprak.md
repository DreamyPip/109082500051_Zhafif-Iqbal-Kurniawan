# <h1 align="center">Laporan Praktikum Modul 5</h1>
<p align="center">[Zhafif Iqbal Kurniawan] - [109082500051]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", fibonaci(i))
	}
}

func fibonaci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonaci(n-1) + fibonaci(n-2)
}

```

#### soal2.go

```go
package main

import "fmt"

func star(n int) {
	if n != 0 {
		fmt.Print("*")
		star(n - 1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		star(i)
		fmt.Println()
	}
}

```

#### soal3.go

```go
package main

import "fmt"

func bagi(n int, i int) {
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	if n >= i {
		bagi(n, i+1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	bagi(n, 1)
}

```

#### soal4.go

```go
package main

import "fmt"

func baris(n int) {
	if n == 1 {
		fmt.Print(n, " ")
		return
	}

	fmt.Print(n, " ")
	baris(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Scan(&n)
	baris(n)
	fmt.Println()
}

```

#### soal5.go

```go
package main

import "fmt"

func ganjil(n int, x int) {
	if x <= n {
		fmt.Print(x, " ")
		ganjil(n, x+2)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	ganjil(n, 1)
}

```

#### soal6.go

```go
package main

import "fmt"

func pangkat(n, x int) int {
	if x == 0 {
		return 1
	}
	return n * pangkat(n, x-1)
}

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	fmt.Print(pangkat(n, x))
}

```

[penjelasan]
<h2>Soal 1</h2>
<br>
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan dan Printf</p>
<p>func : kata kunci untuk membuat fungsi, baik fungsi utama maupun fungsi rekursif buatan</p>
<p>func main() : bisa diartikan sebagai “fungsi utama”</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>n : variabel yang dibuat untuk menyimpan jumlah deret bilangan yang diinginkan</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk memasukan data di terminal</p>
<p>printf : untuk menampilkan hasil ke samping dengan format angka bulat</p>
<p>for : perintah untuk melakukan perulangan (looping)</p>
<p>if : struktur percabangan logika</p>
<p>return : perintah untuk mengembalikan nilai dari fungsi</p>
<p><= : operator “lebih kecil sama dengan”</p>
<p>- : operator pengurangan untuk memecah masalah menjadi lebih kecil</p>
<p>+ : operator penjumlahan untuk menggabungkan hasil rekursif</p>

<p>dalam code di atas terdapat variabel n yang diinputkan oleh user untuk menentukan batas deret bilangan Fibonacci yang akan ditampilkan. Program kemudian masuk ke perulangan for dari angka 0 sampai n. Pada setiap putaran, program memanggil fungsi bernama fibonaci(i). Program kemudian masuk ke logika rekursif di dalam fungsi tersebut:</p>

<p>Jika nilai n kurang dari atau sama dengan 1, maka program akan langsung mengembalikan nilai n itu sendiri (sebagai angka dasar 0 atau 1).</p>

<p>Jika nilai n lebih besar dari 1, maka program akan memanggil dirinya sendiri berulang kali dengan rumus penjumlahan dua angka sebelumnya, yaitu fibonaci(n-1) ditambah fibonaci(n-2).</p>

<p>Hasil dari setiap pemanggilan fungsi tersebut akan dicetak ke samping, sehingga membentuk urutan angka di mana setiap angka adalah hasil penjumlahan dari dua angka di belakangnya.</p>

##### Output
<img width="1920" height="1080" alt="Soal 1" src="https://github.com/user-attachments/assets/ac10fa89-d47d-401a-a63e-3112942adcbb" />


[penjelasan]
<h2>Soal 2</h2>
<br>

<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan, Print, dan Println</p>
<p>func : kata kunci untuk membuat fungsi, baik fungsi utama maupun fungsi rekursif buatan</p>
<p>func main() : bisa diartikan sebagai “fungsi utama”</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>n : variabel yang dibuat untuk menyimpan angka input dari user untuk menentukan tinggi pola bintang</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk memasukan data di terminal</p>
<p>print : untuk menampilkan karakter bintang ke samping tanpa membuat baris baru</p>
<p>println : untuk membuat baris baru setelah satu baris bintang selesai dicetak</p>
<p>for : perintah untuk melakukan perulangan (looping)</p>
<p>if : struktur percabangan logika</p>
<p>!= : operator "tidak sama dengan"</p>
<p>- : operator pengurangan untuk mengurangi jumlah bintang yang perlu dicetak dalam satu baris</p>
<p><= : operator “lebih kecil sama dengan”</p>

<p>dalam code di atas terdapat variabel n yang diinputkan oleh user untuk menentukan berapa banyak baris pola bintang yang ingin dibentuk. Program kemudian menjalankan perulangan for dari i = 1 sampai n. Pada setiap barisnya, program memanggil fungsi rekursif bernama star(i). Program kemudian masuk ke logika rekursif di dalam fungsi tersebut:</p>

<p>Jika nilai n tidak sama dengan 0, maka program akan mencetak satu karakter bintang "*" dan kemudian memanggil dirinya sendiri kembali dengan nilai n yang sudah dikurangi 1 (star(n - 1)).</p>

<p>Proses pemanggilan diri sendiri ini akan terus berulang sampai nilai n menjadi 0, sehingga jumlah bintang yang tercetak di setiap baris akan sesuai dengan urutan barisnya (baris 1 mencetak 1 bintang, baris 2 mencetak 2 bintang, dan seterusnya).</p>

<p>Setelah fungsi rekursif star selesai mencetak bintang dalam satu baris, fungsi utama akan menjalankan fmt.Println() untuk berpindah ke baris di bawahnya sebelum melanjutkan perulangan berikutnya.</p>

##### Output
<img width="1920" height="1080" alt="Soal 2" src="https://github.com/user-attachments/assets/0f36fd4b-7849-4cd2-bfa6-f17def109ac2" />


[penjelasan]
<h2>Soal 3</h2>
<br>
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan dan Print</p>
<p>func : kata kunci untuk membuat fungsi, baik fungsi utama maupun fungsi rekursif buatan</p>
<p>func main() : bisa diartikan sebagai “fungsi utama”</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>n : variabel yang dibuat untuk menyimpan angka bulat positif yang akan dicari faktor pembaginya</p>
<p>i : variabel pembantu yang berfungsi sebagai angka pembagi, dimulai dari 1</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk memasukan data di terminal</p>
<p>print : untuk menampilkan hasil angka pembagi ke samping dengan spasi</p>
<p>if : struktur percabangan logika</p>
<p>== : operator "sama dengan" untuk mengecek sisa bagi</p>
<p>>= : operator “lebih besar sama dengan” untuk memastikan rekursi berlanjut sampai angka pembagi mencapai nilai n</p>
<p>% : modulus atau sisa bagi</p>
<p>+ : operator penjumlahan untuk menaikkan nilai angka pembagi pada setiap tahap rekursi</p>

<p>dalam code di atas terdapat variabel n yang diinputkan oleh user untuk dicari faktor-faktor pembaginya. Program kemudian memanggil fungsi rekursif bernama bagi(n, 1). Program kemudian masuk ke logika rekursif di dalam fungsi tersebut:</p>

<p>Jika n habis dibagi i (n % i == 0), maka program akan mencetak nilai i karena i merupakan salah satu faktor pembagi dari n.</p>

<p>Selanjutnya, program mengecek apakah nilai n masih lebih besar atau sama dengan i. Jika ya, maka fungsi akan memanggil dirinya sendiri kembali dengan menaikkan nilai pembagi sebanyak 1 (bagi(n, i+1)).</p>

<p>Proses ini akan terus berulang secara rekursif mulai dari angka 1 sampai angka pembagi tersebut sama dengan n, sehingga semua faktor pembagi dari bilangan tersebut akan tercetak secara berurutan.</p>

##### Output
<img width="1920" height="1080" alt="Soal 3" src="https://github.com/user-attachments/assets/87bdc9bf-546b-4d13-add4-d49d42b8f19f" />

[penjelasan]
<h2>Soal 4</h2>
<br>
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan, Print, dan Println</p>
<p>func : kata kunci untuk membuat fungsi, baik fungsi utama maupun fungsi rekursif buatan</p>
<p>func main() : bisa diartikan sebagai “fungsi utama”</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>n : variabel yang dibuat untuk menyimpan angka bulat positif sebagai titik awal barisan</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk memasukan data di terminal</p>
<p>print : untuk menampilkan angka ke samping dengan spasi</p>
<p>println : untuk membuat baris baru setelah seluruh barisan selesai dicetak</p>
<p>if : struktur percabangan logika</p>
<p>return : perintah untuk mengakhiri eksekusi fungsi dan kembali ke pemanggilnya</p>
<p>== : operator "sama dengan"</p>
<p>- : operator pengurangan untuk menurunkan nilai angka pada pemanggilan rekursif berikutnya</p>

<p>dalam code di atas terdapat variabel n yang diinputkan oleh user untuk membentuk sebuah barisan bilangan simetris. Program kemudian memanggil fungsi rekursif bernama barisan(n). Program kemudian masuk ke logika rekursif di dalam fungsi tersebut:</p>

<p>Jika nilai n sama dengan 1, program akan mencetak angka 1 lalu berhenti (return) karena sudah mencapai titik tengah barisan.</p>

<p>Jika nilai n masih lebih besar dari 1, program akan mencetak nilai n saat itu, kemudian memanggil dirinya sendiri dengan nilai yang dikurangi 1 (barisan(n - 1)).</p>

<p>Setelah pemanggilan rekursif tersebut selesai (setelah mencapai angka 1), program akan melanjutkan perintah di bawahnya, yaitu mencetak kembali nilai n saat itu. Hal inilah yang menyebabkan angka-angka tersebut muncul kembali secara terbalik sehingga menghasilkan barisan yang menurun lalu menaik lagi secara simetris.</p>

##### Output
<img width="1920" height="1080" alt="Soal 4" src="https://github.com/user-attachments/assets/affa918c-cffb-4ef9-97d4-fffc6e78aa5c" />


[penjelasan]
<h2>Soal 5</h2>
<br>
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan dan Print</p>
<p>func : kata kunci untuk membuat fungsi, baik fungsi utama maupun fungsi rekursif buatan</p>
<p>func main() : bisa diartikan sebagai “fungsi utama”</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>n : variabel yang dibuat untuk menyimpan batas angka maksimal yang diinputkan user</p>
<p>x : variabel pembantu yang merepresentasikan angka ganjil yang sedang diproses, dimulai dari 1</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk memasukan data di terminal</p>
<p>print : untuk menampilkan angka ganjil ke samping dengan spasi</p>
<p>if : struktur percabangan logika</p>
<p><= : operator “lebih kecil sama dengan”</p>
<p>+ : operator penjumlahan untuk meloncat ke angka ganjil berikutnya (x+2)</p>

<p>dalam code di atas terdapat variabel n yang diinputkan oleh user untuk menampilkan deret bilangan ganjil hingga batas tersebut. Program kemudian memanggil fungsi rekursif bernama ganjil(n, 1). Program kemudian masuk ke logika rekursif di dalam fungsi tersebut:</p>

<p>Jika nilai x masih lebih kecil atau sama dengan n, maka program akan mencetak nilai x tersebut ke layar.</p>

<p>Setelah mencetak, fungsi akan memanggil dirinya sendiri kembali dengan mengirimkan nilai x yang sudah ditambah 2 (ganjil(n, x+2)). Penambahan 2 ini memastikan bahwa angka yang diproses selanjutnya tetap ganjil (1, 3, 5, dst).</p>

<p>Proses ini akan terus berulang secara otomatis sampai nilai x melampaui nilai n, sehingga deret bilangan ganjil akan tercetak berurutan sesuai batas yang diinginkan.</p>

##### Output
<img width="1920" height="1080" alt="Soal 5" src="https://github.com/user-attachments/assets/52df1f8d-c387-4bf1-b495-7f8e732f4e13" />

[penjelasan]
<h2>Soal 6</h2>
<br>
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan dan Print</p>
<p>func : kata kunci untuk membuat fungsi, baik fungsi utama maupun fungsi rekursif buatan</p>
<p>func main() : bisa diartikan sebagai “fungsi utama”</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>n, x : variabel yang dibuat untuk menyimpan angka basis (n) dan angka pangkat (x)</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk memasukan data di terminal</p>
<p>print : untuk menampilkan hasil perhitungan ke layar</p>
<p>if : struktur percabangan logika</p>
<p>return : perintah untuk mengembalikan nilai dari sebuah fungsi</p>
<p>== : operator "sama dengan"</p>
<p>* : operator perkalian matematika</p>
<p>- : operator pengurangan</p>

<p>dalam code di atas terdapat variabel n dan x yang diinputkan oleh user melalui fungsi utama. Program kemudian memanggil fungsi buatan bernama pangkat(n, x) untuk menghitung hasil pemangkatan secara rekursif. Program kemudian masuk ke logika rekursif di dalam fungsi tersebut:</p>

<p>Jika nilai pangkat (x) sama dengan 0, maka program akan mengembalikan nilai 1 (karena angka apapun pangkat 0 hasilnya adalah 1).</p>

<p>Jika x belum mencapai 0, program akan mengalikan nilai n dengan hasil dari pemanggilan kembali fungsi pangkat dengan nilai pangkat yang dikurangi 1 (n * pangkat(n, x-1)).</p>

<p>Proses perkalian ini akan terus berulang dan menumpuk sampai nilai x mencapai 0, sehingga akhirnya menghasilkan total nilai n yang dipangkatkan sebanyak x.</p>

##### Output
<img width="1920" height="1080" alt="Soal 6" src="https://github.com/user-attachments/assets/393c0793-6f57-4e79-8ea1-f34a89aaf0b0" />

