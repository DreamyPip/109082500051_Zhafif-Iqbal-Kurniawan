#include <iostream>
#include <string>
#include "pelajaran.h"

using namespace std;

int main() {
    string namapel = "Valo?";
    string kodepel = "Cihuy";
    pelajaran pel = create_pelajaran(namapel, kodepel);
    tampil_pelajaran(pel);
    
    return 0;
}