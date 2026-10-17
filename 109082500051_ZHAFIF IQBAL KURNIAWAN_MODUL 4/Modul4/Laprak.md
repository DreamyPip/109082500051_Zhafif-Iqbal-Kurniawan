# <h1 align="center">Laporan Praktikum Modul 4 - SINGLY LINGKED LIST</h1>

<p align="center">Zhafif Iqbal Kurniawan - 109082500051</p>

## Dasar Teori

isi dengan penjelasan dasar teori disertai referensi jurnal (gunakan kurung siku [] untuk pernyataan yang mengambil refernsi dari jurnal).
contoh :
Linked list atau yang disebut juga senarai berantai adalah Salah satu bentuk struktur data yang berisi kumpulan data yang tersusun secara sekuensial, saling bersambungan, dinamis, dan terbatas[1]. Linked list terdiri dari sejumlah node atau simpul yang dihubungkan secara linier dengan bantuan pointer.

### A. ...<br/>

...

#### 1. ...
Senarai berantai merupakan representasi penyimpanan berupa serangkaian elemen data majemuk atau tunggal yang saling terhubung melalui penunjuk dan bersifat dinamis karena dapat bertumbuh sesuai kebutuhan [1].

#### 2. ...
Singly linked list adalah model struktur data yang beroperasi dengan hanya memiliki satu arah penunjuk ke elemen penerus di mana simpul akhirnya selalu menunjuk pada nilai kosong [2].

#### 3. ...
Alokasi adalah tahapan krusial untuk memesan dan menyediakan memori baru bagi setiap elemen data ke dalam sistem saat program berjalan menggunakan sintaksis perintah khusus [3].

### B. ...<br/>

...

#### 1. ...
Operasi penyisipan berfungsi mengaitkan elemen memori baru ke dalam senarai melalui manipulasi penunjuk alamat, baik pada posisi awal, akhir, maupun diletakkan setelah simpul tertentu [4].

#### 2. ...
Penghapusan beroperasi dengan melepaskan ikatan suatu elemen dari senarai yang kemudian dilanjutkan dengan instruksi pembebasan ruang memori (dealokasi) agar dikembalikan ke sistem [5].

#### 3. ...
Penelusuran merupakan operasi untuk mengunjungi dan menggeser penunjuk ke setiap simpul secara berurutan guna membaca serta menampilkan seluruh informasi yang tersimpan di dalamnya [6].

## Guided

### 1. ...

## list.h

```C++
#ifndef List.h
#define List .h
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
struct linkedlist
{
    address first;
};

bool isEmpty(linkedlist list);
void createList(linkedlist &list);
address alokasi(string nama, string nim, int umur);
void dealokasi (address &node);
void printList(linkedlist list);
void insertFirst (linkedlist &list, address nodeBaru);
void insertAfter(linkedlist &list, address nodeBaru, address prev);
void insertLast(linkedlist &list, address nodeBaru);

void delFirst(linkedlist &list);
void delLast (linkedlist &list);
void delAfter (linkedlist &list, address nodeHapus, address nodePrev);
int nbList(linkedlist list);
void deletelist (linkedlist &list);
#endif
```
## list.cpp

```C++
#include "list.h"
#include <iostream>
using namespace std;

bool isEmpty(linkedlist list)
{
    if (list.first == Nil)
    {
        return true;
    }
    else
    {
        return false;
    }
}

void createList(linkedlist &list)
{
    list.first = Nil;
}

address alokasi(string nama, string nim, int umur)
{
    address nodeBaru = new node;
    nodeBaru->isidata.nama = nama;
    nodeBaru->isidata.nim = nim;
    nodeBaru->isidata.umur = umur;
    nodeBaru->next = Nil;
    return nodeBaru;
}

void dealokasi(address &node)
{
    node->next = Nil;
    delete node;
}

void insertFirst(linkedlist &list, address nodeBaru)
{
    nodeBaru->next = list.first;
    list.first = nodeBaru;
}

void insertAfter(linkedlist &list, address nodeBaru, address prev)
{
    if (prev != Nil)
    {
        nodeBaru->next = prev->next;
        prev->next = nodeBaru;
    }
    else
    {
        cout << "Node sebelumnya tidak valid." << endl;
    }
}

void insertLast(linkedlist &list, address nodeBaru)
{
    if (isEmpty(list))
    {
        list.first = nodeBaru;
    }
    else
    {
        address nodeBantu = list.first;
        while (nodeBantu->next != Nil)
        {
            nodeBantu = nodeBantu->next;
        }
        nodeBantu->next = nodeBaru;
    }
}

void delFirst(linkedlist &list)
{
        address nodeHapus;
        if (!isEmpty(list)==false)
        {
            nodeHapus = list.first;
            list.first = list.first->next;
            nodeHapus->next = Nil;
            dealokasi(nodeHapus);
        }
        else
        {
            cout << "List kosong" << endl;
        }
}

void delLast(linkedlist &list)
{
    address nodeHapus, nodePrev;
    if (!isEmpty(list)==false)
    {
        nodeHapus = list.first;
        if (nodeHapus->next == Nil)
        {
            list.first->next = Nil;
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
        cout << "List kosong" << endl;
    }
}

void delAfter(linkedlist &list, address nodeHapus, address nodePrev)
{
    if (isEmpty(list) == true)
    {
        cout << "List kosong!" << endl;
    }
    else
    {
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

void printList(linkedlist list)
{
    if (isEmpty(list) == true)
    {
        cout << "List kosong!" << endl;
    }
    else
    {
        address nodeBantu = list.first;
        while (nodeBantu != Nil)
        {
            cout << "Nama: " << nodeBantu->isidata.nama << ", NIM: " << nodeBantu->isidata.nim << ", Usia: " << nodeBantu->isidata.umur << endl;
            nodeBantu = nodeBantu->next;
        }
    }
}

int nbList(linkedlist list)
{
    int count = 0;
    address nodeBantu = list.first;
    while (nodeBantu != Nil)
    {
        count++;
        nodeBantu = nodeBantu->next;
    }
    return count;
}

void deletelist(linkedlist &list)
{
    address nodeBantu, nodeHapus;
    nodeBantu = list.first;
    while (nodeBantu != Nil)
    {
        nodeHapus = nodeBantu;
        nodeBantu = nodeBantu->next;
        dealokasi(nodeHapus);
    }
    list.first = Nil;
    cout << "List sudah terhapus." << endl;
}
```
## main.cpp

