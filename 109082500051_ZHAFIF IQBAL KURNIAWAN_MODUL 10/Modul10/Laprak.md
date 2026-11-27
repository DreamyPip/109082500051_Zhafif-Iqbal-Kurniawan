# <h1 align="center">Laporan Praktikum Modul 10 - TREE</h1>

<p align="center">Zhafif Iqbal Kurniawan - 109082500051</p>

## Dasar Teori

isi dengan penjelasan dasar teori disertai referensi jurnal (gunakan kurung siku [] untuk pernyataan yang mengambil refernsi dari jurnal).
contoh :
Linked list atau yang disebut juga senarai berantai adalah Salah satu bentuk struktur data yang berisi kumpulan data yang tersusun secara sekuensial, saling bersambungan, dinamis, dan terbatas[1]. Linked list terdiri dari sejumlah node atau simpul yang dihubungkan secara linier dengan bantuan pointer.

### A. ...<br/>

...

#### 1. ...
Rekursi merupakan mekanisme pemanggilan subprogram terhadap dirinya sendiri secara berulang dengan parameter yang diperkecil hingga mencapai kondisi berhenti (base case) yang terdefinisi [1].

#### 2. ...
Tree merupakan struktur data hierarkis non-linear berupa graf tak berarah yang saling terhubung tanpa mengandung sirkuit, di mana satu simpul utama bertindak sebagai akar (root) dan simpul-simpul di bawahnya memiliki tepat satu simpul induk (parent) [2].

#### 3. ...
Binary Tree membatasi setiap simpul agar memiliki maksimal dua anak (left child dan right child), yang secara dinamis direpresentasikan menggunakan simpul-simpul memori yang saling terhubung melalui pointer [3].

### B. ...<br/>

...

#### 1. ...
Binary Search Tree merupakan varian pohon biner terurut yang menempatkan setiap simpul bernilai lebih kecil pada subpohon kiri dan simpul bernilai lebih besar pada subpohon kanan untuk mempercepat proses penyisipan, pencarian, maupun penghapusan data [4].

#### 2. ...
Penelusuran seluruh simpul pada pohon biner dilakukan secara rekursif melalui tiga urutan kunjungan utama, yaitu Pre-order (root–left–right), In-order (left–root–right), dan Post-order (left–right–root) [5].

#### 3. ...
AVL Tree merupakan pengembangan Binary Search Tree yang menjaga selisih tinggi (height) antara subpohon kiri dan subpohon kanan pada setiap simpul agar tidak melebihi satu tingkat sehingga performa pencarian tetap optimal [6].

## Guided 1

### 1. ...

## bst.h

```C++
#ifndef BST_H
#define BST_H
#define Nil NULL

#include <iostream>
using namespace std;

typedef struct BST *node;

struct Node {
    int angka;
    node left, right;
};

typedef node BinTree;

bool IsEmpty(BinTree tree);
void createTree(BinTree &tree);
node alokasi(int angka);
void dealokasi(node nodeHapus);

void insertNode(BinTree &tree, node nodeBaru);
void searchByData(BinTree tree, int angka);
void preorder(BinTree tree);
void inorder(BinTree tree);
void postorder(BinTree tree);

bool deleteNode(BinTree &tree, int angka);
node mostRight(BinTree tree);
node mostLeft(BinTree tree);
void deletetree(BinTree &tree);
int size(BinTree tree);
int height(BinTree tree);
#endif

```
## bst.cpp

