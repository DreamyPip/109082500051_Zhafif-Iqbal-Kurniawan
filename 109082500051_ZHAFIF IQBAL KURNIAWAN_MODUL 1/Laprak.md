# <h1 align="center">Laporan Praktikum Modul 1 - Codeblocks IDE & Pengenalan Bahas C++ (Bagian Pertama)</h1>

<p align="center">Zhafif Iqbal Kurniawan - 109082500051</p>

## Dasar Teori

isi dengan penjelasan dasar teori disertai referensi jurnal (gunakan kurung siku [] untuk pernyataan yang mengambil refernsi dari jurnal).
contoh :
Linked list atau yang disebut juga senarai berantai adalah Salah satu bentuk struktur data yang berisi kumpulan data yang tersusun secara sekuensial, saling bersambungan, dinamis, dan terbatas[1]. Linked list terdiri dari sejumlah node atau simpul yang dihubungkan secara linier dengan bantuan pointer.

### A. ...<br/>

...

#### 1. ...
Tipe data adalah pengklasifikasian nilai yang menentukan batasan memori yang dialokasikan serta jenis operasi aritmatika maupun logika yang dapat diterapkan pada data tersebut di dalam program [1].

#### 2. ...
Percabangan atau pemilihan merupakan struktur instruksi yang mengarahkan alur eksekusi program ke blok kode tertentu, di mana perintah hanya akan dijalankan apabila kondisi atau syarat logikanya bernilai benar [2].

#### 3. ...
Perulangan adalah teknik pemrograman yang mengeksekusi sebuah blok instruksi secara terus-menerus selama kondisi batas yang dievaluasi masih terpenuhi, sehingga penulisan algoritma menjadi lebih efisien [3].

### B. ...<br/>

...

#### 1. ...
Array adalah struktur data statis yang menyimpan sekumpulan elemen dengan tipe data sejenis pada lokasi memori yang berdekatan, di mana setiap nilainya dapat diakses secara langsung menggunakan nomor indeks [4].

#### 2. ...
Struct merupakan tipe data bentukan yang berfungsi untuk membungkus dan mengelompokkan beberapa variabel dengan tipe data yang berbeda-beda ke dalam satu kesatuan nama agar dapat merepresentasikan sebuah objek secara utuh [5].

#### 3. ...
Fungsi adalah blok kode independen yang dirancang terpisah dari program utama untuk menjalankan tugas spesifik, yang bertujuan membuat susunan kode menjadi lebih rapi, modular, dan dapat digunakan berulang kali [6].

## Guided

### 1. ...

```C++
#include <iostream>
using namespace std;
int main(){
int W, X, Y; float Z;
X = 7; Y = 3; W = 1;
Z = (X + Y)/(Y + W);
cout<< "Nilai z = " << Z << endl;
return 0;
}
```

Kode ini menjalankan operasi perhitungan matematika (7 + 3) / (3 + 1) yang menghasilkan pembagian 10 / 4. Mengingat variabel pembentuknya bertipe bilangan bulat (*integer*), program secara otomatis menerapkan aturan pembagian *integer* yang membuang sisa angka desimal. Akibatnya, perhitungan tersebut dibaca sebagai 2, bukan 2.5. Angka 2 inilah yang kemudian disimpan dan dicetak melalui variabel Z, meskipun variabel Z tersebut telah dideklarasikan menggunakan tipe data *float* (desimal).

### 2. ...

```C++
#include <iostream>
using namespace std;
int main(){
int r = 10;
int s;
s=10 + ++r;
cout<< "Nilai r= "<<r<<endl;
cout<< "Nilai s= "<<s<<endl;
return 0;
}
```

Kode ini menunjukkan cara kerja operator pre-increment pada variabel r dalam sebuah operasi penjumlahan. Pada mulanya, variabel r diberi nilai awal 10. Saat mengeksekusi baris s = 10 + ++r, program akan menambahkan nilai r sebesar 1 terlebih dahulu sehingga r berubah menjadi 11 sebelum mengeksekusi penjumlahannya dengan angka 10. Hasil dari perhitungan tersebut, yaitu 10 ditambah 11, kemudian disimpan ke dalam variabel s. Oleh karena itu, hasil akhir yang dicetak pada layar adalah Nilai r= 11 dan Nilai s= 21.

