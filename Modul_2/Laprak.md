# <h1 align="center">Laporan Praktikum Modul 2</h1>
<p align="center">[Zhafif Iqbal Kurniawan] - [109082500051]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

```

#### soal2.go

```go
package main

import "fmt"

func main() {
	var (
		wrna1, wrna2, wrna3, wrna4 string
		temp                       bool
	)
	fmt.Println("Urutkan wrna sesuai rules")
	temp = true
	x := 1
	for x <= 5 {
		fmt.Print("Percobaan ", x, " : ")
		fmt.Scan(&wrna1, &wrna2, &wrna3, &wrna4)
		if wrna1 != "merah" || wrna2 != "kuning" || wrna3 != "hijau" || wrna4 != "ungu" {
			temp = false
		}
		x++
	}
	fmt.Println("Berhasil :", temp)
}
```

#### soal3.go

```go
package main

import "fmt"

func main() {
	var g, kg, gr, tambahan int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&g)
	kg = g / 1000
	gr = g % 1000
	hasil := kg * 10000
	fmt.Println("Detai berat: Rp.", kg, "kg +", gr, "gr")
	if gr >= 500 && kg <= 10 {
		tambahan = gr * 5
		fmt.Println("Detail biaya: Rp.", hasil, " + Rp.", tambahan)
		hasil = hasil + tambahan
		fmt.Print("Total biaya: ", hasil)
	} else if gr < 500 && kg <= 10 {
		tambahan = gr * 15
		fmt.Println("Detail biaya: Rp.", hasil, " + Rp.", tambahan)
		hasil = hasil + tambahan
		fmt.Print("Total biaya: ", hasil)
	} else if kg > 10 {
		fmt.Println("Detail biaya: Rp.", hasil, " + Rp.", tambahan)
		fmt.Print("Total biaya: ", hasil)
	}
}

