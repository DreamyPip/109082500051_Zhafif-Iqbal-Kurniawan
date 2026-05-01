# <h1 align="center">Laporan Praktikum Modul 5</h1>
<p align="center">[Zhafif Iqbal Kurniawan] - [109082500051]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
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

```

#### soal2.go

```go
package main

import (
	"fmt"
	"math"
)

const KAPASITAS int = 100

func main() {
	var n, x, hapusIndex, cariBilangan int
	var data [KAPASITAS]int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	for i := 1; i < n; i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	for i := 0; i < n; i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	fmt.Scan(&x)
	if x > 0 {
		for i := 0; i < n; i++ {
			if i%x == 0 {
				fmt.Print(data[i], " ")
			}
		}
		fmt.Println()
	}

	fmt.Scan(&hapusIndex)
	if hapusIndex >= 0 && hapusIndex < n {
		for i := hapusIndex; i < n-1; i++ {
			data[i] = data[i+1]
		}
		n--
	}
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	var total float64
	for i := 0; i < n; i++ {
		total += float64(data[i])
	}
	rerata := total / float64(n)
	fmt.Printf("%.2f\n", rerata)

	var jumlahKuadratSelisih float64
	for i := 0; i < n; i++ {
		selisih := float64(data[i]) - rerata
		jumlahKuadratSelisih += selisih * selisih
	}
	stdev := math.Sqrt(jumlahKuadratSelisih / float64(n))
	fmt.Printf("%.2f\n", stdev)

	fmt.Scan(&cariBilangan)
	frekuensi := 0
	for i := 0; i < n; i++ {
		if data[i] == cariBilangan {
			frekuensi++
		}
	}
	fmt.Println(frekuensi)
}

```

#### soal3.go

```go
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

```

#### soal4.go

```go
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

```


```

[penjelasan]
<h2>Soal 1</h2>
<br>
<p>package main : ini adalah paket</p>

<p>import ("fmt", "math") : perintah untuk menggunakan paket fmt (input output) dan math (fungsi matematika)</p>

<p>type titik struct : kata kunci untuk membuat tipe data baru bernama titik</p>

<p>x, y float64 : variabel di dalam struct titik untuk menyimpan koordinat x dan y dengan tipe bilangan desimal</p>

<p>type lingkaran struct : kata kunci untuk membuat tipe data baru bernama lingkaran</p>

<p>pusat titik, radius float64 : field di dalam struct lingkaran, pusat bertipe titik dan radius bertipe float64</p>

<p>func jarak(p, q titik) float64 : fungsi buatan untuk menghitung jarak antara dua titik, hasilnya bertipe float64</p>

<p>math.Sqrt() : fungsi untuk menghitung akar kuadrat</p>

<p>math.Pow() : fungsi untuk menghitung pangkat</p>

<p>func didalam(c lingkaran, p titik) bool : fungsi buatan untuk mengecek apakah titik berada di dalam lingkaran, hasilnya true atau false</p>

<p>return jarak(p, c.pusat) <= c.radius : perintah mengembalikan hasil perbandingan jarak titik ke pusat dengan radius</p>

<p>func main() : bisa diartikan sebagai "fungsi utama"</p>

<p>var daftarL [2]lingkaran : deklarasi variabel array berisi 2 data bertipe lingkaran</p>

<p>var t titik : deklarasi variabel t bertipe titik untuk menyimpan koordinat titik yang akan diperiksa</p>

<p>for i := 0; i < 2; i++ : perulangan yang berjalan 2 kali untuk membaca data dua lingkaran</p>

<p>fmt.Scan() : untuk memasukkan data di terminal</p>

<p>cek1 := didalam(daftarL[0], t) : deklarasi dan pengisian variabel cek1 dengan hasil pengecekan titik pada lingkaran pertama</p>

<p>cek2 := didalam(daftarL[1], t) : deklarasi dan pengisian variabel cek2 dengan hasil pengecekan titik pada lingkaran kedua</p>

<p>if, else if, else : perintah percabangan untuk menentukan output berdasarkan kondisi</p>

<p>&& : operator logika "dan"</p>

<p>fmt.Println() : untuk menampilkan hasil ke layar</p>

<p>dalam code di atas terdapat dua tipe data bentukan yaitu titik (dengan field x dan y) dan lingkaran (dengan field pusat bertipe titik dan radius). Program juga memiliki fungsi jarak untuk menghitung jarak antara dua titik menggunakan rumus Euclidean, serta fungsi didalam yang memanfaatkan fungsi jarak untuk menentukan apakah suatu titik berada di dalam lingkaran.</p>