```C++

#include "bst.h"
#include <iostream>

using namespace std;
// NOTE : parameter tree disini maksudnya merujuk ke node; baik itu node root atau node lain dari tree

bool isEmpty(BinTree tree)
{
    if (tree == Nil)
    {
        return true;
    }
    else
    {
        return false;
    }
}

void createTree(BinTree &tree)
{
    tree = Nil;
}

node alokasi(int angkaInput)
{
    node nodeBaru = new BST;
    nodeBaru->angka = angkaInput;
    nodeBaru->left = Nil;
    nodeBaru->right = Nil;
    return nodeBaru;
}

void dealokasi(node nodeHapus)
{
    delete nodeHapus;
}

void insertNode(BinTree &tree, node nodeBaru)
{
    if (tree == Nil)
    {
        tree = nodeBaru;
        cout << "Node " << nodeBaru->angka << " berhasil ditambahkan ke dalam tree!" << endl;
        return;
    }
    else if (nodeBaru->angka < tree->angka)
    {
        insertNode(tree->left, nodeBaru);
    }
    else if (nodeBaru->angka > tree->angka)
    {
        insertNode(tree->right, nodeBaru);
    }
}

void searchByData(BinTree tree, int angkaCari)
{
    if (isEmpty(tree) == true)
    {
        cout << "Tree kosong!" << endl;
    }
    else
    {
        node nodeBantu = tree;
        node parent = Nil;
        bool ketemu = false;
        while (nodeBantu != Nil)
        {
            if (angkaCari < nodeBantu->angka)
            {
                parent = nodeBantu;
                nodeBantu = nodeBantu->left;
            }
            else if (angkaCari > nodeBantu->angka)
            {
                parent = nodeBantu;
                nodeBantu = nodeBantu->right;
            }
            else if (angkaCari == nodeBantu->angka)
            {
                ketemu = true;
                break;
            }
        }
        if (ketemu == false)
        {
            cout << "Data tidak ditemukan" << endl;
        }
        else if (ketemu == true)
        {
            cout << "Data ditemukan didalam tree!" << endl;
            cout << "Data Angka : " << nodeBantu->angka << endl;

            // menampilkan parentnya & pengecekan sibling
            node sibling = Nil;
            if (parent != Nil)
            {
                cout << "Parent : " << parent->angka << endl;
                if (parent->left == nodeBantu)
                {
                    sibling = parent->right;
                }
                else if (parent->right == nodeBantu)
                {
                    sibling = parent->left;
                }
            }
            else
            {
                cout << "Parent : - (node root)" << endl;
            }

            // menampilkan siblingnya
            if (sibling != Nil)
            {
                cout << "Sibling : " << sibling->angka << endl;
            }
            else
            {
                cout << "Sibling : - " << endl;
            }

            // menampilkan childnya
            if (nodeBantu->left != Nil)
            {
                cout << "Child kiri : " << nodeBantu->left->angka << endl;
            }
            else if (nodeBantu->left == Nil)
            {
                cout << "Child kiri : -" << endl;
            }
            if (nodeBantu->right != Nil)
            {
                cout << "Child kanan : " << nodeBantu->right->angka << endl;
            }
            else if (nodeBantu->right == Nil)
            {
                cout << "Child kanan : -" << endl;
            }
        }
    }
}

void preOrder(BinTree tree)
{ // tengah - kiri - kanan atau root - child kiri - child kanan
    if (tree == Nil)
    {
        return;
    }
    cout << tree->angka << " - ";
    preOrder(tree->left);
    preOrder(tree->right);
}

void inOrder(BinTree tree)
{ // kiri - tengah - kanan atau child kiri - root - child kanan
    if (tree == Nil)
    {
        return;
    }
    inOrder(tree->left);
    cout << tree->angka << " - ";
    inOrder(tree->right);
}

void postOrder(BinTree tree)
{ // kiri - kanan - tengah atau child kiri - child kanan - root
    if (tree == Nil)
    {
        return;
    }
    postOrder(tree->left);
    postOrder(tree->right);
    cout << tree->angka << " - ";
}

bool deleteNode(BinTree &tree, int angka)
{
    if (tree == Nil)
    {
        return false; // data tidak ditemukan di subtree ini
    }
    else
    {
        if (angka < tree->angka)
        {
            return deleteNode(tree->left, angka);
        }
        else if (angka > tree->angka)
        {
            return deleteNode(tree->right, angka);
        }
        else
        {
            // jika node yang mau dihapus ditemukan
            // Case 1 : node yang mau dihapus adalah leaf
            if (tree->left == Nil && tree->right == Nil)
            {
                node tmp = tree;
                tree = Nil;
                dealokasi(tmp);
            }
            // Case 2 : node yang mau dihapus hanya punya right child
            else if (tree->left == Nil)
            {
                node tmp = tree;
                tree = tree->right;
                dealokasi(tmp);
            }
            // Case 3 : node yang mau dihapus hanya punya left child
            else if (tree->right == Nil)
            {
                node tmp = tree;
                tree = tree->left;
                dealokasi(tmp);
            }
            // Case 4 : jika node yang mau dihapus punya dua child, maka ambil mostleft dari subtree kanan untuk menggantikan node yang mau dihapus
            else
            {
                // mostleft dari subtree kanan = node successor (node penerus)
                node successor = mostLeft(tree->right);
                // salin data successor ke node saat ini
                tree->angka = successor->angka;
                // hapus successor pada subtree kanan
                return deleteNode(tree->right, successor->angka);
            }
            return true; // berhasil dihapus
        }
    }
}

node mostRight(BinTree tree)
{
    while (tree->right != Nil)
    {
        tree = tree->right;
    }
    return tree;
}

node mostLeft(BinTree tree)
{
    while (tree->left != Nil)
    {
        tree = tree->left;
    }
    return tree;
}

void deleteTree(BinTree &tree)
{
    if (tree == Nil)
    {
        return;
    }
    else
    {
        deleteTree(tree->left);
        deleteTree(tree->right);
        dealokasi(tree);
        tree = Nil;
    }
}

int size(BinTree tree)
{ // mengembalikan jumlah semua node
    if (isEmpty(tree) == true)
    {
        return 0;
    }
    else
    {
        return 1 + size(tree->left) + size(tree->right);
    }
    cout << endl;
}

int height(BinTree tree)
{ // mengembalikan jumlah level tree
    if (isEmpty(tree) == true)
    {
        return -1; // tree kosong jika height = -1
    }
    else
    {
        int hl = height(tree->left);
        int hr = height(tree->right);
        int maxHeight;
        if (hl > hr)
        {
            maxHeight = hl;
        }
        else
        {
            maxHeight = hr;
        }
        return 1 + maxHeight;
    }
    cout << endl;
}

```
## main.cpp

