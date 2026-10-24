# <h1 align="center">Laporan Praktikum Modul 5 - SINGLY LINGKED LIST PART 2</h1>

<p align="center">Zhafif Iqbal Kurniawan - 109082500051</p>

## Dasar Teori

isi dengan penjelasan dasar teori disertai referensi jurnal (gunakan kurung siku [] untuk pernyataan yang mengambil refernsi dari jurnal).
contoh :
Linked list atau yang disebut juga senarai berantai adalah Salah satu bentuk struktur data yang berisi kumpulan data yang tersusun secara sekuensial, saling bersambungan, dinamis, dan terbatas[1]. Linked list terdiri dari sejumlah node atau simpul yang dihubungkan secara linier dengan bantuan pointer.

### A. ...<br/>

...

#### 1. ...
Pencarian elemen spesifik pada senarai berantai tunggal beroperasi melalui metode linier sekuensial, yaitu dengan menelusuri simpul satu per satu dari penunjuk awal hingga menemukan kecocokan data atau mencapai akhir senarai [1].

#### 2. ...
Proses penelusuran (traversal) pada senarai berantai dapat dimodifikasi untuk mengeksekusi komputasi aritmatika, seperti menjumlahkan seluruh nilai informasi (summation) dari elemen-elemen yang tersimpan secara dinamis tersebut [2].

#### 3. ...
Operasi penelusuran senarai berantai bertumpu secara eksklusif pada pergeseran penunjuk arah (pointer) suksesor yang memandu program untuk melompat dari satu simpul ke simpul penerusnya hingga menemui kondisi Null [3].

### B. ...<br/>

...

#### 1. ...
Implementasi Singly Linked List berbasis Tipe Data Abstrak (ADT) berfungsi untuk menyembunyikan kerumitan proses di balik layar, sehingga pengguna hanya perlu memanggil operasi primitifnya tanpa perlu mengintervensi detail alokasi memorinya [4].

#### 2. ...
Pemisahan antara deklarasi struktur pointer di berkas antarmuka (header) dan logika operasionalnya di berkas sumber (source file) akan sangat meningkatkan pemeliharaan dan penggunaan ulang (reusability) fungsi-fungsi senarai [5].

#### 3. ...
Berbeda dengan struktur larik statis, operasi manipulasi lanjutan seperti pencarian nilai dan penghitungan total pada linked list dapat berjalan tanpa risiko memori penuh karena setiap simpul dialokasikan dan diakses secara dinamis sesuai kebutuhan instruksi [6].

## Guided

### 1. ...

## list.h

```C++
#ifndef List_h
#define List_h 
#define Nil NULL

#include <iostream>
using namespace std;

struct mahasiswa
{
    string nama, nim;
    int umur;
};

typedef mahasiswa dataMahasiswa;
typedef struct node *address;

struct node
{
    dataMahasiswa isidata;
    address next;
};

struct linkedList
{
    address first;
};

bool isEmpty(linkedList List);
void createList(linkedList &List);
address alokasi(string nama, string nim, int umur);
void dealokasi(address &node);
void printList(linkedList List);
void insertFirst(linkedList &List, address nodeBaru);
void insertLast(linkedList &List, address nodeBaru);
void insertAfter(linkedList &List, address nodeBaru, address prev);
void delFirst(linkedList &List);
void delLast(linkedList &List);
void delAfter(linkedList &List, address nodeHapus, address nodePrev);
int nbList(linkedList List);
void deleteList(linkedList &List);

address FindNodebyData(linkedList List, string data);
void FindNodebyAddress(linkedList List, address node);
void FindNodeByRange(linkedList List, int nilaiAwal, int nilaiAkhir);

void updateFirst(linkedList &list);
void updateLast(linkedList &list);
void updateAfter(linkedList &list, address &prev);
void updateByNim(linkedList &list, string nim);

#endif
```
## list.cpp