### 3. ...

```C++
#include <iostream>
#include <stdlib.h>
using namespace std;
int main(){
int r = 10;
int s;
s=10 + r++;
cout<< "Nilai r= "<<r<<endl;
cout<< "Nilai s= "<<s<<endl;
return 0;
}
```

Kode ini menunjukkan cara kerja operator post-increment pada variabel r dalam sebuah operasi penjumlahan. Pada awalnya, variabel r diberi nilai 10. Saat program menjalankan perintah s = 10 + r++, program akan menggunakan nilai r yang lama yaitu 10 terlebih dahulu untuk dijumlahkan dengan angka 10. Hasil penjumlahan tersebut, yaitu 20, kemudian disimpan ke dalam variabel s. Setelah proses perhitungan tersebut selesai dieksekusi, barulah nilai r ditambah 1 secara otomatis sehingga berubah menjadi 11. Oleh karena itu, hasil akhir yang akan dicetak pada layar adalah Nilai r= 11 dan Nilai s= 20.

### 4. ...

```C++
#include <iostream>
using namespace std;
int main(){
double tot_pembelian, diskon;
cout<<"total pembelian: Rp";
cin>>tot_pembelian;
diskon = 0;
if(tot_pembelian >= 100000)
diskon = 0.05*tot_pembelian;
cout<<"besar diskon = Rp" <<diskon;
}
```

Kode ini bertujuan untuk menghitung besaran potongan harga berdasarkan total belanjaan pengguna. Program awalnya akan meminta pengguna memasukkan nominal total pembelian yang kemudian disimpan ke dalam variabel tot_pembelian, sementara variabel diskon diatur dengan nilai awal 0. Selanjutnya, program menggunakan struktur kondisi if untuk memeriksa apakah total pembelian tersebut bernilai 100000 atau lebih. Jika syarat ini terpenuhi, program akan mengubah nilai diskon dengan menghitung 5 persen dari total pembelian tersebut. Pada akhir proses, layar akan menampilkan besaran nominal diskon yang berhasil didapatkan oleh pengguna.

### 5. ...

```C++
#include <iostream>
using namespace std;
int main(){
double tot_pembelian, diskon;
cout<<"total pembelian: Rp";
cin>>tot_pembelian;
diskon = 0;
if(tot_pembelian >= 100000)
diskon = 0.05*tot_pembelian;
else
diskon = 0;
cout<<"besar diskon = Rp" <<diskon;
}
```

Kode ini merupakan pengembangan dari program sebelumnya dengan menambahkan struktur percabangan else. Program tetap mengevaluasi apakah total pembelian mencapai angka 100000 atau lebih untuk memberikan diskon sebesar 5 persen. Perbedaannya, jika total belanja berada di bawah nominal tersebut, program kini secara eksplisit akan diarahkan masuk ke dalam blok else yang menegaskan bahwa nilai diskon ditetapkan menjadi 0. Penambahan fungsi else ini berguna untuk memberikan instruksi alternatif yang jelas apabila syarat utama pada bagian if tidak berhasil terpenuhi.

### 6. ...

```C++
#include <iostream>
using namespace std;
int main(){
int kode_hari;
puts("Menentukan hari kerja/libur\n");
puts("1=Senin 3=Rabu 5=Jumat 7=Minggu ");
puts("2=Selasa 4=Kamis 6=Sabtu ");
cin>>kode_hari;
switch(kode_hari){
case 1:
case 2:
case 3:
case 4:
case 5:
cout<<"Hari Kerja"<<endl;
break;
case 6:
case 7:
cout<<"Hari Libur"<<endl;
break;
default:
cout<<"Kode masukan salah!!!"<<endl;
}
return 0;
}
```

Kode ini menggunakan struktur switch-case untuk mengelompokkan input angka 1 sampai 7. Kasus 1 hingga 5 ditumpuk agar secara bersamaan mencetak teks Hari Kerja, sedangkan kasus 6 dan 7 ditumpuk untuk mencetak teks Hari Libur. Apabila angka yang dimasukkan di luar pilihan tersebut, program otomatis menjalankan perintah default untuk menampilkan pesan peringatan.