```
### Output Unguided :

[penjelasan]
<h2>Soal 1</h2>
<br>
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) dan untuk menjalankan operasi input dan output seperti Scan dan Print</p>
<p>func main() : bisa diartikan sebagai “fungsi utama”</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>satu, dua, tiga : variabel bertipe string</p>
<p>temp : variabel bertipe string</p>
<p>string : tipe data untuk teks</p>
<p>Println / Print : untuk menampilkan hasil atau kalimat ke layar (Println membuat baris baru di akhir, Print tidak)</p>
<p>Scan : untuk memasukan data di terminal</p>

<p>Dalam code di atas, program meminta user untuk menginputkan tiga buah kata satu per satu yang disimpan ke dalam variabel satu, dua, dan tiga
Program kemudian mencetak urutan kata tersebut persis seperti saat diinputkan sebagai "Output awal"
Setelah itu, program melakukan proses pergeseran urutan nilai variabel menggunakan variabel bantuan temp:</p>

<p>Nilai awal pada variabel satu diamankan terlebih dahulu ke dalam variabel temp
Posisi variabel satu yang sudah digandakan nilainya itu kemudian ditimpa oleh nilai dari variabel dua
Posisi variabel dua selanjutnya ditimpa oleh nilai dari variabel tiga
Terakhir, posisi variabel tiga diisi dengan nilai dari temp (yang merupakan nilai asli dari kata satu di awal)
Setelah proses pergeseran ini selesai, program mencetaknya sebagai "Output akhir", di mana urutan kata tersebut telah bergeser posisinya.</p>

##### Output
<img width="1920" height="1080" alt="Soal 2B" src="https://github.com/user-attachments/assets/30dd0c2c-2b63-44f6-927f-5c94b806d060" />

[penjelasan]
<h2>Soal 2</h2>
<br>
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) dan untuk menjalankan operasi input dan output seperti Scan dan Print</p>
<p>func main() : bisa diartikan sebagai “fungsi utama”</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>wrna1, wrna2, wrna3, wrna4 : 4 variabel yang dibuat bertipe string</p>
<p>temp : variabel yang dibuat bertipe boolean</p>
<p>string : tipe data untuk teks</p>
<p>bool : tipe data boolean yang hanya memiliki dua nilai, yaitu benar (true) atau salah (false)</p>
<p>Println / Print : untuk menampilkan hasil atau kalimat ke layar (Println membuat baris baru di akhir, Print tidak)</p>
<p>:= : operator short declaration, digunakan untuk membuat variabel baru sekaligus mengisi nilainya (di sini membuat variabel x dengan nilai awal 1)</p>
<p>for : perintah untuk melakukan perulangan (looping)</p>
<p><= : operator perbandingan "Kurang dari atau sama dengan"</p>
<p>Scan : untuk memasukan data di terminal</p>
<p>if : struktur percabangan logika</p>
<p>!= : operator perbandingan "Tidak sama dengan"</p>
<p>|| : operator logika "ATAU" (OR)</p>
<p>++ : operator tambah 1</p>

<p>Dalam code di atas, program meminta user untuk menguji urutan warna sebanyak 5 kali percobaan (diatur oleh for x <= 5)
Di awal, variabel temp disetel menjadi true, artinya program mengasumsikan user "berhasil" terlebih dahulu
Di dalam setiap perulangan:</p>

<p>Program meminta user menginput 4 nama warna sekaligus (wrna1, wrna2, wrna3, wrna4)
Program mengecek kondisi kesalahan menggunakan if: Apakah wrna1 bukan "merah" ATAU wrna2 bukan "kuning" ATAU wrna3 bukan "hijau" ATAU wrna4 bukan "ungu"
Jika ada satu saja warna yang diinput tidak sesuai dengan urutan tersebut pada percobaan ke berapapun, maka status temp akan langsung diubah menjadi false (gagal)
Setelah 5 kali percobaan selesai, program keluar dari loop dan mencetak hasil akhirnya. Jika selama 5 kali berturut-turut inputnya selalu benar ("merah kuning hijau ungu"), maka hasil akhirnya true. Jika ada satu saja yang meleset, hasil akhirnya false.</p>

##### Output
<img width="1920" height="1080" alt="Soal 2C" src="https://github.com/user-attachments/assets/58e8c57b-3bca-4b2d-9f42-30da52026476" />

[penjelasan]
<h2>Soal 3</h2>
<br>
<p>package main : ini adalah paket</p>
<p>import "fmt": Perintah ini mengimpor paket fmt (format) dan untuk menjalankan operasi input dan output seperti Scan dan Print</p>
<p>func main() : bisa diartikan sebagai “fungsi utama”</p>
<p>var : kata kunci untuk mendeklarasikan variabel</p>
<p>g, kg, gr, tambahan : empat variabel yang dibuat</p>
<p>int : tipe data untuk bilangan bulat</p>
<p>scan : untuk memasukan data di terminal</p>
<p>print : untuk menampilkan hasil atau kalimat</p>
<p>if : perintah kondisional untuk menjalankan kode jika suatu syarat terpenuhi</p>
<p>else if : syarat ke 2 atau setelah if yang pertama</p>
<p>else : jika semua kondisi if tidak terpenuhi</p>
<p>% : modulus</p>
<p>&& : DAN (AND)</p>
<p>> : operator “Lebih dari”</p>
<p>< : operator “Kurang dari”</p>
<p>>= : operator “lebih besar sama dengan”</p>
<p><= : operator “lebih kecil sama dengan”</p>

<p>dalam code di atas terdapat variabel  g, kg, gr, dan tambahan. Di saat code nya di jalankan user memasukkan berat dalam gram (g). Program kemudian memecahnya menjadi kilogram (kg) dan sisa gram (gr) menggunakan rumus pembagian dan sisa bagi (modulus). Biaya dasar dihitung 10.000 per kg.
Kemudian program masuk ke logika if:</p>

<p>Jika sisa gram (gr) >= 500 DAN total berat (kg) < 10, maka biaya tambahan adalah gr * 5.</p>
<p>Jika sisa gram (gr) < 500 DAN total berat (kg) < 10, maka biaya tambahan adalah gr * 15.</p>
<p>Jika total berat (kg) > 10, maka tidak ada biaya tambahan (gratis sisa gram).</p>
<p>Terakhir, program menjumlahkan biaya dasar dengan biaya tambahan dan mencetak totalnya.</p>