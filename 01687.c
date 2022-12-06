// typedef struct {
//     int *elements;
//     int rear, front;
//     int capacity;
// } MyCircularDeque;

// MyCircularDeque* myCircularDequeCreate(int k) {
//     MyCircularDeque *obj = (MyCircularDeque *)malloc(sizeof(MyCircularDeque));
//     obj->capacity = k + 1;
//     obj->rear = obj->front = 0;
//     obj->elements = (int *)malloc(sizeof(int) * obj->capacity);
//     return obj;
// }

// bool myCircularDequeInsertFront(MyCircularDeque* obj, int value) {
//     if ((obj->rear + 1) % obj->capacity == obj->front) {
//         return false;
//     }
//     obj->front = (obj->front - 1 + obj->capacity) % obj->capacity;
//     obj->elements[obj->front] = value;
//     return true;
// }

// bool myCircularDequeInsertLast(MyCircularDeque* obj, int value) {
//     if ((obj->rear + 1) % obj->capacity == obj->front) {
//         return false;
//     }
//     obj->elements[obj->rear] = value;
//     obj->rear = (obj->rear + 1) % obj->capacity;
//     return true;
// }

// bool myCircularDequeDeleteFront(MyCircularDeque* obj) {
//     if (obj->rear == obj->front) {
//         return false;
//     }
//     obj->front = (obj->front + 1) % obj->capacity;
//     return true;
// }

// bool myCircularDequeDeleteLast(MyCircularDeque* obj) {
//     if (obj->rear == obj->front) {
//         return false;
//     }
//     obj->rear = (obj->rear - 1 + obj->capacity) % obj->capacity;
//     return true;
// }

// int myCircularDequeGetFront(MyCircularDeque* obj) {
//     if (obj->rear == obj->front) {
//         return -1;
//     }
//     return obj->elements[obj->front];
// }

// int myCircularDequeGetRear(MyCircularDeque* obj) {
//     if (obj->rear == obj->front) {
//         return -1;
//     }
//     return obj->elements[(obj->rear - 1 + obj->capacity) % obj->capacity];
// }

// bool myCircularDequeIsEmpty(MyCircularDeque* obj) {
//     return obj->rear == obj->front;
// }

// bool myCircularDequeIsFull(MyCircularDeque* obj) {
//     return (obj->rear + 1) % obj->capacity == obj->front;
// }

// void myCircularDequeFree(MyCircularDeque* obj) {
//     free(obj->elements);
//     free(obj);
// }

// int boxDelivering(int** boxes, int boxesSize, int* boxesColSize, int portsCount, int maxBoxes, int maxWeight) {
//     int n = boxesSize;
//     int p[n + 1], w[n + 1], neg[n + 1];
//     long long W[n + 1];
//     memset(neg, 0, sizeof(neg));
//     memset(W, 0, sizeof(W));
//     for (int i = 1; i <= n; ++i) {
//         p[i] = boxes[i - 1][0];
//         w[i] = boxes[i - 1][1];
//         if (i > 1) {
//             neg[i] = neg[i - 1] + (p[i - 1] != p[i]);
//         }
//         W[i] = W[i - 1] + w[i];
//     }
    
//     int f[n + 1], g[n + 1];    
//     memset(f, 0, sizeof(f));
//     memset(g, 0, sizeof(g));
//     MyCircularDeque *opt =  myCircularDequeCreate(n + 1);
//     myCircularDequeInsertLast(opt, 0);
//     for (int i = 1; i <= n; ++i) {
//         while (i - myCircularDequeGetFront(opt) > maxBoxes || 
//                W[i] - W[myCircularDequeGetFront(opt)] > maxWeight) {
//             myCircularDequeDeleteFront(opt);
//         }
        
//         f[i] = g[myCircularDequeGetFront(opt)] + neg[i] + 2;
//         if (i != n) {
//             g[i] = f[i] - neg[i + 1];
//             while (!myCircularDequeIsEmpty(opt) && g[i] <= g[myCircularDequeGetRear(opt)]) {
//                 myCircularDequeDeleteLast(opt);
//             }
//             myCircularDequeInsertLast(opt, i);
//         }
//     }
//     myCircularDequeFree(opt);
//     return f[n];
// }

int main(int argc, char const *argv[])
{
    /* code */
    return 0;
}

