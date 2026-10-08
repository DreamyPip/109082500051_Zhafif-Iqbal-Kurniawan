#ifndef BUKU_H
#define BUKU_H
#include <iostream>

using namespace std;

struct Buku {
    string judul;
    int halaman;
    string penulis;
};

void editIsi(string judul, int halaman, string penulis, Buku &buku);
void tampilkanIsiBuku(Buku buku);
bool checkPenulis(Buku buku);

#endif