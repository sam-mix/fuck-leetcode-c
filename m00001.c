#include <stdlib.h>

// 树的节点
typedef struct treeNode {
    int value; 
    struct treeNode *leftTree;
    struct treeNode *rightTree;
} treeNode;


treeNode * buildTree(int *arr) {
    treeNode * r = (treeNode *) malloc(sizeof(treeNode));
    return r;
}

int main(int argc, char const *argv[])
{
    /* code */
    return 0;
}
