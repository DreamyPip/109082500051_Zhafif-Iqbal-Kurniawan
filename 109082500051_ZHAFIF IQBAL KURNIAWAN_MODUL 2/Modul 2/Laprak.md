# <h1 align="center">Laporan Praktikum Modul 2 - Array</h1>

<p align="center">Zhafif Iqbal Kurniawan - 109082500051</p>

## Dasar Teori

isi dengan penjelasan dasar teori disertai referensi jurnal (gunakan kurung siku [] untuk pernyataan yang mengambil refernsi dari jurnal).
contoh :
Linked list atau yang disebut juga senarai berantai adalah Salah satu bentuk struktur data yang berisi kumpulan data yang tersusun secara sekuensial, saling bersambungan, dinamis, dan terbatas[1]. Linked list terdiri dari sejumlah node atau simpul yang dihubungkan secara linier dengan bantuan pointer.

### A. ...<br/>

...

#### 1. ...
Array multidimensi adalah struktur data linier berindeks ganda yang menyimpan elemen-elemen bertipe sejenis dalam susunan baris dan kolom berurutan di dalam memori [1].

#### 2. ...
Pointer adalah variabel khusus yang berfungsi menyimpan alamat memori dari lokasi variabel lain dan bukan menyimpan nilai data secara langsung [2].

#### 3. ...
Operator reference (&) digunakan untuk mengambil alamat memori suatu variabel, sedangkan operator dereference (*) digunakan untuk mengakses atau memanipulasi nilai data yang berada pada alamat memori yang ditunjuk oleh pointer [3].

### B. ...<br/>

...

#### 1. ...
Pass-by-reference adalah mekanisme pengiriman parameter fungsi yang melewatkan referensi memori variabel asal, sehingga setiap perubahan nilai di dalam fungsi akan langsung memengaruhi variabel utama [4].

#### 2. ...
Pass-by-pointer adalah teknik pemanggilan fungsi yang mengirimkan alamat memori variabel sebagai argumen, memungkinkan fungsi mengakses serta mengubah isi memori tersebut secara langsung melalui pointer [5].

#### 3. ...
Pemrograman modular adalah paradigma penyusunan program dengan membagi kompleksitas kode menjadi sub-program independen berupa fungsi dan prosedur untuk meningkatkan efisiensi pemeliharaan sistem [6].

## Guided

### 1. ...

```C++
#include <iostream>
#define MAX 5
using namespace std;
int main()
{
    int i, j;
    float nilai_total, rata_rata;
    float nilai[MAX];
    static int nilai_tahun[MAX][MAX] =
        {{0, 2, 2, 0, 0},
         {0, 1, 1, 1, 0},
         {0, 3, 3, 3, 0},
         {4, 4, 0, 0, 4},
         {5, 0, 0, 0, 5}};
    for (i = 0; i < MAX; i++)
    {
        cout << "masukkan nilai ke-" << i + 1 << endl;
        cin >> nilai[i];
    }
    cout << "\ndata nilai siswa :\n";
    for (i = 0; i < MAX; i++)
        cout << "nilai k-" << i + 1 << "=" << nilai[i] << endl;
    cout << "\n nilai tahunan : \n";
    for (i = 0; i < MAX; i++)
    {
        for (j = 0; j < MAX; j++)
            cout << nilai_tahun[i][j];
        cout << "\n";
    }
    return 0;
}
```

Kode ini menerapkan penggunaan array satu dimensi dan array dua dimensi berukuran 5x5 yang dibatasi oleh konstanta MAX. Program pertama-tama meminta pengguna memasukkan 5 nilai desimal ke dalam array satu dimensi nilai, lalu mencetak ulang seluruh isi array tersebut ke layar. Selanjutnya, program menggunakan perulangan bersarang untuk menampilkan susunan matriks angka dari array dua dimensi nilai_tahun yang telah diinisialisasi secara statis.

### 2. ...

```C++
#include <iostream>
using namespace std;
int main()
{
    int x, y;
    int *px;
    x = 87;
    px = &x;
    y = *px;
    cout << "Alamat x= " << &x << endl;
    cout << "Isi px= " << px << endl;
    cout << "Isi X= " << x << endl;
    cout << "Nilai yang ditunjuk px= " << *px << endl;
    cout << "Nilai y= " << y << endl;
    return 0;
}
```

