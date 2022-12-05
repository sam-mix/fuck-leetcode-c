#include <stdlib.h>
#include <stdio.h>

int main(int argc, char const *argv[]){
    int *p;
    p = (int *)malloc(sizeof(int));
    *p= 666;
    printf("%d\n",*p);
    free(p);
    return 0;
}