<p>Di fungsi main, program mendeklarasikan array daftarL yang bisa menampung 2 lingkaran dan variabel t untuk titik yang akan diperiksa. Program kemudian melakukan perulangan sebanyak 2 kali untuk membaca input koordinat pusat dan radius dari kedua lingkaran. Setelah itu, program membaca koordinat titik yang akan diperiksa.</p>

<p>Program kemudian memanggil fungsi didalam untuk lingkaran pertama dan kedua, lalu menyimpan hasilnya ke variabel cek1 dan cek2. Program masuk ke logika percabangan: jika cek1 dan cek2 keduanya true, maka titik berada di dalam lingkaran 1 dan 2. Jika hanya cek1 yang true, titik di dalam lingkaran 1 saja. Jika hanya cek2 yang true, titik di dalam lingkaran 2 saja. Jika keduanya false, titik di luar kedua lingkaran. Hasil pengecekan kemudian dicetak ke layar sesuai kondisi yang terpenuhi.</p>

##### Output
<img width="1920" height="1080" alt="Soal 1" src="https://github.com/user-attachments/assets/ac10fa89-d47d-401a-a63e-3112942adcbb" />


[penjelasan]
<h2>Soal 2</h2>
<br>

<p>package main : ini adalah paket</p>

<p>import ("fmt", "math") : perintah untuk menggunakan paket fmt (input output) dan math (fungsi matematika)</p>

<p>const KAPASITAS int = 100 : kata kunci untuk membuat konstanta bernama KAPASITAS dengan nilai 100 dan tipe integer</p>

<p>func main() : bisa diartikan sebagai "fungsi utama"</p>

<p>var n, x, hapusIndex, cariBilangan int : deklarasi variabel untuk menyimpan jumlah data, kelipatan, indeks yang dihapus, dan bilangan yang dicari</p>

<p>var data [KAPASITAS]int : deklarasi array dengan ukuran 100 (dari konstanta KAPASITAS) untuk menyimpan data integer</p>

<p>fmt.Scan(&n) : untuk memasukkan jumlah data di terminal</p>

<p>for i := 0; i < n; i++ : perulangan untuk membaca data sebanyak n kali</p>

<p>fmt.Scan(&data[i]) : untuk memasukkan nilai ke dalam array pada indeks i</p>

<p>fmt.Print(data[i], " ") : untuk menampilkan nilai array diikuti spasi</p>

<p>fmt.Println() : untuk menampilkan baris baru</p>

<p>for i := 1; i < n; i += 2 : perulangan untuk mengakses indeks ganjil (1,3,5,...)</p>

<p>for i := 0; i < n; i += 2 : perulangan untuk mengakses indeks genap (0,2,4,...)</p>

<p>if x > 0 : percabangan jika nilai x lebih besar dari 0</p>

<p>if i%x == 0 : percabangan untuk mengecek apakah indeks i habis dibagi x</p>

<p>if hapusIndex >= 0 && hapusIndex < n : percabangan untuk mengecek apakah indeks yang dihapus valid</p>

<p>data[i] = data[i+1] : perintah menggeser elemen ke kiri untuk menghapus data</p>

<p>n-- : perintah mengurangi jumlah data sebanyak 1</p>

<p>var total float64 : deklarasi variabel total bertipe float64 untuk menampung jumlah data</p>

<p>total += float64(data[i]) : operator penjumlahan yang langsung memperbarui nilai total, data diubah ke float64</p>

<p>rerata := total / float64(n) : operator untuk mendeklarasikan dan mengisi variabel rerata dengan hasil bagi total dan n</p>

<p>fmt.Printf("%.2f\n", rerata) : untuk menampilkan hasil dengan format 2 angka di belakang koma</p>

<p>var jumlahKuadratSelisih float64 : deklarasi variabel untuk menyimpan jumlah kuadrat selisih</p>

<p>selisih := float64(data[i]) - rerata : menghitung selisih antara data dengan nilai rata-rata</p>

<p>stdev := math.Sqrt(jumlahKuadratSelisih / float64(n)) : menghitung standar deviasi dengan akar kuadrat dari varians</p>

<p>frekuensi := 0 : deklarasi dan pengisian variabel frekuensi dengan nilai 0</p>

<p>if data[i] == cariBilangan : percabangan untuk mengecek apakah data sama dengan bilangan yang dicari</p>

