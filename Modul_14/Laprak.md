# <h1 align="center">Laporan Praktikum Modul 14</h1>
<p align="center">[Zhafif Iqbal Kurniawan] - [109082500051]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

const NMAX = 1000000

var arr [NMAX]int

func selectionSort(n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		temp := arr[i]
		arr[i] = arr[idxMin]
		arr[idxMin] = temp
	}
}

func main() {
	var n, m int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		selectionSort(m)

		for j := 0; j < m; j++ {
			fmt.Print(arr[j])
			if j < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

```

#### soal2.go

```go
package main

import "fmt"

const NMAX = 1000000

var arrGanjil [NMAX]int
var arrGenap [NMAX]int

func sortGanjil(n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arrGanjil[j] < arrGanjil[idxMin] {
				idxMin = j
			}
		}
		temp := arrGanjil[i]
		arrGanjil[i] = arrGanjil[idxMin]
		arrGanjil[idxMin] = temp
	}
}

func sortGenap(n int) {
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if arrGenap[j] > arrGenap[idxMax] {
				idxMax = j
			}
		}
		temp := arrGenap[i]
		arrGenap[i] = arrGenap[idxMax]
		arrGenap[idxMax] = temp
	}
}

func main() {
	var n, m, num int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		var nGanjil int = 0
		var nGenap int = 0

		for j := 0; j < m; j++ {
			fmt.Scan(&num)
			if num%2 != 0 {
				arrGanjil[nGanjil] = num
				nGanjil++
			} else {
				arrGenap[nGenap] = num
				nGenap++
			}
		}

		sortGanjil(nGanjil)
		sortGenap(nGenap)

		var pertama int = 1

		for j := 0; j < nGanjil; j++ {
			if pertama == 1 {
				fmt.Print(arrGanjil[j])
				pertama = 0
			} else {
				fmt.Print(" ", arrGanjil[j])
			}
		}

		for j := 0; j < nGenap; j++ {
			if pertama == 1 {
				fmt.Print(arrGenap[j])
				pertama = 0
			} else {
				fmt.Print(" ", arrGenap[j])
			}
		}
		fmt.Println()
	}
}

```

#### soal3.go

```go
package main

import "fmt"

const NMAX = 1000000

var arr [NMAX]int

func insertionSort(n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var n int = 0
	var num int

	fmt.Scan(&num)
	for num != -5313 {
		if num == 0 {
			if n > 0 {
				insertionSort(n)
				if n%2 != 0 {
					fmt.Println(arr[n/2])
				} else {
					fmt.Println((arr[n/2-1] + arr[n/2]) / 2)
				}
			}
		} else {
			arr[n] = num
			n++
		}
		fmt.Scan(&num)
	}
}

```

#### soal4.go

```go
package main

import "fmt"

const NMAX = 1000

type tabInt [NMAX]int

