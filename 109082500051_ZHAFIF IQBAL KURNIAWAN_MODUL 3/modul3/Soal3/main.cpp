#include <iostream>
#include "array.h"

using namespace std;

int main() {
    int matriks1[3][3] = {{6, 7, 6}, {9, 9, 1}, {1, 2, 1}};
    int matriks2[3][3] = {{1, 2, 1}, {1, 9, 9}, {6, 7, 6}};
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