<p>frekuensi++ : operator penambahan untuk menghitung kemunculan bilangan</p>

<p>dalam code di atas terdapat konstanta KAPASITAS yang bernilai 100 untuk batas maksimal array. Program mendeklarasikan variabel n (jumlah data), x (kelipatan indeks), hapusIndex (indeks yang akan dihapus), dan cariBilangan (angka yang dicari frekuensinya), serta array data berukuran 100.</p>

<p>Program pertama-tama membaca nilai n, lalu membaca n buah bilangan ke dalam array menggunakan perulangan. Setelah semua data masuk, program menampilkan seluruh isi array, kemudian menampilkan elemen pada indeks ganjil (1,3,5,...), lalu menampilkan elemen pada indeks genap (0,2,4,...).</p>

<p>Selanjutnya program membaca nilai x. Jika x lebih dari 0, program akan menampilkan elemen-elemen pada indeks kelipatan x (indeks 0, x, 2x, ...).</p>

<p>Program kemudian membaca nilai hapusIndex. Jika hapusIndex valid (antara 0 sampai n-1), program akan menghapus elemen pada indeks tersebut dengan menggeser elemen setelahnya ke kiri, lalu mengurangi nilai n sebanyak 1. Setelah penghapusan, program menampilkan isi array yang baru.</p>

<p>Program menghitung rata-rata dari data yang tersisa dengan menjumlahkan semua data (diubah ke float64) lalu dibagi dengan n. Hasil rata-rata ditampilkan dengan 2 angka desimal.</p>

<p>Program kemudian menghitung standar deviasi dengan rumus: akar kuadrat dari jumlah kuadrat selisih antara setiap data dengan rata-rata, dibagi dengan n. Hasil standar deviasi juga ditampilkan dengan 2 angka desimal menggunakan fungsi math.Sqrt.</p>

<p>Terakhir, program membaca nilai cariBilangan, lalu menghitung berapa kali bilangan tersebut muncul dalam array (frekuensi) dengan melakukan perulangan dan pencocokan nilai. Hasil frekuensi ditampilkan di layar.</p>

##### Output
<img width="1920" height="1080" alt="Soal 2" src="https://github.com/user-attachments/assets/0f36fd4b-7849-4cd2-bfa6-f17def109ac2" />


[penjelasan]
<h2>Soal 3</h2>
<br>
<p>package main : ini adalah paket</p>

<p>import "fmt" : perintah untuk menggunakan paket fmt (input output)</p>

<p>func main() : bisa diartikan sebagai "fungsi utama"</p>

<p>var timA, timB string : deklarasi variabel untuk menyimpan nama klub A dan B bertipe teks</p>

<p>var poinA, poinB int : deklarasi variabel untuk menyimpan poin klub A dan B bertipe bilangan bulat</p>

<p>fmt.Print("Klub A : ") : untuk menampilkan teks ke layar tanpa pindah baris</p>

<p>fmt.Scan(&timA) : untuk memasukkan nama klub A di terminal</p>

<p>daftarPemenang := []string{} : operator untuk mendeklarasikan dan mengisi slice (array dinamis) kosong yang akan menyimpan hasil setiap pertandingan</p>

<p>for i := 1; ; i++ : perulangan tanpa kondisi henti (infinite loop), i dimulai dari 1</p>

<p>fmt.Printf("Pertandingan %d : ", i) : untuk menampilkan teks dengan format, %d akan diganti dengan nomor pertandingan</p>

<p>fmt.Scan(&poinA, &poinB) : untuk memasukkan dua nilai poin sekaligus di terminal</p>

<p>if poinA < 0 || poinB < 0 : percabangan untuk mengecek apakah salah satu poin bernilai negatif</p>

<p>|| : operator logika "atau"</p>

<p>break : perintah untuk menghentikan perulangan</p>

<p>if poinA > poinB : percabangan jika poin A lebih besar dari poin B</p>

<p>daftarPemenang = append(daftarPemenang, timA) : perintah untuk menambahkan nilai (nama tim A) ke dalam slice daftarPemenang</p>

<p>else if poinA < poinB : percabangan jika poin A lebih kecil dari poin B</p>

<p>else : percabangan jika poin A sama dengan poin B</p>

<p>daftarPemenang = append(daftarPemenang, "Draw") : menambahkan kata "Draw" ke dalam slice</p>

<p>for j := 0; j < len(daftarPemenang); j++ : perulangan untuk mengakses semua isi slice, len() digunakan untuk mengetahui panjang slice</p>

