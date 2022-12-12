
#include <stdio.h>
// topk   第k大的数
void maxHeapify(float *a, int i, int heapSize)
{
    int l = i * 2 + 1, r = i * 2 + 2, largest = i;

    if (l < heapSize && a[l] > a[largest])
    {
        largest = l;
    }

    if (r < heapSize && a[r] > a[largest])
    {
        largest = r;
    }

    if (largest != i)
    {
        float t = a[i];
        a[i] = a[largest], a[largest] = t;
        maxHeapify(a, largest, heapSize);
    }
}

// topk   第k小的数
void minHeapify(float *a, int i, int heapSize)
{
    int l = i * 2 + 1, r = i * 2 + 2, largest = i;

    if (l < heapSize && a[l] < a[largest])
    {
        largest = l;
    }

    if (r < heapSize && a[r] < a[largest])
    {
        largest = r;
    }

    if (largest != i)
    {
        float t = a[i];
        a[i] = a[largest], a[largest] = t;
        minHeapify(a, largest, heapSize);
    }
}

void buildMinHeap(float *a, int heapSize)
{
    for (int i = heapSize / 2; i >= 0; --i)
    {
        maxHeapify(a, i, heapSize);
    }
}

float *findKthLargest(float *nums, int numsSize, int k)
{
    int heapSize = numsSize;
    buildMinHeap(nums, heapSize);

    for (int i = numsSize - 1; i >= numsSize - k + 1; --i)
    {
        float t = nums[0];
        nums[0] = nums[i], nums[i] = t;
        --heapSize;
        maxHeapify(nums, 0, heapSize);
    }

    return nums;
}

int main(void)
{
    float array[10] = {0.14, 0.12, 0.13, 0.11, 0.18, 0.11, 0.17, 0.19, 0.16, 0.15};

    float *ans = findKthLargest(array, 10, 5);
    for (int i = 0; i < 5; i++)
    {
        printf("%.2f ", ans[i]);
    }

    return 0;
}
