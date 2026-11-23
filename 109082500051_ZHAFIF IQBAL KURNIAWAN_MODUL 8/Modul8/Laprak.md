# <h1 align="center">Laporan Praktikum Modul 8 - QUEUE</h1>

<p align="center">Zhafif Iqbal Kurniawan - 109082500051</p>

## Dasar Teori

isi dengan penjelasan dasar teori disertai referensi jurnal (gunakan kurung siku [] untuk pernyataan yang mengambil refernsi dari jurnal).
contoh :
Linked list atau yang disebut juga senarai berantai adalah Salah satu bentuk struktur data yang berisi kumpulan data yang tersusun secara sekuensial, saling bersambungan, dinamis, dan terbatas[1]. Linked list terdiri dari sejumlah node atau simpul yang dihubungkan secara linier dengan bantuan pointer.

### A. ...<br/>

...

#### 1. ...
Queue merupakan struktur data linear yang bekerja berdasarkan prinsip First In First Out (FIFO), di mana elemen yang pertama kali masuk ke dalam antrean akan diproses dan dikeluarkan terlebih dahulu [1].

#### 2. ...
Mekanisme manipulasi data pada queue terbagi pada dua titik akses berbeda, yaitu operasi penyisipan elemen baru (enqueue) melalui bagian belakang (tail) dan operasi pengambilan elemen (dequeue) melalui bagian depan (head) [2].

#### 3. ...
Implementasi antrean menggunakan struktur array statis menyimpan elemen pada blok memori berurutan dengan kapasitas terbatas, sehingga penghapusan elemen terdepan memerlukan pergeseran indeks atau penyesuaian posisi penunjuk agar ruang kosong tetap dapat dimanfaatkan [3].

### B. ...<br/>

...

#### 1. ...
Struktur circular queue mengatasi keterbatasan pergeseran elemen pada array linear dengan menghubungkan indeks akhir kembali ke indeks awal secara melingkar menggunakan operasi modulo agar pengelolaan ruang penyangga berjalan efisien [4].

#### 2. ...
Implementasi antrean menggunakan representasi pointer membentuk rangkaian simpul dinamis yang saling terhubung melalui alamat memori sehingga kapasitas antrean dapat bertumbuh secara fleksibel selama memori sistem masih tersedia [5].

#### 3. ...
Penerapan Tipe Data Abstrak (ADT) pada queue memisahkan spesifikasi antarmuka operasi antrean dari struktur penyimpanan internalnya sehingga representasi array maupun pointer dapat dipertukarkan tanpa mengubah logika program pemanggil [6].

## Guided 1

### 1. ...

## queue.h

```C++
#ifndef QUEUE_H
#define QUEUE_H

#include <iostream>
#include <string>
using namespace std;

struct Node {
    string nama;
    Node *next;
};

struct queue {
    Node *head, *tail;
};

void CreateQueue(queue &Q);
bool isEmpty(queue Q);
bool isFull(queue Q);
void enQueue(queue &Q, const string &nama);
void deQueue(queue &Q);
void viewQueue(queue Q);
void clearQueue(queue &Q);
#endif

```
## queue.cpp

```C++

#include "queue.h"

void CreateQueue(queue &Q) {
    Q.head = Q.tail = nullptr;
}

bool isEmpty(queue Q) {
    return Q.head == nullptr;
}

bool isFull(queue Q) {
    int maxQueue = 20;
    int counter = 0;
    Node *nodeBantu = Q.head;
    while (nodeBantu != Q.tail) {
        counter++;
        nodeBantu = nodeBantu->next;
    }
    return counter >= maxQueue;
}

void enQueue(queue &Q, const string &nama) {
    if (isEmpty(Q)) {
        Node *nodeBaru = new Node(nama, nullptr);
        Q.head = Q.tail = nodeBaru;
    } else {
        Node *nodeBaru = new Node(nama, nullptr);
        Q.tail->next = nodeBaru;
        Q.tail = nodeBaru;
    }
}

void deQueue(queue &Q) {
    if (isEmpty(Q)) {
        cout << "Queue kosong";
    } else {
        Node *nodeHapus = Q.head;
        Q.head = Q.head->next;
        delete nodeHapus;
    }
}

void viewQueue(queue Q) {
    if (isEmpty(Q)) {
        cout << "Queue kosong";
    } else {
        Node *nodeBantu = Q.head;
        while (nodeBantu != nullptr) {
            cout << "Nama : " << nodeBantu->nama;
            nodeBantu = nodeBantu->next;
        }
    }
}


void clearQueue(queue &Q) {
    if (!isEmpty(Q)) {
        cout << "Queue kosong";
    }else{
        Node *nodeHapus = Q.head;
        while (nodeHapus != nullptr) {
            nodeHapus = nodeHapus->next;
            deQueue(Q);
        }
    }
}

```
## main.cpp

