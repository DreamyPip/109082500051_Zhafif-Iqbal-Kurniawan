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