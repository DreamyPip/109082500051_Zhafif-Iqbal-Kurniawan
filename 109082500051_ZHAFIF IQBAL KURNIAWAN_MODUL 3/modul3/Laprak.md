# <h1 align="center">Laporan Praktikum Modul 3 - ABSTRACT DATA TYPE</h1>

<p align="center">Zhafif Iqbal Kurniawan - 109082500051</p>

## Dasar Teori

isi dengan penjelasan dasar teori disertai referensi jurnal (gunakan kurung siku [] untuk pernyataan yang mengambil refernsi dari jurnal).
contoh :
Linked list atau yang disebut juga senarai berantai adalah Salah satu bentuk struktur data yang berisi kumpulan data yang tersusun secara sekuensial, saling bersambungan, dinamis, dan terbatas[1]. Linked list terdiri dari sejumlah node atau simpul yang dihubungkan secara linier dengan bantuan pointer.

### A. ...<br/>

...

#### 1. ...
Tipe Data Abstrak merupakan sebuah pemodelan matematika yang mendefinisikan struktur informasi beserta sekumpulan operasinya tanpa memaparkan kerumitan detail implementasi teknis di balik layar memori penggunaannya [1].

#### 2. ...
Struktur Data Bentukan adalah fasilitas pemrograman yang memungkinkan pengelompokan berbagai variabel dengan tipe dasar berbeda ke dalam satu kesatuan nama utuh untuk memodelkan entitas dunia nyata secara rapi [2].

#### 3. ...
Pemrograman Modular bertindak sebagai paradigma pengembangan yang membagi program besar menjadi beberapa unit independen melalui pemisahan berkas antarmuka dan implementasi guna meningkatkan keterbacaan serta kemudahan pemeliharaan kode [3].

### B. ...<br/>

...

#### 1. ...
Transmisi parameter berbasis referensi memungkinkan pengiriman alamat memori variabel sebagai argumen fungsi sehingga subprogram dapat memanipulasi data asli secara langsung tanpa perlu menduplikasi keseluruhan struktur data [4].

#### 2. ...
Larik dua dimensi beroperasi sebagai kumpulan elemen data berindeks ganda yang tersusun berurutan di dalam memori dan direpresentasikan secara logis dalam bentuk baris serta kolom untuk mengelola matriks komputasi terstruktur [5].

#### 3. ...
Enkapsulasi informasi berfungsi sebagai prinsip perancangan sistem yang membatasi interaksi langsung dari luar terhadap struktur data dengan memastikan modifikasi nilai hanya terjadi melalui prosedur fungsi resmi demi menjaga integritas program [6].

## Guided

### 1. ...

## mahasiswa.h

```C++
#ifndef MAHASISWA_H_INCLUDED
#define MAHASISWA_H_INCLUDED
struct mahasiswa
{
    char nim[10];
    int nilai1, nilai2;
};
void inputMhs(mahasiswa &m);
float rata2(mahasiswa m);
#endif
```
## mahasiswa.cpp

```C++
#include <iostream>
#include "mahasiswa.h"

using namespace std;

void inputMhs(mahasiswa &m) {
    cout << "input nim = ";
    cin >> m.nim;
    cout << "input nilai1 = ";
    cin >> m.nilai1;
    cout << "input nilai2 = ";
    cin >> m.nilai2;
}

float rata2(mahasiswa m) {
    return float(m.nilai1 + m.nilai2) / 2;
}
```
## main.cpp

```C++
#include <iostream>
#include "mahasiswa.h"

using namespace std;

int main() {
    mahasiswa mhs;
    inputMhs(mhs);
    cout << "rata-rata = " << rata2(mhs);
    return 0;
}
```

Kode ini mendemonstrasikan pemrograman modular menggunakan struktur data struct yang dipisah ke dalam tiga file, yaitu mahasiswa.h, mahasiswa.cpp, dan main.cpp. File implementasi mendefinisikan prosedur inputMhs dengan metode pass-by-reference untuk mengisi data memori variabel, serta fungsi rata2 untuk menghitung rata-rata nilai. Program utama kemudian menginstansiasi struct tersebut dan memanggil fungsinya untuk menampilkan hasil kalkulasi akhir ke layar.


## Unguided

### 1. (isi dengan soal unguided 1)
## nilaiMahasiswa.h

```C++
#ifndef NILAIMAHASISWA_H
#define NILAIMAHASISWA_H
#include <string>

using namespace std;

struct Mhs {
    string nama;
    string nim;
    float uts;
    float uas;
    float tugas;
    float akhir;
};

void inputData(Mhs &m);
void tampilData(Mhs m);

#endif
```
## nilaiMahasiswa.cpp