### 7. ...

```C++
#include <iostream>
using namespace std;
int main(){
int jum;
cout<<"jumlah perulangan: ";
cin>>jum;
for(int i=0; i<jum; i++){
cout<<"saya pintar\n";
}
return 0;
}
```

Kode ini menggunakan perulangan for untuk mencetak teks saya pintar secara berulang sesuai dengan angka yang dimasukkan oleh pengguna. Program awalnya meminta input angka dari pengguna untuk disimpan ke dalam variabel jum. Setelah itu, instruksi for akan mengeksekusi perintah pencetakan teks tersebut secara terus-menerus hingga jumlah perulangannya persis mencapai batas nilai jum yang telah ditentukan.

### 8. ...

```C++
#include <iostream>
using namespace std;
int main(){
int i=1;
int jum;
cout<<"masukan banyak baris: ";
cin>>jum;
while(i<=jum){
cout<<"baris ke-"<<i<<endl;
i++; 
}
return 0;
}
```

Kode ini menggunakan struktur perulangan while untuk mencetak teks secara berurutan sesuai dengan batas angka yang dimasukkan oleh pengguna. Program meminta input untuk disimpan dalam variabel jum sebagai penentu batas akhir, sedangkan variabel i disiapkan dengan nilai awal 1. Selama nilai i masih lebih kecil atau sama dengan nilai jum, program akan terus mencetak keterangan baris tersebut dan menambahkan nilai i sebesar satu secara bertahap hingga kondisi perulangannya berhenti.

### 9. ...

```C++
#include <iostream>
using namespace std;
int main(){
int i = 1;
int jum;
cin >> jum;
do{
cout << "baris ke-" <<(i+1)<<endl;
i++;
} while(i<jum);
return 0;
}
```

Kode ini menggunakan perulangan do-while yang akan menjalankan perintah minimal satu kali di awal sebelum memeriksa syarat berhentinya di akhir. Program meminta kita memasukkan angka batas perulangan pada variabel jum. Karena variabel i memiliki nilai awal 1 dan program diinstruksikan mencetak hasil i ditambah 1, maka tulisan pertama yang muncul di layar akan langsung dimulai dari baris ke-2. Selanjutnya, nilai i akan terus bertambah satu, dan pencetakan ini akan diulang selama nilai i masih lebih kecil dari jum.

### 10. ...

```C++
#include <iostream>
#define MAX 5
using namespace std;
int main(){
int i;
struct data{
char nama[40];
int nilai;
};
data siswa[MAX];
for(i=0; i<MAX; i++){
cout<<"masukkan data ke-"<<i+1<<endl;
cout<<"nama = ";
cin>>siswa[i].nama;
cout<<"nilai = ";
cin>>siswa[i].nilai;
}
cout<<"\ndata siswa\n";
cout<<"=======";
for(i=0; i<MAX; i++){
cout<<"\n\ndata ke-"<<i+1;
cout<<"\n\nnama="<<siswa[i].nama;
cout<<"\n\nnilai="<<siswa[i].nilai;
}
return 0;
}
```

Kode ini memakai fitur struct untuk menggabungkan variabel nama dan nilai menjadi satu kesatuan data siswa. Program membatasi jumlah data maksimal sebanyak 5 orang menggunakan perintah define di awal, lalu menyiapkan ruang penyimpanannya menggunakan konsep array. Saat beroperasi, program memakai perulangan for yang pertama untuk meminta kita memasukkan data nama dan nilai kelima siswa tersebut satu per satu secara berurutan. Setelah seluruh kotak data selesai diisi, program akan langsung menggunakan perulangan for yang kedua untuk mencetak ulang semua informasi siswa yang sudah kita ketikkan tadi ke layar.

### 11. ...

