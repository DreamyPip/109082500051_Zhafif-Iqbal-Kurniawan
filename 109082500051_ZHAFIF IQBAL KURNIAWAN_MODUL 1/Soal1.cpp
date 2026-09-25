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