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