```C++
#include "queue.h"
#include <iostream>
using namespace std;

int main() {
    queue Q;
    CreateQueue(Q);

    enQueue(Q, "Gilar");
    enQueue(Q, "Aflah");
    enQueue(Q, "Aedil");
    enQueue(Q, "Fathir");
    enQueue(Q, "Kevin");
    enQueue(Q, "Zulfa");
    cout << endl;

    cout << "--- Isi Queue Setelah enQueue ---" << endl;
    viewQueue(Q);

    clearQueue(Q);
    return 0;
}
```

Program ini membangun antrean dinamis berbasis pointer untuk menyimpan data nama, lalu memasukkan enam nama secara berurutan ke bagian belakang antrean. Setelah menampilkan seluruh daftar nama dari urutan paling depan, program mengakhiri eksekusi dengan mengosongkan kembali isi antrean tersebut.

## Guided 2

### 1. ...

## queue.h

```C++
#ifndef QUEUE_H
#define QUEUE_H

#include <iostream>
#include <string>
using namespace std;

const int MAX = 5;

struct queue {
    string nama[MAX];
    int head;
    int tail;
};

bool isEmpty(queue Q);
bool isFull(queue Q);
void createQueue(queue &Q);
void enQueue(queue &Q, string nama);
void deQueue(queue &Q);
void viewQueue(queue Q);

#endif

```
## stack.cpp

```C++
#include "queue.h"
#include <iostream>

using namespace std;

// NOTE : 
// Implementasi 1 = head diam, tail bergerak (Queue Linear Statis, kerana head nya tetap diam)
// Implementasi 2 = head bergerak, tail bergerak (Queue Linear Dinamis, karena head & tail nya sama-sama bergerak)
// Implementasi 3 = head dan tail berputar (Queue Circular, karena jika udh mentok tapi masih ada space, diputar sehingga tail bisa ada didepan head)

bool isEmpty(queue Q){
    if(Q.head == -1 && Q.tail == -1){
        return true;
    } else {
        return false;
    }
}

//isFull implmenetasi 1 & 2
bool isFull(queue Q){
    if(Q.tail == MAKSIMAL - 1){
        return true;
    } else {
        return false;
    }
}

//isFull implementasi 3
// bool isFull(queue Q){
//     if((Q.tail + 1) % MAKSIMAL == Q.head){
//         return true;
//     } else {
//         return false;
//     }
// }

void CreateQueue(queue &Q){ //terbentuk queue dengan head = -1 dan tail = -1 
    Q.head = -1;
    Q.tail = -1;
}
 

//enqueue implementasi 1 & 2
void enQueue(queue &Q, string nama){
    if(isFull(Q) == true){
        cout << "Queue sudah penuh!" << endl;
    } else {
        if(isEmpty(Q) == true){
            Q.head = Q.tail = 0;
        } else {
            Q.tail++;
        }
        Q.nama[Q.tail] = nama;
        cout << "nama " << nama << " berhasil ditambahkan kedalam queue!" << endl;
    }
}

// //enQueue implementasi 3
// void enQueue(queue &Q, string nama){
//     if(isFull(Q) == true){
//         cout << "Queue sudah penuh!" << endl;
//     } else {
//         if(isEmpty(Q) == true){
//             Q.head = Q.tail = 0;
//         } else {
//             Q.tail = (Q.tail + 1) % MAKSIMAL; // bergerak melingkar
//         }
//         Q.nama[Q.tail] = nama;
//         cout << "nama " << nama << " berhasil ditambahkan kedalam queue!" << endl;
//     }
// }

//dequeue implementasi 1
void deQueue(queue &Q){
    if(isEmpty(Q) == true){
        cout << "Queue kosong!" << endl;
    } else {
        cout << "Mengahapus data " << Q.nama[Q.head] << "..." << endl;
        for(int i = 0; i < Q.tail; i++){
            Q.nama[i] =  Q.nama[i+1];
        }
        Q.tail--;
        if(Q.tail < 0){ //kalo semua isi queue nya udh dikeluarin, set head & tail ke -1
            Q.head = -1;
            Q.tail = -1;
        }
    }
}

// //dequeue implementasi 2
// void deQueue(queue &Q){
//     if(isEmpty(Q) == true){
//         cout << "Queue kosong!" << endl;
//     } else {
//         cout << "Mengahapus data " << Q.nama[Q.head] << "..." << endl;
//         Q.head++;
//         if(Q.head > Q.tail){ //kalo elemennya udh abis (head akan lebih 1 dari tail), maka reset ulang head & tail ke -1
//             Q.head = -1;
//             Q.tail = -1;
//         }
//     }
// }

// //deQueue implementasi 3
// void deQueue(queue &Q){
//     if(isEmpty(Q) == true){
//         cout << "Queue kosong!" << endl;
//     } else {
//         cout << "Mengahapus data " << Q.nama[Q.head] << "..." << endl;
//         if(Q.head == Q.tail){ //kalo elemennya tinggal 1, langsungkan saja head & tail nya reset ke -1
//             Q.head = -1;
//             Q.tail = -1;
//         } else {
//             Q.head = (Q.head + 1) % MAKSIMAL; // bergerak melingkar
//         }
//     }
// }

//viewQueue implementasi 1 & 2
void viewQueue(queue Q){
    if(isEmpty(Q) == true){
        cout << "Queue kosong!" << endl;
    } else {
        for(int i = Q.head; i <= Q.tail; i++){
            cout << i -  Q.head + 1 << ". " << Q.nama[i] << endl;
        }
    }
    cout << endl;
}

// //viewQueue implementasi 3
// void viewQueue(queue Q){
//     if(isEmpty(Q) == true){
//         cout << "Queue kosong!" << endl;
//     } else {
//         int i = Q.head;
//         int count = 1;
//         while(true){
//             cout << count << ". " << Q.nama[i] << endl;
//             if(i == Q.tail){
//                 break;
//             }
//             i = (i + 1) % MAKSIMAL;
//             count++;
//         }   
//     }
// }

```
## main.cpp

