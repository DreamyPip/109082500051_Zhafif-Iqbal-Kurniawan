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