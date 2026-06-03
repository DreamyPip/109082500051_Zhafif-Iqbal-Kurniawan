# <h1 align="center">Laporan Praktikum Modul 12</h1>
<p align="center">[Zhafif Iqbal Kurniawan] - [109082500051]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

const NCALON int = 20

type arrKandidat [NCALON]int

func seqSearch(T arrKandidat, n int, X int) int {
    var found int = -1
    var j int = 0
    for j < n && found == -1 {
        if T[j] == X {
            found = j
        }
        j = j + 1
    }
    return found
}

func main() {
    var daftar arrKandidat
    var rekap [NCALON]int
    var suara, suaraMasuk, suaraSah int

    for i := 0; i < NCALON; i++ {
        daftar[i] = i + 1
    }

    for {
        fmt.Scan(&suara)
        if suara == 0 {
            break
        }

        suaraMasuk++
        idx := seqSearch(daftar, NCALON, suara)

        if idx != -1 {
            suaraSah++
            rekap[idx]++
        }
    }

    fmt.Printf("Suara masuk: %d\n", suaraMasuk)
    fmt.Printf("Suara sah: %d\n", suaraSah)

    for i := 0; i < NCALON; i++ {
        if rekap[i] > 0 {
            fmt.Printf("%d: %d\n", daftar[i], rekap[i])
        }
    }
}

```

#### soal2.go

```go
package main

import "fmt"

const NCALON int = 20

type arrKandidat [NCALON]int

func seqSearch(T arrKandidat, n int, X int) int {
	var found int = -1
	var j int = 0
	for j < n && found == -1 {
		if T[j] == X {
			found = j
		}
		j = j + 1
	}
	return found
}