func insertionSort(arr *tabInt, n int) {
	for i := 1; i < n; i++ {
		key := (*arr)[i]
		j := i - 1
		for j >= 0 && (*arr)[j] > key {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = key
	}
}

func cekJarak(arr tabInt, n int) {
	if n > 1 {
		diff := arr[1] - arr[0]
		status := 1

		for i := 1; i < n-1; i++ {
			if arr[i+1]-arr[i] != diff {
				status = 0
			}
		}

		if status == 1 {
			fmt.Printf("Data berjarak %d\n", diff)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}

func main() {
	var arr tabInt
	var n int
	var num int

	fmt.Scan(&num)
	for num >= 0 {
		arr[n] = num
		n++
		fmt.Scan(&num)
	}

	insertionSort(&arr, n)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	cekJarak(arr, n)
}


```

#### soal5.go

```go
package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax + 1]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i <= n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n > 0 {
		maxRating := pustaka[1].rating
		for i := 2; i <= n; i++ {
			if pustaka[i].rating > maxRating {
				maxRating = pustaka[i].rating
			}
		}

		for i := 1; i <= n; i++ {
			if pustaka[i].rating == maxRating {
				fmt.Printf("%s, %s, %s, %d\n", pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].tahun)
			}
		}
	}
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 2; i <= n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 1 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 1; i <= limit; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 1
	kanan := n
	found := -1

	for kiri <= kanan && found == -1 {
		med := (kiri + kanan) / 2
		if pustaka[med].rating == r {
			found = med
		} else if pustaka[med].rating < r {
			kanan = med - 1
		} else {
			kiri = med + 1
		}
	}

	if found != -1 {
		fmt.Printf("%s, %s, %s, %d, %d, %d\n", pustaka[found].judul, pustaka[found].penulis, pustaka[found].penerbit, pustaka[found].tahun, pustaka[found].eksemplar, pustaka[found].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var n, targetRating int
	var pustaka DaftarBuku

	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)
	fmt.Scan(&targetRating)

	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	CariBuku(pustaka, n, targetRating)
}


```


[penjelasan]
<h2>Soal 1</h2>
<br>
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan, Print, dan Println</p>
<p>const : kata kunci untuk mendeklarasikan nilai tetap (konstanta)</p>
<p>NMAX = 1000000 : konstanta yang menentukan batas maksimal ukuran array, yaitu 1 juta elemen</p>
<p>var arr [NMAX]int : deklarasi variabel array global bernama arr yang bisa diakses dan diubah secara langsung oleh semua fungsi di dalam program</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func selectionSort() : fungsi buatan untuk mengurutkan angka dari yang terkecil ke terbesar menggunakan metode Selection Sort</p>
<p>func main() : bisa diartikan sebagai “fungsi utama” tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>n, m : variabel untuk menyimpan jumlah kelompok data atau baris (n) dan jumlah elemen di masing-masing kelompok (m)</p>
<p>idxMin, temp : variabel bantuan di dalam fungsi untuk menyimpan letak (indeks) angka terkecil dan nilai sementara saat proses penukaran posisi angka</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk memasukan data di terminal</p>
<p>print : untuk menampilkan hasil ke layar secara menyamping (tanpa membuat baris baru)</p>
<p>println : untuk menampilkan hasil ke layar dan secara otomatis membuat baris baru</p>
<p>for : perintah perulangan untuk mengeksekusi blok kode secara berulang</p>
<p>:= : operator untuk mendeklarasikan sekaligus mengisi nilai variabel baru secara singkat</p>
<p>++ : operator increment untuk menambah nilai variabel sebanyak 1</p>
<p>if : perintah kondisional untuk mengecek sebuah syarat</p>
<p>< : operator "kurang dari"</p>
<p>+ / - : operator matematika (penjumlahan dan pengurangan)</p>
<p>dalam code di atas, program utama mula-mula meminta input n untuk menentukan berapa banyak baris/kelompok data (test case) yang akan diproses. Program kemudian menjalankan perulangan luar sebanyak n kali.</p>
<p>Di setiap perulangannya, program meminta input m yang merupakan jumlah angka untuk baris tersebut. Kemudian, program menjalankan perulangan untuk menerima input m buah angka dan menyimpannya secara berurutan ke dalam array global bernama arr.</p>
<p>Setelah sekumpulan angka dalam satu baris selesai dimasukkan, program memanggil fungsi selectionSort(m) untuk mengurutkannya. Cara kerja algoritma ini: program membagi data menjadi bagian yang sudah urut dan yang belum. Melalui dua perulangan (perulangan i dan j), program mencari nilai paling kecil di bagian yang belum urut dengan mencatat posisinya (idxMin). Setelah angka terkecil ditemukan, program menukar angka tersebut dengan angka yang ada di posisi i menggunakan bantuan variabel temp. Proses pencarian dan penukaran ini diulang terus sampai semua elemen urut membesar.</p>
<p>Setelah proses pengurutan selesai, program utama (main) menjalankan perulangan terakhir untuk mencetak isi array arr yang sekarang sudah tersusun rapi dari kecil ke besar. Setiap angka dipisahkan oleh spasi, dan diakhiri dengan baris baru (Println) untuk persiapan memproses kelompok data selanjutnya.</p>

##### Output
<img width="1920" height="1080" alt="Soal 1" src="https://github.com/user-attachments/assets/4c15fc33-19e6-42cf-b77a-1fcef3e190ea" />


[penjelasan]
<h2>Soal 2</h2>
<br>

<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan, Print, dan Println</p>
<p>const : kata kunci untuk mendeklarasikan nilai tetap (konstanta)</p>
<p>NMAX = 1000000 : konstanta batas maksimal ukuran array</p>
<p>var arrGanjil, arrGenap [NMAX]int : deklarasi dua variabel array global untuk menampung kumpulan angka ganjil dan angka genap secara terpisah</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func sortGanjil() : fungsi untuk mengurutkan isi array ganjil dari yang terkecil ke terbesar (ascending) menggunakan Selection Sort</p>
<p>func sortGenap() : fungsi untuk mengurutkan isi array genap dari yang terbesar ke terkecil (descending) menggunakan Selection Sort</p>
<p>func main() : bisa diartikan sebagai “fungsi utama” tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>n, m, num : variabel untuk menyimpan jumlah baris/test case (n), jumlah angka per baris (m), dan angka inputan user saat ini (num)</p>
<p>nGanjil, nGenap : variabel untuk menghitung ada berapa banyak angka ganjil dan genap yang berhasil dikumpulkan</p>
<p>pertama : variabel penanda (flag) yang digunakan untuk mengatur posisi spasi saat mencetak hasil agar rapi</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk memasukan data di terminal</p>
<p>print / println : untuk menampilkan hasil menyamping dan membuat baris baru</p>
<p>for : perintah perulangan untuk mengeksekusi blok kode secara berulang</p>
<p>if / else : perintah kondisional untuk mengecek syarat dan menentukan tindakan alternatif</p>
<p>% : operator modulus (sisa bagi) yang digunakan untuk mengecek apakah num % 2 != 0 (jika dibagi 2 sisanya tidak 0, berarti ganjil)</p>
<p>:= : operator deklarasi variabel singkat</p>
<p>++ : operator increment untuk menambah nilai 1</p>
<p>< / > / != / == : operator perbandingan matematika</p>
<p>dalam code di atas, program meminta input n sebagai penentu berapa kali proses (baris) akan diulang. Di setiap perulangan barisnya, program meminta input m untuk mengetahui berapa banyak angka yang akan dimasukkan.</p>
<p>Kemudian, program menjalankan perulangan untuk membaca angka tersebut satu per satu (disimpan di variabel num). Setiap kali angka masuk, program mengeceknya: jika ganjil (num%2 != 0), angka dimasukkan ke arrGanjil dan penghitung nGanjil bertambah. Jika genap (else), angka dimasukkan ke arrGenap dan penghitung nGenap bertambah.</p>
<p>Setelah pemisahan selesai, program memanggil fungsi sortGanjil untuk mengurutkan array ganjil dengan mencari nilai terkecil (idxMin) dan memindahkannya ke depan. Program juga memanggil fungsi sortGenap yang bekerja sebaliknya, yaitu mencari nilai terbesar (idxMax) untuk dipindahkan ke depan agar urut menurun.</p>
<p>Langkah terakhir adalah mencetak hasilnya. Program menggunakan variabel bantuan bernama "pertama" bernilai 1. Saat mencetak elemen paling pertama dari array ganjil, program tidak menambahkan spasi di depannya dan mengubah nilai "pertama" menjadi 0. Untuk elemen-elemen berikutnya (baik di array ganjil yang tersisa maupun seluruh array genap), program akan selalu mencetak spasi lebih dulu diikuti angkanya. Hal ini memastikan susunan hasil rapi: kelompok ganjil dicetak duluan, disambung langsung oleh kelompok genap, tanpa ada spasi berlebih di awal baris. Setelah semua tercetak, program membuat baris baru (Println).</p>

##### Output
<img width="1920" height="1080" alt="Soal 2" src="https://github.com/user-attachments/assets/c39229a2-1819-417e-9a62-df1abb4508d5" />



[penjelasan]
<h2>Soal 3</h2>
<br>

<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan dan Println</p>
<p>const : kata kunci untuk mendeklarasikan nilai tetap (konstanta)</p>
<p>NMAX = 1000000 : konstanta batas maksimal ukuran array</p>
<p>var arr [NMAX]int : deklarasi variabel array global bernama arr untuk menampung data angka</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func insertionSort() : fungsi buatan untuk mengurutkan isi array dari yang terkecil ke terbesar menggunakan metode Insertion Sort (pengurutan sisip)</p>
<p>func main() : bisa diartikan sebagai “fungsi utama” tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>n, num : variabel untuk menyimpan total jumlah data yang telah dimasukkan (n) dan angka inputan user saat ini (num)</p>
<p>key, j : variabel bantuan di dalam fungsi pengurutan untuk menyimpan nilai angka yang sedang dievaluasi (key) dan indeks pembandingnya (j)</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk memasukan data di terminal</p>
<p>println : untuk menampilkan hasil ke layar dan secara otomatis membuat baris baru</p>
<p>for : perintah perulangan untuk mengeksekusi blok kode secara berulang</p>
<p>if / else : perintah kondisional untuk mengecek syarat dan menentukan tindakan alternatif</p>
<p>!= : operator "tidak sama dengan" (digunakan untuk mengecek apakah input bukan -5313)</p>
<p>== : operator "sama dengan" (digunakan untuk mengecek apakah input user adalah angka 0)</p>
<p>% : operator modulus (sisa bagi) yang digunakan pada n % 2 != 0 untuk mengecek apakah jumlah total data bernilai ganjil</p>
<p>:= : operator deklarasi variabel singkat</p>
<p>++ / -- : operator increment (tambah 1) dan decrement (kurang 1)</p>
<p>> / >= / < : operator perbandingan matematika</p>
<p>+ / - / / : operator matematika dasar (penjumlahan, pengurangan, pembagian)</p>
<p>dalam code di atas, program dibuat untuk menerima input angka dari pengguna secara terus-menerus (looping). Perulangan ini akan berhenti sepenuhnya dan program ditutup hanya jika pengguna memasukkan angka penanda khusus, yaitu -5313.</p>
<p>Selama angka yang dimasukkan bukan -5313 dan bukan 0, program akan menyimpan angka tersebut (num) ke dalam array arr dan menaikkan jumlah total data (n). Namun, jika pengguna memasukkan angka 0, program tidak akan menyimpannya, melainkan menganggapnya sebagai "perintah" untuk menghitung dan mencetak nilai tengah (median) dari angka-angka yang sudah terkumpul sejauh ini.</p>
<p>Saat perintah 0 diberikan dan terdeteksi ada data yang bisa diproses (n > 0), program pertama-tama memanggil fungsi insertionSort. Fungsi ini bekerja mengurutkan array dengan cara mengambil satu per satu angka (key) lalu membandingkannya dengan angka-angka sebelumnya. Jika angka sebelumnya lebih besar, angka tersebut digeser ke kanan, lalu angka key disisipkan di posisi yang tepat. Begitu seterusnya sampai semua data urut membesar.</p>
<p>Setelah array tersusun rapi, program utama mengecek jumlah total datanya (n) untuk mencari nilai median. Jika jumlah datanya ganjil (n%2 != 0), mediannya ada tepat di satu nilai tengah, sehingga program mencetak angka pada indeks (n/2). Namun, jika genap (else), mediannya berada di antara dua angka di tengah, sehingga program menjumlahkan kedua angka tersebut lalu membaginya 2 sebelum dicetak ke layar.</p>

##### Output
<img width="1920" height="1080" alt="Soal 3" src="https://github.com/user-attachments/assets/4ed60e1c-9164-4fe5-adb3-a27a18afe767" />


[penjelasan]
<h2>Soal 4</h2>
<br>
Penjelasan Program (Pengurutan dan Pengecekan Jarak Antar Data)
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan, Print, Printf, dan Println</p>
<p>const NMAX = 1000 : kata kunci untuk mendeklarasikan nilai tetap (konstanta) sebagai batas maksimal ukuran array, yaitu 1000 elemen</p>
<p>type tabInt [NMAX]int : pembuatan tipe data alias (kustom) berupa array bernama tabInt yang menampung bilangan bulat</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func insertionSort() : fungsi buatan untuk mengurutkan isi array dari yang terkecil ke terbesar. Fungsi ini menggunakan parameter pointer (*tabInt) sehingga bisa langsung mengubah susunan array asli di memori</p>
<p>func cekJarak() : fungsi buatan untuk memeriksa apakah selisih (jarak) nilai antar angka yang berdekatan di dalam array selalu konstan/sama</p>
<p>func main() : bisa diartikan sebagai “fungsi utama” tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>arr, n, num : variabel untuk menampung kumpulan data (arr), jumlah total data (n), dan angka inputan user saat ini (num)</p>
<p>key, j, diff, status : variabel tambahan di dalam fungsi untuk menyimpan nilai yang sedang disisipkan (key), indeks pembanding (j), selisih angka awal (diff), dan penanda kondisi/flag (status)</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk membaca input dari terminal</p>
<p>print / printf / println : untuk menampilkan teks menyamping, teks dengan format, dan teks dengan baris baru</p>
<p>for : perintah perulangan untuk mengeksekusi blok kode secara berulang</p>
<p>if / else : perintah kondisional untuk mengecek syarat dan menentukan tindakan alternatif</p>
<p>* (pointer) : digunakan pada parameter fungsi (*tabInt) dan pemanggilan elemen (*arr)[i] untuk mengakses dan memodifikasi data pada memori aslinya</p>
<p>& (ampersand) : digunakan saat memanggil fungsi atau scan untuk mengirim alamat memori variabel</p>
<p>:= : operator deklarasi variabel secara singkat</p>
<p>++ / -- : operator increment (tambah 1) dan decrement (kurang 1)</p>
<p>> / >= / < : operator perbandingan (lebih dari, lebih dari sama dengan, kurang dari)</p>
<p>== / != : operator kesamaan (sama dengan, tidak sama dengan)</p>
<p>- : operator matematika untuk pengurangan</p>
<p>dalam code di atas, program meminta user memasukkan angka secara berulang. Selama angka yang dimasukkan bernilai positif atau 0 (num >= 0), program akan terus menyimpannya ke dalam array arr dan menambah jumlah data (n). Jika user memasukkan angka negatif, perulangan berhenti dan angka tersebut tidak disimpan.</p>
<p>Setelah penginputan selesai, program memanggil fungsi insertionSort(&arr, n). Karena menggunakan pointer (&), array arr di fungsi utama akan langsung disusun ulang dari yang terkecil hingga terbesar oleh fungsi tersebut tanpa perlu perintah return.</p>
<p>Selanjutnya, program utama mencetak seluruh isi array yang sudah terurut rapi, dipisahkan dengan spasi.</p>
<p>Langkah terakhir, program memanggil fungsi cekJarak(arr, n). Fungsi ini mula-mula menghitung selisih antara elemen kedua dan pertama (arr[1] - arr[0]) untuk dijadikan patokan awal (diff), lalu mengatur status = 1 (artinya sementara dianggap jaraknya konstan). Program lalu mengecek selisih elemen-elemen selanjutnya. Jika ditemukan ada selisih dua angka berurutan yang tidak sama dengan diff, maka status langsung diubah menjadi 0 (jarak tidak konstan). Pada akhir pengecekan, jika status masih bernilai 1, program mencetak berapa jarak tetapnya tersebut. Jika tidak, program mencetak bahwa "Data berjarak tidak tetap".</p>

##### Output
<img width="1920" height="1080" alt="Soal 4" src="https://github.com/user-attachments/assets/0c6d7648-b7b9-40e0-b268-eb21b44e50ba" />


[penjelasan]
<h2>Soal 5</h2>
<br>
Penjelasan Program (Manajemen Data Buku dengan Struct, Sort, dan Search)
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan, Printf, dan Println</p>
<p>const nMax = 7919 : konstanta batas maksimal ukuran array</p>
<p>type Buku struct : kata kunci untuk membuat tipe data terstruktur (struct) yang memungkinkan kita mengelompokkan beberapa variabel dengan tipe data berbeda (seperti string untuk judul/penulis dan int untuk tahun/rating) ke dalam satu kesatuan entitas bernama Buku</p>
<p>type DaftarBuku [nMax + 1]Buku : pembuatan tipe data array yang isinya adalah kumpulan struct Buku. Ukurannya nMax + 1 karena array ini menggunakan indeks mulai dari 1 (bukan 0)</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func DaftarkanBuku() : fungsi untuk meminta input seluruh data buku dan menyimpannya ke array</p>
<p>func CetakTerfavorit() : fungsi untuk mencari rating tertinggi dan mencetak buku dengan rating tersebut</p>
<p>func UrutBuku() : fungsi untuk mengurutkan array buku berdasarkan rating dari terbesar ke terkecil (descending) menggunakan algoritma Insertion Sort</p>
<p>func Cetak5Terbaru() : fungsi untuk menampilkan 5 judul buku urutan teratas</p>
<p>func CariBuku() : fungsi untuk mencari buku berdasarkan rating tertentu menggunakan algoritma Binary Search (Pencarian Biner) pada data yang sudah urut menurun</p>
<p>func main() : fungsi utama tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>n, targetRating, pustaka : variabel untuk menyimpan total buku (n), rating yang ingin dicari (targetRating), dan variabel array struktur penyimpan data buku (pustaka)</p>
<p>int / string : tipe data untuk bilangan bulat dan teks</p>
<p>scan / printf / println : perintah untuk memasukkan input, mencetak dengan format khusus (%s untuk string, %d untuk angka bulat), dan mencetak baris baru</p>
<p>for : perintah perulangan untuk mengeksekusi blok kode secara berulang</p>
<p>if / else if / else : perintah kondisional untuk mengecek syarat beruntun</p>
<p>* (pointer) / & (ampersand) : digunakan agar fungsi dapat mengubah isi memori dari array pustaka secara langsung</p>
<p>== / != / < / > / <= / >= : operator perbandingan matematika</p>
<p>+ / - / / : operator matematika dasar</p>
<p>dalam code di atas, program utama mula-mula meminta pengguna memasukkan jumlah buku (n). Program lalu memanggil fungsi DaftarkanBuku di mana program melakukan perulangan dari indeks 1 sampai n untuk meminta detail masing-masing buku (id, judul, penulis, penerbit, eksemplar, tahun, dan rating) secara berurutan. Setelah itu, pengguna diminta memasukkan satu angka (targetRating) yang akan dicari nanti.</p>
<p>Setelah data terkumpul, program memanggil fungsi CetakTerfavorit. Fungsi ini mencari nilai rating paling besar di dalam array. Setelah ketemu, ia mengecek kembali seluruh daftar buku dan mencetak detail buku yang ratingnya sama dengan rating maksimal tersebut.</p>
<p>Selanjutnya, program memanggil fungsi UrutBuku. Fungsi ini menyusun ulang daftar buku menggunakan Insertion Sort, namun dikonfigurasi untuk mengurutkan dari rating yang paling besar ke yang paling kecil (descending). Karena data sudah urut membesar ke mengecil, pemanggilan fungsi Cetak5Terbaru setelahnya akan langsung mencetak maksimal 5 judul buku dengan rating terbaik di daftar tersebut.</p>
<p>Langkah terakhir, program memanggil fungsi CariBuku untuk mencari buku dengan rating sama dengan targetRating. Fungsi ini menggunakan Binary Search yang membelah data menjadi dua secara terus-menerus. Karena susunan datanya menurun, logikanya sedikit dibalik: jika rating tengah (pustaka[med].rating) lebih kecil dari target, maka batas kanan digeser (kanan = med - 1), dan sebaliknya. Jika buku ditemukan, program mencetak detailnya; jika tidak, program mencetak pesan "Tidak ada buku dengan rating seperti itu".</p>

##### Output
<img width="1920" height="1080" alt="Soal 5" src="https://github.com/user-attachments/assets/b083c327-bc76-4dd0-b24d-bbdadef2dd07" />