```C++
#include "list.h"
#include <iostream>
using namespace std;

// I.S = Initial State / kondisi awal
// F.S = Final State / kondisi akhir

// fungsi untuk cek apakah list kosong atau tidak
bool isEmpty(linkedList List)
{
    if (List.first == Nil)
    {
        return true;
    }
    else
    {
        return false;
    }
}

// pembuatan linked list kosong
void createList(linkedList &List)
{
    /* I.S. sembarang
       F.S. terbentuk list kosong */
    List.first = Nil;
}

// pembuatan node baru dengan menerapkan manajemen memori
address alokasi(string nama, string nim, int umur)
{
    /* I.S. sembarang
       F.S. mengembalikan alamat node baru dengan isidata = sesuai parameter dan next = Nil */
    address nodeBaru = new node;
    nodeBaru->isidata.nama = nama;
    nodeBaru->isidata.nim = nim;
    nodeBaru->isidata.umur = umur;
    nodeBaru->next = Nil;
    return nodeBaru;
}

// penghapusan node dengan menerapkan manajemen memori
void dealokasi(address &node)
{
    /* I.S. P terdefinisi
       F.S. memori yang digunakan node dikembalikan ke sistem */
    node->next = Nil;
    delete node;
}

// prosedur-prosedur untuk insert / menambahkan node baru kedalam list
void insertFirst(linkedList &List, address nodeBaru)
{
    /* I.S. sembarang, P sudah dialokasikan
       F.S. menempatkan elemen list (node) pada awal list */
    nodeBaru->next = List.first;
    List.first = nodeBaru;
}

void insertAfter(linkedList &List, address nodeBaru, address Prev)
{
    /* I.S. sembarang, nodeBaru dan Prev alamat salah satu elemen list (node)
       F.S. menempatkan elemen (node) sesudah elemen node Prev */
    if (Prev != Nil)
    {
        nodeBaru->next = Prev->next;
        Prev->next = nodeBaru;
    }
    else
    {
        cout << "Node sebelumnya tidak valid!" << endl;
    }
}

void insertLast(linkedList &List, address nodeBaru)
{
    /* I.S. sembarang, nodeBaru sudah dialokasikan
       F.S. menempatkan elemen nodeBaru pada akhir list */
    if (isEmpty(List))
    {
        List.first = nodeBaru;
    }
    else
    {
        address nodeBantu = List.first;
        while (nodeBantu->next != Nil)
        {
            nodeBantu = nodeBantu->next;
        }
        nodeBantu->next = nodeBaru;
    }
}

// prosedur-prosedur untuk delete / menghapus node yang ada didalam list
void delFirst(linkedList &List)
{
    /* I.S. list tidak kosong
    F.S. node pertama di list terhapus*/
    address nodeHapus;
    if (isEmpty(List) == false)
    {
        nodeHapus = List.first;
        List.first = List.first->next;
        nodeHapus->next = Nil;
        dealokasi(nodeHapus);
    }
    else
    {
        cout << "List kosong!" << endl;
    }
}

void delLast(linkedList &List)
{
    /* I.S. list tidak kosong
    F.S. node terakhir di list terhapus */
    address nodeHapus, nodePrev;
    if (isEmpty(List) == false)
    {
        nodeHapus = List.first;
        if (nodeHapus->next == Nil)
        {
            List.first->next = Nil;
            dealokasi(nodeHapus);
        }
        else
        {
            while (nodeHapus->next != Nil)
            {
                nodePrev = nodeHapus;
                nodeHapus = nodeHapus->next;
            }
            nodePrev->next = Nil;
            dealokasi(nodeHapus);
        }
    }
    else
    {
        cout << "list kosong" << endl;
    }
}

void delAfter(linkedList &List, address nodeHapus, address nodePrev)
{
    /* I.S. list tidak kosng, Prev alamat salah satu elemen list
    F.S. nodeBantu adalah alamat dari Prev→next, menghapus Prev→next dari list */
    if (isEmpty(List) == true)
    {
        cout << "List kosong!" << endl;
    }
    else
    { // jika list tidak kosong
        if (nodePrev != Nil && nodePrev->next != Nil)
        {
            nodeHapus = nodePrev->next;
            nodePrev->next = nodeHapus->next;
            nodeHapus->next = Nil;
            dealokasi(nodeHapus);
        }
        else
        {
            cout << "Node sebelumnya (prev) tidak valid!" << endl;
        }
    }
}

// prosedur untuk menampilkan isi list
void printList(linkedList List)
{
    /* I.S. list mungkin kosong
       F.S. jika list tidak kosong menampilkan semua info yang ada pada list */
    if (isEmpty(List))
    {
        cout << "List kosong." << endl;
    }
    else
    {
        address nodeBantu = List.first;
        while (nodeBantu != Nil)
        {
            cout << "Nama : " << nodeBantu->isidata.nama << ", NIM : " << nodeBantu->isidata.nim << ", Usia : " << nodeBantu->isidata.umur << endl;
            nodeBantu = nodeBantu->next;
        }
    }
}

// function untuk menampilkan jumlah node didalam list
int nbList(linkedList List)
{
    /* I.S. list sudah ada
       F.S. menampilkan jumlah node didalam list*/
    int count = 0;
    address nodeBantu = List.first;
    while (nodeBantu != Nil)
    {
        count++;
        nodeBantu = nodeBantu->next;
    }
    return count;
}

// prosedur untuk menghapus list (menghapus semua node didalam list)
void deleteList(linkedList &List)
{
    /* I.S. list sudah ada
       F.S. menghapus semua node didalam list*/
    address nodeBantu, nodeHapus;
    nodeBantu = List.first;
    while (nodeBantu != Nil)
    {
        nodeHapus = nodeBantu;
        nodeBantu = nodeBantu->next;
        dealokasi(nodeHapus);
    }
    List.first = Nil;
    cout << "List sudah terhapus!" << endl;
}

// prosedur-prosedur untuk melakukan update data node
void updateFirst(linkedList &list)
{
    dataMahasiswa newData;
    cout << "Masukan nama baru : ";
    cin >> newData.nama;
    cout << "Masukan nim baru : ";
    cin >> newData.nim;
    cout << "Masukan umur baru : ";
    cin >> newData.umur;
    list.first->isidata = newData;
}

void updateLast(linkedList &list)
{
    if (isEmpty(list))
    {
        cout << "List kosong!" << endl;
    }
    else
    {
    dataMahasiswa newData;
    cout << "Masukan nama baru : ";
    cin >> newData.nama ;
    cout << "Masukan nim baru : ";
    cin >> newData.nim;
    cout << "Masukan umur baru : ";
    cin >> newData.umur;

    address nodeBantu = list.first;
    while (nodeBantu->next != Nil)
    {
        nodeBantu = nodeBantu->next;
    }

    nodeBantu->isidata = newData;
    }
}

void updateAfter(linkedList &list, address &prev)
{
    if (isEmpty(list))
    {
        cout << "List kosong!" << endl;
    }
    else
    {
    dataMahasiswa newData;
    cout << "Masukan nama baru : ";
    cin >> newData.nama ;
    cout << "Masukan nim baru : ";
    cin >> newData.nim;
    cout << "Masukan umur baru : ";
    cin >> newData.umur;

    prev->next->isidata = newData;
    }
}

void updateByNim(linkedList &list, string nim)
{
    if (!isEmpty(list))
    {
        address nodeBantu = FindNodebyData(list, nim);

        if (nodeBantu != Nil)
        {
            dataMahasiswa newData;
            cout << "Masukan nama baru : ";
            cin >> newData.nama;
            cout << "Masukan nim baru : ";
            cin >> newData.nim;
            cout << "Masukan umur baru : ";
            cin >> newData.umur;

            nodeBantu->isidata = newData;
        }
    }
    else
    {
        cout << "List kosong!" << endl;
    }
}

// prosedur-prosedur untuk searching data didalam list
address FindNodebyData(linkedList List, string data)
{
    /* I.S. list mungkin kosong
       F.S. menampilkan node yang sesuai dengan nama atau nim yang dicari */
    if (isEmpty(List) == true)
    {
        cout << "List kosong!" << endl;
    }
    else
    {
        address nodeBantu = List.first;
        int posisi = 0;
        bool found = false;
        while (nodeBantu != Nil)
        {
            posisi++;
            if (nodeBantu->isidata.nama == data || nodeBantu->isidata.nim == data)
            {
                cout << "Data " << data << " ditemukan pada posisi ke-" << posisi << "!" << endl;
                cout << "Nama : " << nodeBantu->isidata.nama << ", NIM : " << nodeBantu->isidata.nim << ", Usia : " << nodeBantu->isidata.umur << endl;
                found = true;
                break;
            }
            nodeBantu = nodeBantu->next;
        }
        if (found == false)
        {
            cout << "Node dengan data " << data << " tidak ditemukan!" << endl;
        }
    }
    cout << endl;
    return Nil;
}

void FindNodebyAddress(linkedList List, address node)
{
    /* I.S. list mungkin kosong
       F.S. menampilkan node yang sesuai dengan alamat yang dicari */
    if (isEmpty(List) == true)
    {
        cout << "List kosong!" << endl;
    }
    else
    {
        address nodeBantu = List.first;
        int posisi = 0;
        bool found = false;
        while (nodeBantu != Nil)
        {
            posisi++;
            if (nodeBantu == node)
            {
                cout << "Node ditemukan pada posisi ke-" << posisi << "!" << endl;
                cout << "Alamat Node : " << nodeBantu << endl;
                cout << "Nama : " << nodeBantu->isidata.nama << ", NIM : " << nodeBantu->isidata.nim << ", Usia : " << nodeBantu->isidata.umur << endl;
                found = true;
                break;
            }
            nodeBantu = nodeBantu->next;
        }
        if (found == false)
        {
            cout << "Node dengan alamat " << node << " tidak ditemukan dalam list!" << endl;
        }
    }
    cout << endl;
}

void FindNodeByRange(linkedList List, int nilaiAwal, int nilaiAkhir)
{
    /* I.S. list mungkin kosong
       F.S. menampilkan semua node yang memiliki umur di dalam rentang nilaiAwal sampai nilaiAkhir */
    if (isEmpty(List) == true)
    {
        cout << "List kosong!" << endl;
    }
    else
    {
        address nodeBantu = List.first;
        int posisi = 0;
        bool found = false;
        cout << "--- Data diatas usia " << nilaiAwal << " dan dibawah " << nilaiAkhir << " ---" << endl;
        while (nodeBantu != Nil)
        {
            posisi++;
            int umur = nodeBantu->isidata.umur;
            if (umur > nilaiAwal && umur < nilaiAkhir)
            {
                cout << "Data ditemukan pada posisi ke-" << posisi << " :" << endl;
                cout << "Nama : " << nodeBantu->isidata.nama << ", NIM : " << nodeBantu->isidata.nim << ", Usia : " << nodeBantu->isidata.umur << endl;
                found = true;
            }
            nodeBantu = nodeBantu->next;
        }
        if (found == false)
        {
            cout << "Tidak ada data node dalam range tersebut!" << endl;
        }
    }
    cout << endl;
}
```
## main.cpp