```C++
#include <iostream>
#include "bst.h"

using namespace std;

int main()
{
    BinTree tree;
    createTree(tree);

    int pilih, angka;

    do
    {
        cout << "========= MENU BST =========" << endl;
        cout << "1. Insert Node" << endl;
        cout << "2. Delete Node" << endl;
        cout << "3. Search Data" << endl;
        cout << "4. Tampilkan PreOrder" << endl;
        cout << "5. Tampilkan InOrder" << endl;
        cout << "6. Tampilkan PostOrder" << endl;
        cout << "7. Size Tree (jumlah node)" << endl;
        cout << "8. Height Tree (tinggi level)" << endl;
        cout << "9. Tampilkan mostright" << endl;
        cout << "10. Tampilkan mostleft" << endl;
        cout << "11. Delete Seluruh Tree" << endl;
        cout << "0. Keluar" << endl;
        cout << "pilihan anda : ";
        cin >> pilih;
        cout << endl;

        switch (pilih)
        {
        case 1:
            cout << "Masukkan angka: ";
            cin >> angka;
            insertNode(tree, alokasi(angka));
            cout << endl;
            break;

        case 2:
            if (isEmpty(tree) == true)
            {
                cout << "Tree kosong!" << endl;
            }
            else
            {
                cout << "Masukkan angka yang ingin dihapus: ";
                cin >> angka;
                if (deleteNode(tree, angka))
                {
                    cout << "Data " << angka << " berhasil dihapus!" << endl;
                }
                else
                {
                    cout << "Data " << angka << " tidak ditemukan!" << endl;
                }
            }
            cout << endl;
            break;

        case 3:
            if (isEmpty(tree) == true)
            {
                cout << "Tree kosong!" << endl;
            }
            else
            {
                cout << "Masukkan angka yang ingin dicari: ";
                cin >> angka;
                searchByData(tree, angka);
            }
            cout << endl;
            break;

        case 4:
            if (isEmpty(tree) == true)
            {
                cout << "Tree kosong!" << endl;
            }
            else
            {
                cout << "PreOrder : ";
                preOrder(tree);
                cout << endl;
            }
            cout << endl;
            break;

        case 5:
            if (isEmpty(tree) == true)
            {
                cout << "Tree kosong!" << endl;
            }
            else
            {
                cout << "InOrder : ";
                inOrder(tree);
                cout << endl;
            }
            cout << endl;
            break;

        case 6:
            if (isEmpty(tree) == true)
            {
                cout << "Tree kosong!" << endl;
            }
            else
            {
                cout << "PostOrder : ";
                postOrder(tree);
                cout << endl;
            }
            cout << endl;
            break;

        case 7:
            cout << "Size Tree = " << size(tree) << endl;
            cout << endl;
            break;

        case 8:
            cout << "Height Tree = " << height(tree) << endl;
            cout << endl;
            break;

        case 9:
            if (isEmpty(tree) == true)
            {
                cout << "Tree kosong!" << endl;
                cout << endl;
            }
            else
            {
                cout << "Mostright : " << mostRight(tree)->angka << endl;
                cout << endl;
            }
            break;

        case 10:
            if (isEmpty(tree) == true)
            {
                cout << "Tree kosong!" << endl;
                cout << endl;
            }
            else
            {
                cout << "Mostleft : " << mostLeft(tree)->angka << endl;
                cout << endl;
            }
            break;

        case 11:
            if (isEmpty(tree) == true)
            {
                cout << "Tree kosong!" << endl;
            }
            else
            {
                deleteTree(tree);
                cout << "Seluruh tree berhasil dihapus!" << endl;
            }
            cout << endl;
            break;

        case 0:
            cout << "Keluar dari program..." << endl;
            break;

        default:
            cout << "Pilihan tidak valid!" << endl;
            break;
        }

    } while (pilih != 0);

    return 0;
}
```