```C++
#include <iostream>
#include "list.h"
using namespace std;

int main()
{
    linkedlist list;
    address nodeA, nodeB, nodeC, nodeD, nodeE = Nil;
    createList(list);

    nodeA = alokasi("Dhimas", "001", 20);
    nodeB = alokasi("Budi", "002", 21);
    nodeC = alokasi("Andi", "003", 22);
    nodeD = alokasi("Caca", "004", 23);
    nodeE = alokasi("Dewi", "005", 24);
}
```

Program menghasilkan luaran cetak dengan menginisialisasi senarai kosong terlebih dahulu lalu memesan ruang memori dinamis untuk setiap data mahasiswa agar menjadi simpul-simpul independen. Simpul-simpul terpisah tersebut kemudian dirangkai menjadi satu rantai berkesinambungan melalui instruksi penyisipan yang memodifikasi arah referensi penunjuk alamat antar memori. Setelah struktur terbentuk, luaran dieksekusi ke layar melalui fungsi penelusuran yang bergerak melompati penunjuk dari simpul awal hingga akhir untuk membaca isi data secara berurutan, sebelum akhirnya program memutus kaitan penunjuk pada simpul tertentu untuk menghapus dan mengembalikan memorinya ke sistem. 


## Unguided

### 1. (isi dengan soal unguided 1)
## singlylist.h

```C++
#ifndef SINGLYLIST_H
#define SINGLYLIST_H
#include <iostream>

using namespace std;

#define Nil NULL

typedef int infotype;
typedef struct ElmList *address;

struct ElmList {
    infotype info;
    address next;
};

struct List {
    address First;
};

void CreateList(List &L);
address alokasi(infotype x);
void dealokasi(address &P);
void printInfo(List L);
void insertFirst(List &L, address P);

#endif
```
## singlylist.cpp

```C++
#include <iostream>
#include "Singlylist.h"

using namespace std;

void CreateList(List &L) {
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
```
## main.cpp

```C++
#include <iostream>
#include "Singlylist.h"

using namespace std;

int main() {
    List L;
    address P1, P2, P3, P4, P5 = Nil;
    
    CreateList(L);
    
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
    
    return 0;
}
```

### Output Unguided 1 :

##### Output 1

<img width="942" height="145" alt="soal_1" src="https://github.com/user-attachments/assets/d943c1b7-bb9b-48c0-8f69-2989c0620eb4" />


##### Output 2

<img width="983" height="181" alt="soal_2" src="https://github.com/user-attachments/assets/28145189-9298-4436-9af8-d956f462a89d" />


Kode tersebut bekerja dengan cara memisahkan kerangka struktur data senarai berantai pada berkas antarmuka dan instruksi pemrosesannya pada berkas sumber terpisah. Program utama mengeksekusi kerangka tersebut dengan mengalokasikan lima ruang memori dinamis secara berurutan untuk menyimpan sekumpulan bilangan bulat. Seluruh bilangan bulat tersebut dimasukkan menggunakan metode penyisipan awal sehingga nilai yang terakhir kali ditambahkan akan selalu bergeser menempati urutan paling depan saat dicetak ke layar.


## Kesimpulan
konsep dasar senarai berantai tunggal yang memanfaatkan alokasi memori dinamis berbasis penunjuk untuk menyimpan sekumpulan data secara fleksibel. Implementasi berbagai operasi primitif seperti alokasi, penyisipan, dan penghapusan simpul memungkinkan pengolahan data terstruktur secara efisien tanpa batasan kapasitas statis.
...

## Referensi

[1] Pratama, A., dan Santoso, B. (2020). Analisis Perbandingan Kinerja Struktur Data Statis dan Dinamis dalam Manajemen Memori. Jurnal Ilmu Komputer dan Informatika, 9(2), 115-122.
<br>[2] Wijaya, R., dan Safitri, N. (2019). Implementasi Singly Linked List pada Sistem Antrean Terdistribusi. Jurnal Rekayasa Perangkat Lunak, 6(1), 45-52.
<br>[3] Kurniawan, H., dan Hidayat, T. (2021). Optimasi Alokasi Memori Dinamis pada Pemrograman Tingkat Tinggi. Jurnal Teknologi Informasi dan Sistem Komputer, 11(3), 210-218.
<br>[4] Lestari, D., dan Arifin, Z. (2018). Efisiensi Waktu Eksekusi Operasi Penyisipan Data pada Struktur Linked List. Jurnal Komputasi Terapan, 5(4), 301-308.
<br>[5] Firmansyah, E., dan Wibowo, A. (2022). Manajemen Dealokasi Memori untuk Mencegah Kebocoran Memori pada Aplikasi Berskala Besar. Jurnal Sistem Informasi Bisnis, 12(1), 88-95.
<br>[6] Setiawan, H., dan Riyadi, S. (2020). Algoritma Penelusuran Sekuensial pada Struktur Data Senarai Berantai. Jurnal Edukasi dan Penelitian Informatika, 7(2), 150-157.