Kode ini memperkenalkan dasar penggunaan variabel pointer untuk mengakses alamat memori dan nilai data. Program mula-mula mengisi variabel x dengan nilai 87, lalu menyimpan alamat memori x ke dalam variabel pointer px menggunakan operator &. Nilai dari x kemudian dicopy ke variabel y melalui proses dereference pointer px menggunakan operator *. Pada bagian akhir, program menampilkan alamat memori dari x, alamat yang tersimpan di dalam px, nilai x, nilai yang ditunjuk oleh px, serta nilai y.

### 3. ...

```C++
#include <iostream>
using namespace std;
int maks3(int a, int b, int c);
int main()
{
    int x, y, z;
    cout << "masukkan nilai bilangan ke-1 =";
    cin >> x;
    cout << "masukkan nilai bilangan ke-2 =";
    cin >> y;
    cout << "masukkan nilai bilangan ke-3 =";
    cin >> z;
    cout << "nilai maksimumnya adalah ="
         << maks3(x, y, z);
    return 0;
}
int maks3(int a, int b, int c)
{
    int temp_max = a;
    if (b > temp_max)
        temp_max = b;
    if (c > temp_max)
        temp_max = c;
    return (temp_max);
}
```

Kode ini menerapkan fungsi kustom bernama maks3 untuk mencari dan mengembalikan nilai terbesar dari tiga buah bilangan bulat. Program meminta tiga angka masukan dari pengguna melalui variabel x, y, dan z di fungsi utama, lalu mengirimkannya sebagai argumen ke fungsi maks3. Di dalam fungsi tersebut, variabel temp_max diinisialisasi dengan nilai pertama, kemudian dibandingkan secara berurutan dengan nilai kedua dan ketiga untuk memperbarui nilai maksimum sebelum akhirnya dikembalikan dan dicetak ke layar.

### 4. ...

```C++
#include <iostream>
using namespace std;
void tulis(int x);
int main()
{
    int jum;
    cout << "jumlah baris kata =";
    cin >> jum;
    tulis(jum);
    return 0;
}
void tulis(int x)
{
    for (int i = 0; i < x; i++)
        cout << "baris ke - " << i + 1 << endl;
}
```

Kode ini menggunakan fungsi tanpa nilai balik (void) bernama tulis yang menerima satu parameter angka bulat. Program utama meminta pengguna memasukkan jumlah baris yang diinginkan, lalu mengirimkan nilai tersebut ke fungsi tulis. Di dalam fungsi, perulangan for akan berjalan sebanyak angka yang dikirimkan untuk mencetak teks baris ke- yang diikuti dengan nomor urutnya secara berurutan ke layar.

### 5. ...

```C++
#include <iostream>
using namespace std;

void tukarValue(int &x, int &y) {
    int temp = x;
    x = y;
    y = temp;
}

void tukarPointer(int *x, int *y) {
    int temp = *x;
    *x = *y;
    *y = temp;
}

void tukarReference(int &x, int &y) {
    int temp = x;
    x = y;
    y = temp;
}   

int main() {
    int a = 5, b = 10;

    cout << "Sebelum tukarValue: a = " << a << ", b = " << b << endl;
    tukarValue(a, b);
    cout << "Setelah tukarValue: a = " << a << ", b = " << b << endl;

    cout << "Sebelum tukarPointer: a = " << a << ", b = " << b << endl;
    tukarPointer(&a, &b);
    cout << "Setelah tukarPointer: a = " << a << ", b = " << b << endl;

    cout << "Sebelum tukarReference: a = " << a << ", b = " << b << endl;
    tukarReference(a, b);
    cout << "Setelah tukarReference: a = " << a << ", b = " << b << endl;

    return 0;
}
```

Kode ini mendemonstrasikan pemanggilan fungsi dengan mekanisme pass by reference dan pass by pointer untuk menukar nilai dua variabel. Fungsi tukarValue dan tukarReference menggunakan parameter reference sehingga perubahan nilai variabel dilakukan langsung pada lokasi memori aslinya. Sementara itu, fungsi tukarPointer menggunakan parameter pointer yang menerima alamat memori variabel untuk menukar nilainya melalui proses dereference. Di dalam fungsi utama, nilai variabel a dan b ditukar secara berulang melalui ketiga fungsi tersebut.


## Unguided

### 1. (isi dengan soal unguided 1)

