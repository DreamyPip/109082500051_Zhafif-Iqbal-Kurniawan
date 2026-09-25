#include <iostream>
using namespace std;

int main() {
    int n;
    
    cout << "input: ";
    cin >> n;
    
    cout << "output:\n";
    
    for (int i = n; i >= 0; i--) {
        
        for (int j = 0; j < n - i; j++) {
            cout << "  ";
        }

        for (int j = i; j >= 1; j--) {
            cout << j << " ";
        }

        cout << "*";

        for (int j = 1; j <= i; j++) {
            cout << " " << j;
        }
        
        cout << endl;
    }

    return 0;
}