```C++
#include <iostream>
using namespace std;

float ctof(float celcius);
int main() {
float celcius, fahrenheit;
cout <<"nilai Celcius? ";
cin >> celcius;
fahrenheit = ctof(celcius);
cout<<celcius<<" Celcius adalah "<<fahrenheit<<" Fahrenheit"<<endl;
return 0;
}

float ctof(float celcius){
return (celcius * 1.8) + 32;
}
```

Kode ini berfungsi untuk mengonversi suhu dari Celcius ke Fahrenheit dengan menggunakan sebuah fungsi terpisah bernama ctof. Program awalnya meminta kita memasukkan angka suhu Celcius yang kemudian dikirim ke dalam fungsi ctof tersebut. Di dalam fungsi itu, angka suhu dihitung menggunakan rumus perkalian 1.8 ditambah 32. Hasil perhitungannya kemudian dikembalikan lagi ke program utama untuk langsung dicetak ke layar.

## Unguided

### 1. (isi dengan soal unguided 1)

```C++
#include <iostream>
using namespace std;

int main() {
    float b1, b2, total;
    cout << "Masukkan bilangan pertama: ";
    cin >> b1;
    cout << "Masukkan bilangan kedua: ";
    cin >> b2;
    total = b1 + b2;
    cout << "Penjumlahan: " << total << endl;
    total = b1 - b2;
    cout << "Pengurangan: " << total << endl;
    total = b1 * b2;
    cout << "Perkalian: " << total << endl;
    total = b1 / b2;
    cout << "Pembagian: " << total << endl;
    return 0;
}
```

### Output Unguided 1 :

##### Output 1

<img width="1658" height="215" alt="Soal1_1" src="https://github.com/user-attachments/assets/cee92872-47d6-4a48-b116-8c004efe7897" />


##### Output 2

<img width="1728" height="203" alt="Soal1_2" src="https://github.com/user-attachments/assets/3aed382c-65f2-43b0-8ae5-2d48763ef21f" />


Kode ini berfungsi sebagai kalkulator sederhana yang dapat memproses angka desimal menggunakan tipe data float. Awalnya, program meminta kita memasukkan dua buah angka yang akan disimpan ke dalam variabel b1 dan b2. Setelah kedua angka tersebut diterima, program langsung menjalankan empat operasi aritmatika dasar secara berurutan, yaitu penjumlahan, pengurangan, perkalian, dan pembagian. Setiap hasil perhitungan tersebut disimpan secara bergantian ke dalam satu variabel bernama total untuk kemudian dicetak ke layar satu per satu.

### 2. (isi dengan soal unguided 2)

```C++
#include <iostream>
#include <string>
using namespace std;

int main() {
    int angka;
    string teks[] = {"nol", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"};
    string kapital[] = {"", "Satu", "Dua", "Tiga", "Empat", "Lima", "Enam", "Tujuh", "Delapan", "Sembilan"};

    cout << "Masukkan angka (0 - 100): ";
    cin >> angka;

    if (angka < 0 || angka > 100) {
        cout << "Hanya angka 0 - 100" << endl;
        return 0;
    }

    cout << angka << " : ";

    if (angka >= 0 && angka <= 11) {
        cout << teks[angka];
    } 
    else if (angka >= 12 && angka <= 19) {
        cout << teks[angka % 10] << " belas";
    } 
    else if (angka >= 20 && angka <= 99) {
        cout << teks[angka / 10] << " puluh";
        if (angka % 10 != 0) {
            cout << " " << kapital[angka % 10];
        }
    } 
    else if (angka == 100) {
        cout << "seratus";
    }

    cout << endl;
    return 0;
}
```

### Output Unguided 2 :

##### Output 1

<img width="1721" height="200" alt="Soal2_1" src="https://github.com/user-attachments/assets/dfea7550-03c7-4a79-98d9-8f284a32760f" />


##### Output 2

<img width="1720" height="191" alt="Soal2_2" src="https://github.com/user-attachments/assets/01c4bbda-388b-4e28-bc3e-943854990d49" />