func main() {
	var daftar arrKandidat
	var rekap [NCALON + 1]int
	var suara, suaraMasuk, suaraSah int

	for i := 0; i < NCALON; i++ {
		daftar[i] = i + 1
	}

	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}

		suaraMasuk++

		idx := seqSearch(daftar, NCALON, suara)

		if idx != -1 {
			suaraSah++
			rekap[daftar[idx]]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	var max1, max2 int = -1, -1
	var id1, id2 int = 0, 0

	for i := 1; i <= NCALON; i++ {
		if rekap[i] > max1 {
			max2 = max1
			id2 = id1
			max1 = rekap[i]
			id1 = i
		} else if rekap[i] > max2 {
			max2 = rekap[i]
			id2 = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", id1)
	fmt.Printf("Wakil ketua: %d\n", id2)
}

```

#### soal3.go

```go
package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	isiArray(n)

	idx := posisi(n, k)
	if idx != -1 {
		fmt.Println(idx)
	} else {
		fmt.Println("TIDAK ADA")
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	var found int = -1
	var med int
	var kr int = 0
	var kn int = n - 1

	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if k < data[med] {
			kn = med - 1
		} else if k > data[med] {
			kr = med + 1
		} else {
			found = med
		}
	}
	
	return found
}

```

[penjelasan]
<h2>Soal 1</h2>
<br>
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan dan Printf</p>
<p>const : kata kunci untuk mendeklarasikan nilai tetap (konstanta) yang tidak bisa diubah nilainya</p>
<p>NCALON : konstanta yang menyimpan batas maksimal jumlah calon kandidat yaitu 20</p>
<p>type : kata kunci untuk membuat tipe data bentukan atau alias baru</p>
<p>arrKandidat [NCALON]int : tipe data array buatan bernama arrKandidat yang bisa menyimpan bilangan bulat hingga batas NCALON (20 elemen)</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func seqSearch() : fungsi buatan untuk mencari letak (indeks) sebuah data di dalam array menggunakan metode pencarian berurutan (sequential search)</p>
<p>func main() : bisa diartikan sebagai “fungsi utama” tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>daftar, rekap, suara, suaraMasuk, suaraSah : variabel untuk menyimpan daftar kandidat valid, jumlah perolehan tiap kandidat (rekap), input suara, total seluruh suara masuk, dan total suara yang sah</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk memasukan data di terminal</p>
<p>printf : untuk menampilkan teks dan hasil dengan format tertentu (seperti %d untuk menampilkan bilangan bulat dan \n untuk baris baru)</p>
<p>for : perintah perulangan untuk mengeksekusi blok kode berkali-kali</p>
<p>for { ... } : perulangan tanpa batas (infinite loop) yang akan terus berjalan dan meminta input sampai dihentikan paksa oleh perintah break</p>
<p>break : perintah untuk menghentikan dan keluar dari perulangan secara paksa</p>
<p>:= : operator untuk mendeklarasikan sekaligus mengisi nilai variabel baru secara singkat</p>
<p>++ : operator increment untuk menambah nilai variabel sebanyak 1</p>
<p>if : perintah kondisional untuk mengecek sebuah syarat</p>
<p>== : operator "sama dengan" (misalnya untuk mengecek apakah data cocok, atau apakah input user bernilai 0)</p>
<p>!= : operator "tidak sama dengan" (digunakan untuk mengecek apakah hasil pencarian indeks valid/tidak bernilai -1)</p>
<p>&& : operator logika "DAN" (AND), yang berarti kedua syarat di kiri dan kanannya harus terpenuhi</p>
<p>return : perintah untuk mengembalikan nilai hasil perhitungan atau pencarian dari dalam fungsi ke pemanggilnya</p>
<p>dalam code di atas, program utama mula-mula mengisi array daftar dengan nomor urut kandidat dari 1 sampai 20 (berdasarkan batas NCALON). Kemudian, program masuk ke dalam perulangan tanpa batas untuk terus-menerus meminta input suara (nomor kandidat) dari user.</p>
<p>Setiap kali suara dimasukkan, variabel suaraMasuk akan bertambah. Jika user memasukkan angka 0, perintah break akan dijalankan untuk menghentikan proses input. Namun jika bukan 0, program akan memanggil fungsi seqSearch untuk mengecek apakah nomor suara yang dimasukkan terdaftar sebagai kandidat yang sah (antara 1-20).</p>
<p>Fungsi seqSearch bekerja dengan cara menelusuri nomor kandidat satu per satu dari awal. Jika nomornya cocok, fungsi akan mengembalikan letak indeksnya. Jika indeksnya valid (ditandai dengan nilai yang tidak sama dengan -1), maka suara tersebut dihitung sebagai suaraSah, dan perolehan suara kandidat tersebut di array rekap akan ditambahkan 1 (+1).</p>
<p>Setelah proses input suara selesai (saat user memasukkan 0), program akan mencetak total suara yang masuk dan total suara yang sah. Terakhir, program mencetak rekapitulasi hasil pemilu, menampilkan hanya kandidat-kandidat yang berhasil mendapatkan minimal 1 suara atau lebih.</p>
##### Output
<img width="1920" height="1080" alt="Soal 1" src="https://github.com/user-attachments/assets/a78145a0-37fa-418f-a927-9674ed9409b4" />



<br>
<br>
[penjelasan]
<h2>Soal 2</h2>
<br>
Penjelasan Program (Pemilihan Ketua dan Wakil Ketua RT)
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan dan Printf</p>
<p>const : kata kunci untuk mendeklarasikan nilai tetap (konstanta) yang tidak bisa diubah nilainya</p>
<p>NCALON : konstanta yang menyimpan batas maksimal jumlah calon kandidat yaitu 20</p>
<p>type : kata kunci untuk membuat tipe data bentukan atau alias baru</p>
<p>arrKandidat [NCALON]int : tipe data array buatan bernama arrKandidat yang bisa menyimpan bilangan bulat hingga 20 elemen</p>
<p>func : kata kunci untuk membuat fungsi atau sub-program terpisah</p>
<p>func seqSearch() : fungsi buatan untuk mencari letak (indeks) sebuah data di dalam array menggunakan metode pencarian berurutan (sequential search)</p>
<p>func main() : bisa diartikan sebagai “fungsi utama” tempat program mulai berjalan</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>daftar, rekap, suara, suaraMasuk, suaraSah : variabel untuk menyimpan daftar kandidat valid, jumlah perolehan tiap kandidat, input suara user, total suara masuk, dan total suara sah</p>
<p>max1, max2, id1, id2 : variabel tambahan khusus untuk menampung perolehan suara terbanyak pertama (max1) dan kedua (max2), beserta nomor kandidatnya (id1 untuk ketua, id2 untuk wakil)</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk memasukan data di terminal</p>
<p>printf : untuk menampilkan teks dan hasil dengan format tertentu</p>
<p>for : perintah perulangan untuk mengeksekusi blok kode berkali-kali</p>
<p>for { ... } : perulangan tanpa batas (infinite loop) yang akan terus berjalan dan meminta input sampai dihentikan paksa oleh perintah break</p>
<p>break : perintah untuk menghentikan dan keluar dari perulangan secara paksa</p>
<p>:= : operator untuk mendeklarasikan sekaligus mengisi nilai variabel baru secara singkat</p>
<p>++ : operator increment untuk menambah nilai variabel sebanyak 1</p>
<p>if / else if : perintah kondisional untuk mengecek syarat beruntun, digunakan di sini untuk mencari dua perolehan suara terbesar</p>
<p>== : operator "sama dengan"</p>
<p>!= : operator "tidak sama dengan" (digunakan untuk mengecek jika hasil pencarian valid)</p>
<p>< / <= / > : operator perbandingan matematika (kurang dari, kurang dari atau sama dengan, lebih dari)</p>
<p>&& : operator logika "DAN" (AND)</p>
<p>return : perintah untuk mengembalikan nilai hasil pencarian dari dalam fungsi ke pemanggilnya</p>
<p>dalam code di atas, program memulainya dengan mengisi array daftar kandidat secara berurutan dari angka 1 hingga 20. Setelah itu, program meminta input suara secara terus-menerus (looping tak hingga) sampai pengguna memasukkan angka 0.</p>
<p>Setiap input suara akan dihitung sebagai suaraMasuk. Program lalu mengecek suara tersebut lewat fungsi seqSearch. Jika suara sah (bukan -1), maka suaraSah bertambah, dan kotak penyimpanan (array rekap) untuk kandidat tersebut ditambahkan 1. Array rekap sengaja dibuat berukuran [NCALON + 1] (yaitu 21 elemen) agar indeks array-nya bisa langsung sejajar dengan nomor urut kandidat (1-20).</p>
<p>Setelah pengumpulan suara selesai, program mencetak total suara yang masuk dan total suara yang sah.</p>
<p>Selanjutnya, program memasuki tahap pencarian juara 1 dan 2. Program melakukan perulangan dari kandidat 1 sampai 20 pada array rekap. Jika suara kandidat saat ini (rekap[i]) lebih besar dari juara pertama (max1), maka juara pertama yang lama akan digeser menjadi juara kedua (max2 = max1), dan kandidat saat ini naik menjadi juara pertama yang baru. Namun, jika suaranya tidak mengalahkan juara pertama tapi lebih besar dari juara kedua (else if), maka ia hanya merebut posisi juara kedua.</p>
<p>Pada akhir program, kandidat pemegang posisi pertama (id1) ditetapkan dan dicetak sebagai Ketua RT, sedangkan pemegang posisi kedua (id2) dicetak sebagai Wakil Ketua.</p>
##### Output
<img width="1920" height="1080" alt="Soal 2" src="https://github.com/user-attachments/assets/57e8cd0c-d074-4d81-b796-5658eb8286d2" />


<br>
<br>

[penjelasan]
<h2>Soal 3</h2>
<br>
Penjelasan Program (Pencarian Data dengan Algoritma Binary Search)
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) untuk menjalankan operasi input dan output seperti Scan dan Println</p>
<p>const : kata kunci untuk mendeklarasikan nilai tetap (konstanta)</p>
<p>NMAX = 1000000 : konstanta yang menentukan batas maksimal ukuran array, yaitu 1 juta elemen</p>
<p>var data [NMAX]int : deklarasi variabel array bernama data yang bersifat "global", artinya array ini dibuat di luar fungsi agar bisa diakses dan diubah secara langsung oleh semua fungsi di dalam program</p>
<p>func : kata kunci untuk membuat fungsi (terdapat fungsi main, isiArray, dan posisi)</p>
<p>var n, k, kr, kn, med, found : variabel untuk menyimpan banyak data (n), angka yang dicari (k), batas pencarian kiri (kr), batas pencarian kanan (kn), posisi tengah (med), dan status/indeks penemuan (found)</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk memasukan data di terminal</p>
<p>println : untuk menampilkan hasil ke layar dan membuat baris baru</p>
<p>for : perintah perulangan untuk mengeksekusi blok kode secara berulang</p>
<p>if / else if / else : perintah kondisional bersarang untuk mengecek syarat tertentu, dan memberikan aksi alternatif (else) jika semua syarat sebelumnya tidak terpenuhi</p>
<p>!= : operator "tidak sama dengan" (untuk mengecek apakah indeks hasil pencarian valid)</p>
<p>== : operator "sama dengan"</p>
<p>< / > / <= : operator perbandingan matematika (kurang dari, lebih dari, kurang dari atau sama dengan)</p>
<p>+ / - / / : operator matematika (penjumlahan, pengurangan, pembagian)</p>
<p>return : perintah untuk mengembalikan nilai hasil pencarian dari dalam fungsi ke pemanggilnya</p>
<p>dalam code di atas, program utama meminta user memasukkan jumlah total data (n) dan angka target yang ingin dicari (k). Setelah itu, program memanggil fungsi isiArray.</p>
<p>Di dalam fungsi isiArray, program melakukan perulangan sebanyak n kali untuk meminta user memasukkan angka-angka tersebut satu per satu ke dalam array global bernama data. (Syarat utama algoritma ini: angka-angka yang dimasukkan harus sudah dalam keadaan urut membesar/mengecil).</p>
<p>Setelah array terisi, program memanggil fungsi posisi untuk mencari angka k dengan metode Binary Search. Cara kerjanya: program menetapkan batas awal penelusuran dari ujung paling kiri (kr = 0) dan ujung paling kanan (kn = n - 1). Program lalu mencari posisi tengahnya (med) dengan rumus (kr + kn) / 2.</p>
<p>Selanjutnya, program membandingkan angka yang dicari (k) dengan angka yang ada di posisi tengah (data[med]). Jika k lebih kecil dari angka di tengah, berarti angka k pasti ada di sebelah kiri, sehingga batas kanan (kn) digeser ke med - 1. Jika k lebih besar, pencarian difokuskan ke sebelah kanan dengan menggeser batas kiri (kr) ke med + 1. Jika nilainya pas/sama, maka angka berhasil ditemukan dan nilai found diubah menjadi posisi med tersebut.</p>
<p>Proses pembelahan data ini terus diulang (looping) selama batas kiri tidak melewati batas kanan (kr <= kn). Setelah perulangan selesai, fungsi posisi mengembalikan letak (indeks) angka yang ditemukan. Di program utama, jika indeksnya bukan -1, maka letak indeks tersebut dicetak ke layar. Jika -1, maka dicetak peringatan "TIDAK ADA".</p>

##### Output
<img width="1920" height="1080" alt="Soal 3" src="https://github.com/user-attachments/assets/0631d810-44d2-4664-8c4a-d97f5921a4c9" />