```C++
#include "queue.h"
#include <iostream>

using namespace std;

int main(){
    queue Q;

    CreateQueue(Q);
    enQueue(Q, "Dimas");
    enQueue(Q, "Arvin");
    enQueue(Q, "Rizal");
    enQueue(Q, "Hafizh");
    enQueue(Q, "Fathur");
    enQueue(Q, "Daffa");
    cout << endl;

    cout << "--- Isi Queue Setelah enQueue ---" << endl;
    viewQueue(Q);
    cout << endl;

    deQueue(Q);
    deQueue(Q);
    deQueue(Q);
    deQueue(Q);
    // deQueue(Q);
    // deQueue(Q);
    cout << endl;

    cout << "--- Isi Queue Setelah deQueue ---" << endl;
    viewQueue(Q);

    return 0;
}
```

Program ini membuat antrean array berkapasitas lima nama dengan posisi head diam, lalu mencoba memasukkan enam nama sehingga data terakhir ditolak karena penuh. Setelah menampilkan isi antrean, program mengeluarkan empat nama terdepan sambil menggeser sisa data ke depan dan mencetak elemen yang tersisa.


## Unguided

### 1. (isi dengan soal unguided 1)
## queue.h

```C++
#ifndef QUEUE_H
#define QUEUE_H

#include <iostream>

using namespace std;

typedef int infotype;

struct Queue {
    infotype info[5];
    int head;
    int tail;
};

void createQueue(Queue &Q);
bool isEmptyQueue(Queue Q);

bool isFullQueueAlt1(Queue Q);
void enqueueAlt1(Queue &Q, infotype x);
infotype dequeueAlt1(Queue &Q);
void printInfoAlt1(Queue Q);

bool isFullQueueAlt2(Queue Q);
void enqueueAlt2(Queue &Q, infotype x);
infotype dequeueAlt2(Queue &Q);
void printInfoAlt2(Queue Q);

bool isFullQueueAlt3(Queue Q);
void enqueueAlt3(Queue &Q, infotype x);
infotype dequeueAlt3(Queue &Q);
void printInfoAlt3(Queue Q);

#endif
```
## queue.cpp