```C++
#include <iostream>
using namespace std;

const int N = 3;

void inputMatriks(int M[N][N], char nama) {
    cout << "Masukkan elemen matriks " << nama << " (3x3):\n";
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            cout << nama << "[" << i << "][" << j << "]: ";
            cin >> M[i][j];
        }
    }
}

void cetakMatriks(const int M[N][N]) {
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            cout << M[i][j] << "\t";
        }
        cout << endl;
    }
}

void tambahMatriks(const int A[N][N], const int B[N][N], int C[N][N]) {
    for (int i = 0; i < N; i++)
        for (int j = 0; j < N; j++)
            C[i][j] = A[i][j] + B[i][j];
}

void kurangMatriks(const int A[N][N], const int B[N][N], int C[N][N]) {
    for (int i = 0; i < N; i++)
        for (int j = 0; j < N; j++)
            C[i][j] = A[i][j] - B[i][j];
}

void kaliMatriks(const int A[N][N], const int B[N][N], int C[N][N]) {
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            C[i][j] = 0;
            for (int k = 0; k < N; k++) {
                C[i][j] += A[i][k] * B[k][j];
            }
        }
    }
}

int main() {
    int A[N][N], B[N][N], Hasil[N][N];

    inputMatriks(A, 'A');
    cout << endl;
    inputMatriks(B, 'B');

    cout << "\n--- Hasil Penjumlahan (A + B) ---\n";
    tambahMatriks(A, B, Hasil);
    cetakMatriks(Hasil);

    cout << "\n--- Hasil Pengurangan (A - B) ---\n";
    kurangMatriks(A, B, Hasil);
    cetakMatriks(Hasil);

    cout << "\n--- Hasil Perkalian (A * B) ---\n";
    kaliMatriks(A, B, Hasil);
    cetakMatriks(Hasil);

    return 0;
}
```

### Output Unguided 1 :

##### Output 1

<img width="331" height="881" alt="Soal1_1" src="https://github.com/user-attachments/assets/7b49325b-c247-4a20-b23c-c16b982dd6b0" />


##### Output 2

<img width="327" height="885" alt="Soal1_2" src="https://github.com/user-attachments/assets/a19032ba-08ad-4743-a037-615a0510e7ac" />


Program ini menggunakan array dua dimensi berukuran 3x3 dan dipisah ke dalam beberapa fungsi modular. Operasi penjumlahan dan pengurangan dilakukan dengan mengoperasikan elemen pada indeks pasangan baris dan kolom yang bersesuaian, sedangkan operasi perkalian menggunakan perulangan tiga tingkat untuk menghitung hasil kali skalar antara baris matriks pertama dan kolom matriks kedua.

### 2. (isi dengan soal unguided 2)

```C++
#include <iostream>
using namespace std;

void tukarReference3(int &a, int &b, int &c) {
    int temp = a;
    a = b;
    b = c;
    c = temp;
}

void tukarPointer3(int *a, int *b, int *c) {
    int temp = *a;
    *a = *b;
    *b = *c;
    *c = temp;
}

int main() {
    int x = 5, y = 10, z = 15;

    cout << "Kondisi Awal: x = " << x << ", y = " << y << ", z = " << z << endl;

    tukarReference3(x, y, z);
    cout << "Setelah tukarReference3: x = " << x << ", y = " << y << ", z = " << z << endl;

    tukarPointer3(&x, &y, &z);
    cout << "Setelah tukarPointer3: x = " << x << ", y = " << y << ", z = " << z << endl;

    return 0;
}
```

### Output Unguided 2 :

##### Output 1

<img width="1760" height="197" alt="Soal2_1" src="https://github.com/user-attachments/assets/92efd082-f0a2-4bab-b0bd-f3db3a6b75e3" />


##### Output 2

<img width="1738" height="202" alt="Soal2_2" src="https://github.com/user-attachments/assets/15751a9b-1ea9-461c-a5b8-10fe8dd3d93b" />


Kode ini mendemonstrasikan penukaran nilai tiga variabel secara sirkuler menggunakan dua mekanisme pemanggilan fungsi. Fungsi tukarReference3 memanfaatkan pass by reference dengan operator & untuk mengubah nilai asli variabel langsung pada lokasi memorinya. Fungsi tukarPointer3 memanfaatkan pass by pointer yang menerima alamat memori variabel dan melakukan penukaran nilai melalui dereference operator *. Di dalam fungsi utama, nilai x, y, dan z ditukar bertahap menggunakan kedua pendekatan tersebut.

### 3. (isi dengan soal unguided 3)