Kode ini berfungsi untuk mengonversi angka dari 0 hingga 100 menjadi teks ejaan. Program menyiapkan dua buah array untuk menyimpan daftar kosakata angka dasar. Setelah pengguna memasukkan angka, program akan mengecek apakah nilainya berada dalam rentang yang diizinkan. Jika sesuai, program menggunakan serangkaian struktur percabangan untuk menerjemahkan angka tersebut berdasarkan kelompok nilainya. Angka 0 hingga 11 akan langsung mengambil teks dari array, angka belasan diformat menggunakan operasi sisa bagi yang ditambah kata belas, dan angka puluhan dihitung melalui kombinasi hasil pembagian untuk memanggil kata puluh serta operasi sisa bagi untuk memanggil angka satuannya. Khusus untuk angka 100, program akan langsung mencetaknya sebagai seratus.

### 3. (isi dengan soal unguided 3)

```C++
#include <iostream>
using namespace std;

int main() {
    int n;
    
    cout << "input: ";
    cin >> n;
    
    cout << "output:\n";
    
    for (int i = n; i >= 0; i--) {
        
        for (int j = 0; j < n - i; j++) {
            cout << "  ";
        }

        for (int j = i; j >= 1; j--) {
            cout << j << " ";
        }

        cout << "*";

        for (int j = 1; j <= i; j++) {
            cout << " " << j;
        }
        
        cout << endl;
    }

    return 0;
}
```

### Output Unguided 3 :

##### Output 1

<img width="1732" height="208" alt="Soal3_1" src="https://github.com/user-attachments/assets/f8cbe46d-5b4e-47c8-b26a-a679edd7acc3" />


##### Output 2

<img width="1717" height="366" alt="Soal3_2" src="https://github.com/user-attachments/assets/26bf1c66-ae02-4d81-9f06-fe14e6884af2" />


Kode ini berfungsi untuk mencetak pola angka dan simbol berbentuk segitiga terbalik yang simetris berdasarkan angka yang dimasukkan oleh pengguna. Program akan meminta sebuah angka batas awal, lalu menjalankan perulangan for utama yang menghitung mundur untuk mengatur perpindahan baris. Di dalam perulangan utama tersebut, terdapat tiga perulangan for tambahan yang bekerja secara berurutan, yaitu untuk mencetak spasi kosong agar letak karakter semakin menjorok ke tengah, mencetak deret angka yang menurun ke arah pusat, menampilkan simbol bintang sebagai titik tengah, dan diakhiri dengan mencetak deret angka yang menaik ke sisi kanan. Kombinasi bersarang dari perulangan ini secara otomatis menyusun barisan angka yang semakin menyempit pada setiap barisnya hingga mengerucut sempurna pada satu baris terbawah.

## Kesimpulan
dasar pemrograman C++ seperti tipe data, operasi aritmatika, percabangan (if, switch), dan berbagai perulangan (for, while, bersarang(nested loop)). penggunaan array, struct.
...

## Referensi

[1] Sitorus, L., & Tambunan, J. (2020). Analisis Penerapan Algoritma dan Struktur Data dalam Pemrograman Dasar. Jurnal Ilmiah Komputer dan Informatika, 9(2), 55-62.
<br>[2] Supriyanto, A. (2019). Penerapan Struktur Logika Percabangan pada Algoritma Sistem Pengambilan Keputusan. Jurnal Teknologi Informasi dan Ilmu Komputer, 6(4), 411-418.
<br>[3] Nugroho, A., & Purnomo, H. (2021). Evaluasi Kinerja Algoritma Perulangan dalam Pemrosesan Data Skala Besar. Jurnal Rekayasa Sistem dan Teknologi Informasi, 5(1), 112-119.
<br>[4] Hidayat, R. (2018). Implementasi Array dan Linked List pada Simulasi Penjadwalan Prosesor. Jurnal Edukasi dan Penelitian Informatika, 4(1), 32-38.
<br>[5] Kurniawan, B., & Setiawan, D. (2022). Pemodelan Struktur Data Record dalam Pengembangan Sistem Informasi Akademik Berbasis C++. Jurnal Sistem Informasi Bisnis, 10(2), 145-152.
<br>[6] Pratama, E. A. (2017). Pendekatan Pemrograman Modular dalam Pengembangan Perangkat Lunak Skala Menengah. Jurnal Informatika Terpadu, 3(2), 89-96.
