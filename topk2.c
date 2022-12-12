/* topK算法实现 */

#include <stdio.h>

/* 调整小顶堆，pos：唯一违反堆性质的点的位置 */

void heapify(int *arr, const size_t len, size_t pos){

    int min;

    size_t child = (pos * 2) + 1; // 左孩子

    while (1)

    {

        if (child + 1 < len) // 有两个孩子

        {

            min = arr[child] < arr[child + 1] ?

                                              arr[child]
                                              : arr[++child];
        }

        else if (child + 1 == len) // 只有左孩子

        {

            min = arr[child];
        }

        else
            break; // 数组结束，调整完成

        if (arr[pos] > min)

        {

            arr[child] = arr[pos];

            arr[pos] = min;

            pos = child; // 更新父节点索引

            child = (pos * 2) + 1; // 下一个调整点
        }

        else
            break; // 已经是堆，无需调整
    }
}

/* 建立小顶堆 */

void build_heap(int *arr_k, const int k)

{

    int i;

    for (i = k / 2 - 1; i >= 0; --i)

    {

        heapify(arr_k, k, i);
    }
}

/* 选出数组中最大的k个数，存入数组arr_k中 */

void top_k(const int *arr, const size_t len,

           int *arr_k, const size_t k)

{

    size_t i;

    for (i = 0; i < k; ++i)
        arr_k[i] = arr[i];

    build_heap(arr_k, k); // 用arr的前k个数建堆

    for (; i < len; ++i)

    {

        arr_k[0] = arr[i];

        heapify(arr_k, k, 0);
    }
}

/* 测试代码 */

int main()

{

    int a[] = {8, 1, 2, 7, 3, 4, 5, 6, 9};

#define K 4

    size_t i;

    int arr_k[K];

    top_k(a, 9, arr_k, K);

    for (i = 0; i < K; ++i)

    {

        printf("%d ", arr_k[i]);
    }

    return 0;
}