<p>fmt.Printf("Hasil %d : %s\n", j+1, daftarPemenang[j]) : menampilkan hasil dengan format, %d untuk nomor, %s untuk nama pemenang</p>

<p>fmt.Println("Pertandingan selesai") : untuk menampilkan teks diakhiri pindah baris</p>

<p>dalam code di atas terdapat program yang mencatat hasil pertandingan antara dua klub. Program mendeklarasikan variabel timA, timB untuk nama klub, serta poinA dan poinB untuk skor setiap pertandingan.</p>

<p>Program pertama-tama meminta pengguna memasukkan nama Klub A dan Klub B. Kemudian membuat slice kosong bernama daftarPemenang untuk menyimpan hasil setiap pertandingan.</p>

<p>Program masuk ke perulangan tanpa henti dimulai dari i=1. Di setiap perulangan, program menampilkan "Pertandingan ke-i" lalu membaca poinA dan poinB. Jika poinA atau poinB bernilai negatif, program akan keluar dari perulangan menggunakan break.</p>

<p>Setelah membaca poin, program membandingkan nilai poinA dan poinB. Jika poinA lebih besar, maka nama timA ditambahkan ke daftarPemenang. Jika poinB lebih besar, maka nama timB yang ditambahkan. Jika kedua poin sama, maka kata "Draw" yang ditambahkan.</p>

<p>Setelah perulangan berhenti (karena ada input negatif), program menampilkan semua hasil pertandingan menggunakan perulangan for. Setiap hasil ditampilkan dengan format "Hasil ke-j : pemenang", lalu di akhir program menampilkan teks "Pertandingan selesai".</p>

##### Output
<img width="1920" height="1080" alt="Soal 3" src="https://github.com/user-attachments/assets/87bdc9bf-546b-4d13-add4-d49d42b8f19f" />

[penjelasan]
<h2>Soal 4</h2>
<br>
<p>package main : ini adalah paket</p>

<p>import "fmt" : perintah untuk menggunakan paket fmt (input output)</p>

<p>const MAX int = 127 : kata kunci untuk membuat konstanta bernama MAX dengan nilai 127 dan tipe integer</p>

<p>type daftarKarakter [MAX]rune : kata kunci untuk membuat tipe data baru bernama daftarKarakter berupa array dengan ukuran MAX (127) yang bertipe rune (karakter Unicode)</p>

<p>func isi(t *daftarKarakter, jumlah *int) : fungsi buatan untuk mengisi array karakter, parameter t adalah pointer ke array dan jumlah adalah pointer ke integer</p>

<p>*daftarKarakter : tanda bintang menunjukkan pointer (alamat memori) ke tipe data daftarKarakter</p>

<p>var huruf rune : deklarasi variabel huruf bertipe rune untuk menyimpan karakter sementara</p>

<p>*jumlah = 0 : perintah mengisi nilai yang ditunjuk oleh pointer jumlah dengan 0 (menggunakan tanda bintang)</p>

<p>for *jumlah < MAX : perulangan berjalan selama nilai yang ditunjuk pointer jumlah masih kurang dari MAX</p>

<p>fmt.Scanf("%c", &huruf) : untuk membaca satu karakter dari terminal, %c adalah format untuk karakter</p>

<p>if huruf == '.' : percabangan jika karakter yang dimasukkan adalah tanda titik</p>

<p>break : perintah untuk menghentikan perulangan</p>

<p>if huruf != ' ' && huruf != '\n' && huruf != '\r' : percabangan untuk mengabaikan spasi, enter, dan carriage return</p>

<p>t[*jumlah] = huruf : mengisi array t pada indeks sesuai nilai yang ditunjuk pointer jumlah dengan karakter huruf</p>

<p>*jumlah++ : operator penambahan untuk menambah nilai yang ditunjuk pointer jumlah sebanyak 1</p>

<p>func tampil(t daftarKarakter, jumlah int) : fungsi buatan untuk menampilkan isi array, t adalah array (bukan pointer) dan jumlah adalah integer</p>

<p>for i := 0; i < jumlah; i++ : perulangan untuk mengakses indeks dari 0 sampai jumlah-1</p>

<p>fmt.Printf("%c ", t[i]) : untuk menampilkan karakter diikuti spasi, %c untuk format karakter</p>

<p>fmt.Println() : untuk menampilkan baris baru</p>

<p>func balik(t *daftarKarakter, jumlah int) : fungsi buatan untuk membalik urutan array, t adalah pointer ke array</p>