```C++
#include <iostream>
#include "nilaiMahasiswa.h"

using namespace std;

void inputData(Mhs &m) {
    cout << "Masukkan Nama : ";
    cin.ignore();
    getline(cin, m.nama);
    cout << "Masukkan NIM  : ";
    cin >> m.nim;
    cout << "Nilai UTS     : ";
    cin >> m.uts;
    cout << "Nilai UAS     : ";
    cin >> m.uas;
    cout << "Nilai Tugas   : ";
    cin >> m.tugas;
    
    m.akhir = (0.3 * m.uts) + (0.4 * m.uas) + (0.3 * m.tugas);
}

void tampilData(Mhs m) {
    cout << "Nama        : " << m.nama << endl;
    cout << "NIM         : " << m.nim << endl;
    cout << "Nilai Akhir : " << m.akhir << endl;
}
```
## main.cpp

```C++
#include <iostream>
#include "nilaiMahasiswa.h"

using namespace std;

int main() {
    Mhs data[10];
    int n = 0;
    
    cout << "Berapa data mahasiswa yang ingin dimasukkan (max 10)? : ";
    cin >> n;
    
    if(n > 10) n = 10;
    
    for(int i = 0; i < n; i++) {
        cout << "\nData Mahasiswa ke-" << i+1 << endl;
        inputData(data[i]);
    } 
    
    for(int i = 0; i < n; i++) {
        tampilData(data[i]);
    }
    
    return 0;
}
```

### Output Unguided 1 :

##### Output 1

<img width="1736" height="708" alt="Soal1_1" src="https://github.com/user-attachments/assets/b904411b-4d44-4cbe-8893-835b7ed18120" />


##### Output 2

<img width="1625" height="497" alt="Soal1_2" src="https://github.com/user-attachments/assets/a3eae2a8-36a1-4d19-a187-b5f034e88408" />


Program ini menggunakan array dari tipe data struktur untuk menyimpan sekumpulan data mahasiswa beserta kalkulasi nilai akhirnya secara terorganisir. Proses pengisian data pada setiap elemen array dilakukan melalui parameter reference agar perubahan memori langsung tersimpan pada variabel utama. Setelah batas jumlah mahasiswa yang ditentukan selesai diinput, program menggunakan perulangan untuk mencetak seluruh rekapitulasi data tersebut ke layar secara berurutan.

### 2. (isi dengan soal unguided 2)


```C++
#ifndef PELAJARAN_H
#define PELAJARAN_H
#include <string>

using namespace std;

struct pelajaran {
    string namaMapel;
    string kodeMapel;
};

pelajaran create_pelajaran(string namapel, string kodepel);
void tampil_pelajaran(pelajaran pel);

#endif
```


```C++
#include <iostream>
#include "pelajaran.h"

using namespace std;

pelajaran create_pelajaran(string namapel, string kodepel) {
    pelajaran p;
    p.namaMapel = namapel;
    p.kodeMapel = kodepel;
    return p;
}

void tampil_pelajaran(pelajaran pel) {
    cout << "nama pelajaran : " << pel.namaMapel << endl;
    cout << "nilai : " << pel.kodeMapel << endl;
}
```


```C++
#include <iostream>
#include <string>
#include "pelajaran.h"

using namespace std;

int main() {
    string namapel = "Struktur Data";
    string kodepel = "STD";
    pelajaran pel = create_pelajaran(namapel, kodepel);
    tampil_pelajaran(pel);
    
    return 0;
}
```

### Output Unguided 2 :

##### Output 1

<img width="1560" height="140" alt="Soal2_1" src="https://github.com/user-attachments/assets/208fce6c-fe97-468f-b143-141a810d95ae" />


##### Output 2

<img width="1552" height="202" alt="Soal2_2" src="https://github.com/user-attachments/assets/f28e6129-e924-4225-9600-dd8e337192db" />


Program ini mengimplementasikan tipe data abstrak dengan memisahkan kerangka struktur dan prototipe fungsinya ke dalam file tajuk terpisah. Fungsi pembuatan data bertugas menerima dua argumen teks untuk menyusun objek pelajaran baru yang kemudian dikembalikan sebagai nilai utuh ke program utama. Objek tersebut selanjutnya dikirim menuju prosedur penampilan yang akan mengekstrak serta mencetak atribut nama dan kode pelajaran ke layar sesuai format yang diminta.

### 3. (isi dengan soal unguided 3)
## array.h

```C++
#ifndef ARRAY_H
#define ARRAY_H

void tampilArray(int arr[3][3]);
void tukarPosisiArray(int arr1[3][3], int arr2[3][3], int baris, int kolom);
void tukarPointer(int *p1, int *p2);

#endif
```
## array.cpp