Program ini menyediakan menu pilihan untuk mengolah data angka pada Binary Search Tree, mulai dari menambah, menghapus, hingga mencari posisi suatu angka beserta info parent, sibling, dan child-nya. Di dalamnya juga ada fitur untuk menampilkan urutan isi pohon (PreOrder, InOrder, PostOrder), melihat nilai paling kiri dan kanan, menghitung jumlah serta tinggi node, sampai menghapus seluruh isi pohon sekaligus.



## Unguided

### 1. (isi dengan soal unguided 1)
## bstree.h

```C++
#ifndef BSTREE_H
#define BSTREE_H

#include <iostream>
#define Nil NULL

using namespace std;

typedef int infotype;
typedef struct Node *address;

struct Node {
    infotype info;
    address left;
    address right;
};

address alokasi(infotype x);
void insertNode(address &root, infotype x);
address findNode(infotype x, address root);
void printInorder(address root);
void printPreorder(address root);
void printPostorder(address root);

int hitungJumlahNode(address root);
int hitungTotalInfo(address root, int start = 0);
int hitungKedalaman(address root, int start = 0);

#endif
```
## bstree.cpp

```C++
#include "bstree.h"

address alokasi(infotype x) {
    address P = new Node;
    P->info = x;
    P->left = Nil;
    P->right = Nil;
    return P;
}

void insertNode(address &root, infotype x) {
    if (root == Nil) {
        root = alokasi(x);
    } else if (x < root->info) {
        insertNode(root->left, x);
    } else if (x > root->info) {
        insertNode(root->right, x);
    }
}

address findNode(infotype x, address root) {
    if (root == Nil || root->info == x) {
        return root;
    }
    if (x < root->info) {
        return findNode(x, root->left);
    }
    return findNode(x, root->right);
}

void printInorder(address root) {
    if (root != Nil) {
        printInorder(root->left);
        cout << root->info << " - ";
        printInorder(root->right);
    }
}

void printPreorder(address root) {
    if (root != Nil) {
        cout << root->info << " - ";
        printPreorder(root->left);
        printPreorder(root->right);
    }
}

void printPostorder(address root) {
    if (root != Nil) {
        printPostorder(root->left);
        printPostorder(root->right);
        cout << root->info << " - ";
    }
}

int hitungJumlahNode(address root) {
    if (root == Nil) {
        return 0;
    }
    return 1 + hitungJumlahNode(root->left) + hitungJumlahNode(root->right);
}

int hitungTotalInfo(address root, int start) {
    if (root == Nil) {
        return start;
    }
    return root->info + hitungTotalInfo(root->left, 0) + hitungTotalInfo(root->right, 0) + start;
}

int hitungKedalaman(address root, int start) {
    if (root == Nil) {
        return start;
    }
    int kiri = hitungKedalaman(root->left, start + 1);
    int kanan = hitungKedalaman(root->right, start + 1);
    return (kiri > kanan) ? kiri : kanan;
}
```
## main.cpp