<p>for i := 0; i < jumlah/2; i++ : perulangan sebanyak setengah dari jumlah (hanya sampai tengah)</p>

<p>temp := t[i] : menyimpan nilai sementara dari indeks i ke variabel temp</p>

<p>t[i] = t[jumlah-1-i] : mengisi indeks i dengan nilai dari indeks sebelah kanan (simetris)</p>

<p>t[jumlah-1-i] = temp : mengisi indeks kanan dengan nilai temp yang tadi disimpan</p>

<p>func cekPalindrom(t daftarKarakter, jumlah int) bool : fungsi buatan untuk mengecek apakah array adalah palindrom, mengembalikan nilai true atau false</p>

<p>if t[i] != t[jumlah-1-i] : percabangan jika karakter di indeks i tidak sama dengan karakter di indeks kanan simetrisnya</p>

<p>return false : perintah mengembalikan nilai false (bukan palindrom)</p>

<p>return true : perintah mengembalikan nilai true (palindrom)</p>

<p>func main() : bisa diartikan sebagai "fungsi utama"</p>

<p>var data daftarKarakter : deklarasi variabel data dengan tipe daftarKarakter (array berukuran 127)</p>

<p>var n int : deklarasi variabel n untuk menyimpan jumlah karakter yang dimasukkan</p>

<p>fmt.Print("Teks : ") : untuk menampilkan teks ke layar</p>

<p>isi(&data, &n) : memanggil fungsi isi dengan mengirimkan alamat memori dari data dan n (menggunakan tanda &)</p>

<p>& : operator untuk mengambil alamat memori dari suatu variabel</p>

<p>isPalindrom := cekPalindrom(data, n) : mendeklarasikan dan mengisi variabel isPalindrom dengan hasil pengecekan palindrom</p>

<p>fmt.Print("Reverse teks : ") : menampilkan teks "Reverse teks : " ke layar</p>

<p>balik(&data, n) : memanggil fungsi balik dengan mengirimkan alamat data (pointer) dan jumlah n</p>

<p>tampil(data, n) : memanggil fungsi tampil untuk menampilkan array yang sudah dibalik</p>

<p>fmt.Printf("Palindrom : %v\n", isPalindrom) : menampilkan hasil palindrom, %v adalah format untuk nilai default (boolean akan tampil true/false)</p>

<p>dalam code di atas terdapat konstanta MAX bernilai 127 untuk batas maksimal array. Program membuat tipe data bentukan daftarKarakter berupa array rune berukuran 127. Terdapat empat fungsi buatan: isi (mengisi array), tampil (menampilkan array), balik (membalik urutan array), dan cekPalindrom (memeriksa apakah array simetris).</p>

<p>Program menggunakan pointer untuk fungsi isi dan balik. Pointer memungkinkan fungsi mengubah nilai asli variabel karena yang dikirim adalah alamat memori, bukan salinan nilainya. Pada fungsi isi, parameter t *daftarKarakter dan jumlah *int memungkinkan fungsi mengisi array dan mengubah nilai n secara langsung.</p>

<p>Di fungsi main, program menampilkan "Teks : " lalu memanggil fungsi isi(&data, &n). Fungsi isi akan membaca karakter satu per satu menggunakan fmt.Scanf dengan format %c. Karakter yang valid (bukan spasi, enter, carriage return) disimpan ke array data. Jika karakter yang dimasukkan adalah tanda titik (.), maka perulangan berhenti.</p>

<p>Setelah array terisi, program mengecek apakah teks tersebut palindrom dengan memanggil fungsi cekPalindrom(data, n). Fungsi cekPalindrom membandingkan karakter dari awal dan akhir secara bersamaan. Jika ada yang tidak sama, fungsi langsung mengembalikan false. Semua karakter cocok maka mengembalikan true.</p>

<p>Program kemudian membalik urutan array dengan memanggil fungsi balik(&data, n). Fungsi balik menggunakan perulangan dari indeks 0 sampai tengah array, menukar karakter indeks i dengan karakter indeks jumlah-1-i menggunakan variabel bantuan temp.</p>

<p>Setelah dibalik, program memanggil fungsi tampil(data, n) untuk menampilkan isi array yang sudah dibalik. Terakhir, program menampilkan hasil pengecekan palindrom (true jika palindrom, false jika tidak) dengan format %v.</p>

##### Output
<img width="1920" height="1080" alt="Soal 4" src="https://github.com/user-attachments/assets/affa918c-cffb-4ef9-97d4-fffc6e78aa5c" />

