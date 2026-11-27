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