```C++
#include "queue.h"

void createQueue(Queue &Q) {
    Q.head = -1;
    Q.tail = -1;
}

bool isEmptyQueue(Queue Q) {
    return (Q.head == -1 && Q.tail == -1);
}

bool isFullQueueAlt1(Queue Q) {
    return (Q.tail == 4);
}

void enqueueAlt1(Queue &Q, infotype x) {
    if (!isFullQueueAlt1(Q)) {
        if (isEmptyQueue(Q)) {
            Q.head = 0;
            Q.tail = 0;
            Q.info[0] = x;
        } else {
            Q.tail++;
            Q.info[Q.tail] = x;
        }
    }
}

infotype dequeueAlt1(Queue &Q) {
    infotype x = -1;
    if (!isEmptyQueue(Q)) {
        x = Q.info[Q.head];
        if (Q.head == Q.tail) {
            Q.head = -1;
            Q.tail = -1;
        } else {
            for (int i = Q.head; i < Q.tail; i++) {
                Q.info[i] = Q.info[i + 1];
            }
            Q.tail--;
        }
    }
    return x;
}

void printInfoAlt1(Queue Q) {
    if (isEmptyQueue(Q)) {
        cout << Q.head << " - " << Q.tail << " | empty queue" << endl;
    } else {
        cout << Q.head << " - " << Q.tail << "   | ";
        for (int i = Q.head; i <= Q.tail; i++) {
            cout << Q.info[i] << " ";
        }
        cout << endl;
    }
}

bool isFullQueueAlt2(Queue Q) {
    return (Q.head == 0 && Q.tail == 4);
}

void enqueueAlt2(Queue &Q, infotype x) {
    if (!isFullQueueAlt2(Q)) {
        if (isEmptyQueue(Q)) {
            Q.head = 0;
            Q.tail = 0;
            Q.info[0] = x;
        } else {
            if (Q.tail == 4) {
                int j = 0;
                for (int i = Q.head; i <= Q.tail; i++) {
                    Q.info[j] = Q.info[i];
                    j++;
                }
                Q.tail = Q.tail - Q.head;
                Q.head = 0;
            }
            Q.tail++;
            Q.info[Q.tail] = x;
        }
    }
}

infotype dequeueAlt2(Queue &Q) {
    infotype x = -1;
    if (!isEmptyQueue(Q)) {
        x = Q.info[Q.head];
        if (Q.head == Q.tail) {
            Q.head = -1;
            Q.tail = -1;
        } else {
            Q.head++;
        }
    }
    return x;
}

void printInfoAlt2(Queue Q) {
    if (isEmptyQueue(Q)) {
        cout << Q.head << " - " << Q.tail << " | empty queue" << endl;
    } else {
        cout << Q.head << " - " << Q.tail << "   | ";
        for (int i = Q.head; i <= Q.tail; i++) {
            cout << Q.info[i] << " ";
        }
        cout << endl;
    }
}

bool isFullQueueAlt3(Queue Q) {
    return ((Q.tail + 1) % 5 == Q.head);
}

void enqueueAlt3(Queue &Q, infotype x) {
    if (!isFullQueueAlt3(Q)) {
        if (isEmptyQueue(Q)) {
            Q.head = 0;
            Q.tail = 0;
            Q.info[0] = x;
        } else {
            Q.tail = (Q.tail + 1) % 5;
            Q.info[Q.tail] = x;
        }
    }
}

infotype dequeueAlt3(Queue &Q) {
    infotype x = -1;
    if (!isEmptyQueue(Q)) {
        x = Q.info[Q.head];
        if (Q.head == Q.tail) {
            Q.head = -1;
            Q.tail = -1;
        } else {
            Q.head = (Q.head + 1) % 5;
        }
    }
    return x;
}

void printInfoAlt3(Queue Q) {
    if (isEmptyQueue(Q)) {
        cout << Q.head << " - " << Q.tail << " | empty queue" << endl;
    } else {
        cout << Q.head << " - " << Q.tail << "   | ";
        int i = Q.head;
        while (true) {
            cout << Q.info[i] << " ";
            if (i == Q.tail) {
                break;
            }
            i = (i + 1) % 5;
        }
        cout << endl;
    }
}
```
## main.cpp