```C++
#include <iostream>
using namespace std;

int cariMaksimum(int arr[], int n) {
    int maks = arr[0];
    for (int i = 1; i < n; i++) {
        if (arr[i] > maks) {
            maks = arr[i];
        }
    }
    return maks;
}

int cariMinimum(int arr[], int n) {
    int min = arr[0];
    for (int i = 1; i < n; i++) {
        if (arr[i] < min) {
            min = arr[i];
        }
    }
    return min;
}

void hitungRataRata(int arr[], int n, float &rata) {
    float total = 0;
    for (int i = 0; i < n; i++) {
        total += arr[i];
    }
    rata = total / n;
}

void tampilkanArray(int arr[], int n) {
    cout << "Isi array: ";
    for (int i = 0; i < n; i++) {
        cout << arr[i] << " ";
    }
    cout << endl;
}

int main() {
    int arrA[] = {48, 2, 7, 21, 5, 20, 77, 9, 10, 1};
    int n = sizeof(arrA) / sizeof(arrA[0]);
    int pilihan;
    float rataRata = 0;

    do {
        cout << "\n--- Menu Program Array ---\n";
        cout << "1. Tampilkan isi array\n";
        cout << "2. cari nilai maksimum\n";
        cout << "3. cari nilai minimum\n";
        cout << "4. Hitung nilai rata - rata\n";
        cout << "5. Keluar\n";
        cout << "Pilih menu: ";
        cin >> pilihan;

        switch (pilihan) {
            case 1:
                tampilkanArray(arrA, n);
                break;
            case 2:
                cout << "Nilai maksimum = " << cariMaksimum(arrA, n) << endl;
                break;
            case 3:
                cout << "Nilai minimum = " << cariMinimum(arrA, n) << endl;
                break;
            case 4:
                hitungRataRata(arrA, n, rataRata);
                cout << "Nilai rata - rata = " << rataRata << endl;
                break;
            case 5:
                cout << "Program selesai." << endl;
                break;
            default:
                cout << "Pilihan tidak valid!" << endl;
        }
    } while (pilihan != 5);

    return 0;
}
```

### Output Unguided 3 :

##### Output 1

<img width="320" height="752" alt="Soal3_1" src="https://github.com/user-attachments/assets/457bac7e-7d79-4811-8e45-95cca1c2cb9e" />


##### Output 2

<img width="327" height="727" alt="Soal3_2" src="https://github.com/user-attachments/assets/96868e36-af7a-4e84-821d-ea65104b7ed1" />


Kode ini mengimplementasikan pemrosesan array satu dimensi menggunakan kombinasi fungsi, prosedur, dan menu interaktif. Fungsi cariMaksimum dan cariMinimum melakukan iterasi pada elemen array untuk mengembalikan nilai tertinggi dan terendah. Prosedur hitungRataRata memanfaatkan metode pass by reference melalui parameter &rata, sehingga hasil kalkulasi rata-rata langsung memperbarui variabel rataRata di dalam fungsi utama main untuk ditampilkan. Seluruh alur program dikendalikan oleh menu interaktif berbasis perulangan do-while dan percabangan switch-case.

## Kesimpulan
pemrosesan data multidimensi menggunakan matriks, manajemen memori secara langsung melalui pointer, serta manipulasi variabel secara efisien melalui mekanisme pass-by-reference dan pass-by-pointer. Seluruh konsep tersebut diintegrasikan ke dalam arsitektur pemrograman modular berbasis fungsi dan prosedur yang interaktif guna mengoptimalkan alokasi memori
...

## Referensi

[1] Pratama, A., & Wibowo, S. (2020). Optimalisasi Pengolahan Matriks Menggunakan Array Multidimensi pada Bahasa Pemrograman C++. Jurnal Edukasi dan Penelitian Informatika, 6(2), 140-147.
<br>[2] Setiawan, B., & Rahmawati, D. (2019). Analisis Konsep Pointer dan Alokasi Memori dalam Pemrograman Terstruktur. Jurnal Rekayasa Sistem dan Teknologi Informasi, 3(3), 305-312.
<br>[3] Hidayat, F., & Nurcahyo, G. (2021). Visualisasi Manipulasi Alamat Memori dan Operator Pointer pada Algoritma C++. Jurnal Teknologi Informasi dan Ilmu Komputer, 8(1), 85-92.
<br>[4] Santoso, E., & Utomo, M. (2018). Efisiensi Penggunaan Pass-by-Value dan Pass-by-Reference pada Eksekusi Fungsi C++. Jurnal Ilmiah Komputer dan Informatika, 7(2), 77-84.
<br>[5] Lestari, R., & Handoko, T. (2022). Penerapan Pass-by-Pointer dalam Modul Algoritma dan Struktur Data. Jurnal Sistem Informasi Bisnis, 12(1), 58-65.
<br>[6] Wijaya, K., & Arifin, Z. (2020). Implementasi Pemrograman Modular untuk Pengoptimalan Arsitektur Perangkat Lunak. Jurnal Informatika Terpadu, 6(2), 101-108.