```C++
#include <iostream>
#include "array.h"

using namespace std;

void tampilArray(int arr[3][3]) {
    for (int i = 0; i < 3; i++) {
        for (int j = 0; j < 3; j++) {
            cout << arr[i][j] << " ";
        }
        cout << endl;
    }
}

void tukarPosisiArray(int arr1[3][3], int arr2[3][3], int baris, int kolom) {
    int temp = arr1[baris][kolom];
    arr1[baris][kolom] = arr2[baris][kolom];
    arr2[baris][kolom] = temp;
}

void tukarPointer(int *p1, int *p2) {
    int temp = *p1;
    *p1 = *p2;
    *p2 = temp;
}
```
## main.cpp

```C++
#include <iostream>
#include "array.h"

using namespace std;

int main() {
    int matriks1[3][3] = {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}};
    int matriks2[3][3] = {{9, 8, 7}, {6, 5, 4}, {3, 2, 1}};
    int a = 10, b = 20;
    int *ptr1 = &a;
    int *ptr2 = &b;

    cout << "Matriks 1 awal:\n";
    tampilArray(matriks1);
    cout << "\nMatriks 2 awal:\n";
    tampilArray(matriks2);

    tukarPosisiArray(matriks1, matriks2, 1, 1);

    cout << "\nMatriks 1 setelah pertukaran indeks baris 1 dan kolom 1:\n";
    tampilArray(matriks1);
    cout << "\nMatriks 2 setelah pertukaran indeks baris 1 dan kolom 1:\n";
    tampilArray(matriks2);

    cout << "\nNilai penunjuk memori awal penunjuk pertama berisi " << *ptr1 << " dan penunjuk kedua berisi " << *ptr2 << endl;
    tukarPointer(ptr1, ptr2);
    cout << "Nilai penunjuk memori akhir penunjuk pertama berisi " << *ptr1 << " dan penunjuk kedua berisi " << *ptr2 << endl;

    return 0;
}
```
### Output Unguided 3 :

##### Output 1

<img width="1641" height="557" alt="Soal3_1" src="https://github.com/user-attachments/assets/76bf1008-8c85-46bc-870d-7d842476d530" />


##### Output 2

<img width="1522" height="536" alt="Soal3_2" src="https://github.com/user-attachments/assets/8faac01c-b4e8-41b8-8056-77bc115ed595" />


Program ini memisahkan deklarasi dan implementasi fungsi manipulasi matriks serta penunjuk memori ke dalam tiga berkas berbeda yaitu tajuk dan sumber utama. Berkas antarmuka mendefinisikan prototipe prosedur pencetakan array dan pertukaran nilai yang kemudian logika pemrosesannya dijabarkan secara rinci pada berkas implementasi. Berkas utama selanjutnya bertugas menginisialisasi matriks dan variabel penunjuk memori untuk dieksekusi menggunakan modul yang telah diimpor tersebut.

## Kesimpulan
implementasi tipe data abstrak dan arsitektur pemrograman modular melalui pemisahan kerangka kerja ke dalam beberapa berkas khusus. Pengorganisasian kode yang terstruktur ini dipadukan dengan manipulasi larik serta penunjukan referensi memori secara langsung guna menciptakan sistem pengolahan data yang sangat efisien dan teratur.
...

## Referensi

[1] Haryanto, D., dan Riyadi, S. (2020). Analisis Penerapan Tipe Data Abstrak pada Pengembangan Sistem Informasi. Jurnal Ilmu Komputer dan Informatika, 8(2), 112-119.
<br>[2] Susanti, E., dan Wibowo, A. (2019). Implementasi Struktur Data Majemuk dalam Pembuatan Aplikasi Manajemen Akademik. Jurnal Teknologi Informasi, 7(1), 55-61.
<br>[3] Pratama, R., dan Nurhayati, N. (2021). Evaluasi Pemrograman Modular Berbasis Bahasa Tingkat Tinggi untuk Efisiensi Pengembangan Perangkat Lunak. Jurnal Edukasi dan Penelitian Sistem Komputer, 9(3), 201-208.
<br>[4] Firmansyah, A., dan Lestari, D. (2018). Optimasi Penggunaan Memori Melalui Mekanisme Pass by Reference pada Algoritma Pencarian. Jurnal Rekayasa Sistem Komputer, 5(2), 77-84.
<br>[5] Wijaya, K., dan Santoso, B. (2022). Pengolahan Data Matriks Menggunakan Larik Dua Dimensi pada Pemrograman Terstruktur. Jurnal Komputasi Terapan, 10(1), 34-41.
<br>[6] Setiawan, H., dan Arifin, Z. (2020). Penerapan Konsep Penyembunyian Informasi pada Arsitektur Perangkat Lunak Skala Menengah. Jurnal Sistem Informasi Bisnis, 11(2), 150-157.