```C++
#include <iostream>
#include "queue.h"

using namespace std;

int main() {
    Queue Q;

    cout << "=== ALTERNATIF 1 ===" << endl;
    createQueue(Q);
    printInfoAlt1(Q);
    enqueueAlt1(Q, 5); printInfoAlt1(Q);
    enqueueAlt1(Q, 2); printInfoAlt1(Q);
    enqueueAlt1(Q, 7); printInfoAlt1(Q);
    dequeueAlt1(Q); printInfoAlt1(Q);
    dequeueAlt1(Q); printInfoAlt1(Q);
    enqueueAlt1(Q, 4); printInfoAlt1(Q);
    dequeueAlt1(Q); printInfoAlt1(Q);
    dequeueAlt1(Q); printInfoAlt1(Q);
    cout << endl;

    cout << "=== ALTERNATIF 2 ===" << endl;
    createQueue(Q);
    printInfoAlt2(Q);
    enqueueAlt2(Q, 5); printInfoAlt2(Q);
    enqueueAlt2(Q, 2); printInfoAlt2(Q);
    enqueueAlt2(Q, 7); printInfoAlt2(Q);
    dequeueAlt2(Q); printInfoAlt2(Q);
    dequeueAlt2(Q); printInfoAlt2(Q);
    enqueueAlt2(Q, 4); printInfoAlt2(Q);
    dequeueAlt2(Q); printInfoAlt2(Q);
    dequeueAlt2(Q); printInfoAlt2(Q);
    cout << endl;

    cout << "=== ALTERNATIF 3 ===" << endl;
    createQueue(Q);
    printInfoAlt3(Q);
    enqueueAlt3(Q, 5); printInfoAlt3(Q);
    enqueueAlt3(Q, 2); printInfoAlt3(Q);
    enqueueAlt3(Q, 7); printInfoAlt3(Q);
    dequeueAlt3(Q); printInfoAlt3(Q);
    dequeueAlt3(Q); printInfoAlt3(Q);
    enqueueAlt3(Q, 4); printInfoAlt3(Q);
    dequeueAlt3(Q); printInfoAlt3(Q);
    dequeueAlt3(Q); printInfoAlt3(Q);

    return 0;
}
```

### Output Unguided 1 :

##### Output 1

<img width="852" height="298" alt="Soal_1" src="https://github.com/user-attachments/assets/ab3359de-0290-4b68-b1fa-ba86cccfa8e9" />


##### Output 2

<img width="742" height="496" alt="Soal_2" src="https://github.com/user-attachments/assets/706506f4-42ea-4295-ae6b-d88936319e66" />


Program ini menjalankan simulasi antrean berbasis array lima elemen menggunakan tiga mekanisme berbeda secara berurutan. Alternatif 1 selalu menggeser sisa data ke depan tiap ada elemen keluar, Alternatif 2 memajukan posisi head dan baru menggeser data saat mentok di ujung array, sedangkan Alternatif 3 memutar indeks head dan tail kembali ke awal tanpa perlu menggeser data sama sekali.


## Kesimpulan
Queue merupakan struktur data berprinsip First In First Out (FIFO) dengan operasi penambahan (enqueue) di bagian belakang (tail) dan penghapusan (dequeue) di bagian depan (head), yang dapat diimplementasikan secara dinamis memakai pointer maupun statis memakai array. Pada representasi array, mekanisme antrean melingkar (circular buffer) pada Alternatif 3 terbukti paling efisien karena indeks head dan tail dapat berputar otomatis tanpa perlu menggeser elemen data seperti pada Alternatif 1 dan 2.
...

## Referensi

[1] Little, J. D. C. (1961). A proof for the queuing formula: L = λW. Operations Research, 9(3), 383–387.
<br>[2] Hood, R., & Melville, R. (1981). Real-time queue operations in pure LISP. Information Processing Letters, 13(2), 50–54.
<br>[3] Burton, F. W. (1982). An efficient functional implementation of FIFO queues. Information Processing Letters, 14(5), 205–206.
<br>[4] Hoare, C. A. R. (1974). Monitors: An operating system structuring concept. Communications of the ACM, 17(10), 549–557.
<br>[5] Mellor-Crummey, J. M., & Scott, M. L. (1991). Algorithms for scalable synchronization on shared-memory multiprocessors. ACM Transactions on Computer Systems, 9(1), 21–65.
<br>[6] Guttag, J. (1977). Abstract data types and the development of data structures. Communications of the ACM, 20(6), 396–404.
