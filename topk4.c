/*
寻找n个不同数据中第k大的元素
接口：int top_k(int *a,int n,int k);
输入：一个由n个不同元素组成的数组a，a的长度n以及k的值
输出：数组a中第k大的元素
*/

#include <stdio.h>
#define max_size 10001

float min_heap[max_size] = {0.0}; // 假定数组中的数都是正数
int heap_count = 0;

void reheap(int k) // 更新根节点的数值后，调整使其符合最小堆
{
    int root = 0;
    int child = root * 2 + 1;
    while (child < k)
    {
        if ((child + 1) < k && min_heap[child + 1] < min_heap[child])
            child++;
        if (min_heap[root] < min_heap[child])
            return;
        float temp = min_heap[root];
        min_heap[root] = min_heap[child];
        min_heap[child] = temp;
        root = child;
        child = root * 2 + 1;
    }
}

float top_k(float *a, int n, int k)
{
    for (int i = 0; i < n; i++)
        if (a[i] > min_heap[0])
        {
            min_heap[0] = a[i];
            reheap(k);
        }
    return min_heap[0];
}

int main(void)
{

    float array[10] = {1.2, 4.5, 2.4, 5.6, 7.6, 7.8, 5.8, 7.8, 9.1, 3.8};
    float x = top_k(array, 9, 5);
    printf("%.1f\n", x);

    return 0;
}