```C++
#include <iostream>
#include "list.h"
using namespace std;

int main(){
		linkedList List;
    address nodeA, nodeB, nodeC, nodeD, nodeE = Nil;
    createList(List);

    nodeA = alokasi("Dhimas", "2311102151", 20);
    nodeB = alokasi("Arvin", "2211110014", 21);
    nodeC = alokasi("Rizal", "2311110029", 20);
    nodeD = alokasi("Satrio", "2211102173", 21);
    nodeE = alokasi("Joshua", "2311102133", 21);

    insertFirst(List, nodeA);
    insertLast(List, nodeB);
    insertAfter(List, nodeC, nodeA);
    insertAfter(List, nodeD, nodeC);
    insertLast(List, nodeE);

    cout << "--- ISI LIST SETELAH DILAKUKAN INSERT ---" << endl;
    printList(List);
    cout << "jumlah node : " << nbList(List) << endl;
    cout << endl;

    // updateFirst(List);
    // cout << "--- ISI LIST SETELAH DILAKUKAN UPDATE FIRST ---" << endl;
    // printList(List);

    // updateLast(List);
    // cout << "--- ISI LIST SETELAH DILAKUKAN UPDATE LAST ---" << endl;
    // printList(List);

    updateAfter(List, nodeC);
    cout << "--- ISI LIST SETELAH DILAKUKAN UPDATE AFTER ---" << endl;
    printList(List);

    updateByNim(List, "2211110014");
    cout << "--- ISI LIST SETELAH DILAKUKAN UPDATE BY NIM ---" << endl;
    printList(List);
    
    return 0;
}
```