```C++
#include <iostream>
#include "bstree.h"

using namespace std;

int main() {
    cout << "Hello world!" << endl;
    address root = Nil;

    insertNode(root, 1);
    insertNode(root, 2);
    insertNode(root, 6);
    insertNode(root, 4);
    insertNode(root, 5);
    insertNode(root, 3);
    insertNode(root, 6);
    insertNode(root, 7);

    cout << "In-order   : ";
    printInorder(root);
    cout << endl;

    cout << "Pre-order  : ";
    printPreorder(root);
    cout << endl;

    cout << "Post-order : ";
    printPostorder(root);
    cout << endl;

    cout << "kedalaman  : " << hitungKedalaman(root, 0) << endl;
    cout << "jumlah node: " << hitungJumlahNode(root) << endl;
    cout << "total      : " << hitungTotalInfo(root, 0) << endl;

    return 0;
}
```

### Output Unguided 1 :

##### Output 1



Program ini membangun struktur Binary Search Tree (BST) berbasis pointer dengan memasukkan deretan angka secara rekursif, di mana nilai lebih kecil diarahkan ke cabang kiri, nilai lebih besar ke cabang kanan, dan angka duplikat otomatis diabaikan. Setelah pohon terbentuk, program menampilkan isi simpul melalui traversal In-order, Pre-order, dan Post-order, lalu menghitung kedalaman pohon, jumlah simpul, serta total seluruh nilainya.


## Kesimpulan
Fungsi rekursif sangat efektif untuk menyederhanakan pengolahan struktur data non-linear seperti Tree selama memiliki kondisi berhenti yang jelas. Penerapan aturan Binary Search Tree (BST)—anak kiri lebih kecil dan anak kanan lebih besar dari induknya—membuat operasi penyisipan, pencarian, maupun penelusuran data (Pre-order, In-order, dan Post-order) berjalan terstruktur dan efisien.
...

## Referensi

[1] Burstall, R. M. (1969). Proving properties of programs by structural induction. The Computer Journal, 12(1), 41–48.
<br>[2] Otter, R. (1948). The number of trees. Annals of Mathematics, 49(3), 583–599.
<br>[3] Perlis, A. J., & Thornton, C. (1960). Symbol manipulation by threaded lists. Communications of the ACM, 3(4), 195–204.
<br>[4] Hibbard, T. N. (1962). Some combinatorial properties of certain trees with applications to searching and sorting. Journal of the ACM, 9(1), 13–28.
<br>[5] Morris, J. M. (1979). Traversing binary trees simply and cheaply. Information Processing Letters, 9(5), 197–200.
<br>[6] GAdelson-Velsky, G. M., & Landis, E. M. (1962). An algorithm for organization of information. Doklady Akademii Nauk SSSR, 146(2), 263–266.