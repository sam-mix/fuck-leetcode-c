#include <stdio.h>
#include <stdlib.h>
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
void AdjustDown(int*a,int n,int root)
{
    int parent=root;
    int child=parent*2+1;
    while(child<n)
    {
        if(child+1<n&&a[child+1]>a[child])
        {
            ++child;
        }

        if(a[child]>a[parent])
        {
            int tmp=a[child];
            a[child]=a[parent];
            a[parent]=tmp;

            parent=child;
            child=parent*2+1;
        }
        else
        {
            break;
        }

    }
}

int* getLeastNumbers(int* arr, int arrSize, int k, int* returnSize){
    *returnSize=k;


    int* retArr=(int*)malloc(sizeof(int)*k);
    if(k==0)
        return retArr;
    for(int i=0;i<k;++i)
    {
        retArr[i]=arr[i];
    }
    //建k个的大堆
    for(int i=(k-2)/2;i>=0;--i)
    {
        AdjustDown(retArr,k,i);
    }

    for(int j=k;j<arrSize;++j)
    {
        if(arr[j]<retArr[0])
        {
            retArr[0]=arr[j];
            AdjustDown(retArr,k,0);
        }
    }
    

    return retArr;

}

int main() {
    int arr1[] = {1,2,3,4,5,6,7,8,9,10};
    int *arr = arr1;
    int arrSize = 10;
    int k = 3;
    int * returnSize = (int*)malloc(sizeof(int*));
    int* ans = getLeastNumbers(arr, arrSize,k, returnSize);
    for (int i = 0; i < *returnSize;i++) {
        printf("%d\t", *ans++);
    }
    return 0;
}