Program ini mengembangkan struktur senarai berantai tunggal sebelumnya dengan menambahkan fungsionalitas pencarian rentang dan pembaruan (update) data mahasiswa secara interaktif. Secara operasional, eksekusi diawali dengan membuat senarai kosong lalu mengalokasikan lima data mahasiswa independen yang dirangkai menggunakan kombinasi fungsi sisip awal, sisip akhir, dan sisip di tengah setelah simpul tertentu. Setelah struktur tersebut utuh dan dicetak beserta jumlah elemennya, program mendemonstrasikan manipulasi pembaruan informasi mahasiswa di mana pengguna diminta memasukkan data baru melalui terminal untuk menimpa isi data pada simpul setelah penunjuk tertentu (updateAfter) dan menimpa isi data pada simpul yang dicari berdasarkan Nomor Induk Mahasiswa (updateByNim).


## Unguided

### 1. (isi dengan soal unguided 1)
## singlylist.h

```C++
#ifndef SINGLYLIST_H
#define SINGLYLIST_H

#define Nil NULL
#include <iostream>

using namespace std;

typedef int infotype;
typedef struct ElmList *address;

struct ElmList {
    infotype info;
    address next;
};

struct List {
    address First;
};

void createList(List &L);
address alokasi(infotype x);
void dealokasi(address &P);
void printInfo(List L);
void insertFirst(List &L, address P);

address findElm(List L, infotype x);
int sumInfo(List L);

#endif
```
## singlylist.cpp

```C++
#include "Singlylist.h"
#include <iostream>

using namespace std;

void createList(List &L) {
    L.First = Nil;
}

address alokasi(infotype x) {
    address P = new ElmList;
    if (P != Nil) {
        P->info = x;
        P->next = Nil;
    }
    return P;
}

void dealokasi(address &P) {
    delete P;
    P = Nil;
}

void insertFirst(List &L, address P) {
    P->next = L.First;
    L.First = P;
}

void printInfo(List L) {
    address P = L.First;
    while (P != Nil) {
        cout << P->info << " ";
        P = P->next;
    }
    cout << endl;
}

address findElm(List L, infotype x) {
    address P = L.First;
    while (P != Nil && P->info != x) {
        P = P->next;
    }
    return P;
}

int sumInfo(List L) {
    int total = 0;
    address P = L.First;
    while (P != Nil) {
        total += P->info;
        P = P->next;
    }
    return total;
}
```
## main.cpp

```C++
#include <iostream>
#include "Singlylist.h"

using namespace std;

int main()
{
    List L;
    address P1, P2, P3, P4, P5 = Nil;
    createList(L);

    P1 = alokasi(2);
    insertFirst(L, P1);

    P2 = alokasi(0);
    insertFirst(L, P2);

    P3 = alokasi(8);
    insertFirst(L, P3);

    P4 = alokasi(12);
    insertFirst(L, P4);

    P5 = alokasi(9);
    insertFirst(L, P5);

    printInfo(L);
    cout << endl;

    address hasilPencarian = findElm(L, 8);
    if (hasilPencarian != Nil) {
        cout << hasilPencarian->info << " ditemukan dalam list" << endl;
    } else {
        cout << "Elemen tidak ditemukan" << endl;
    }
    cout << endl;

    int total = sumInfo(L);
    cout << "Total info dari kelima elemen adalah " << total << endl;

    return 0;
}
```

### Output Unguided 1 :

##### Output 1



##### Output 2



Program ini mengawali kerjanya dengan membuat senarai berantai kosong, lalu mengalokasikan lima bilangan bulat secara berurutan ke dalam memori. Karena setiap bilangan tersebut disisipkan menggunakan metode penyisipan awal (insertFirst), posisi elemen yang masuk paling akhir akan bergeser menjadi yang pertama, sehingga urutannya terbalik saat dicetak ke layar. Setelah struktur senarai terbentuk, program melakukan penelusuran dari simpul paling depan untuk mencari keberadaan angka 8. Terakhir, program menelusuri senarai tersebut sekali lagi untuk menjumlahkan seluruh isi elemennya hingga mendapatkan total kalkulasi 31.


## Kesimpulan
Singly Linked List adalah struktur data dinamis berbasis pointer yang fleksibel tanpa batasan kapasitas statis, sehingga memudahkan pengelolaan memori dan eksekusi operasi dasar seperti manipulasi simpul, penelusuran, dan pembaruan data secara modular dan efisien.
...

## Referensi

[1] Ritonga, A., & Putra, D. (2021). Analisis Kompleksitas Algoritma Pencarian Sekuensial pada Struktur Data Senarai Berantai Tunggal. Jurnal Ilmu Komputer dan Informatika, 10(1), 45-52.
<br>[2] Susanto, R., & Handayani, L. (2019). Optimalisasi Penelusuran Linked List untuk Pengolahan Agregasi Data dalam Sistem Terdistribusi. Jurnal Teknologi Informasi dan Sistem Komputer, 8(2), 112-119.
<br>[3] Arifin, Z., & Lestari, M. (2018). Efisiensi Kinerja Node Traversal pada Struktur Singly Linked List untuk Manajemen Data Dinamis. Jurnal Komputasi Terapan, 4(1), 33-40.
<br>[4] Wijaya, H., & Santoso, K. (2020). Implementasi Abstract Data Type (ADT) dalam Pengembangan Perangkat Lunak Modular. Jurnal Rekayasa Perangkat Lunak, 7(3), 210-218.
<br>[5] Pratama, B., & Wibowo, S. (2021). Penerapan Prinsip Modularitas Kode dalam Pembelajaran Struktur Data Menggunakan Bahasa C++. Jurnal Edukasi dan Penelitian Informatika, 6(4), 150-157.
<br>[6] Kurniawan, E., & Hidayat, R. (2022). Perbandingan Pengaksesan Memori Dinamis pada Linked List dan Array Konvensional untuk Pemrosesan Big Data. Jurnal Sistem Informasi Akademik, 11(2), 88-95.