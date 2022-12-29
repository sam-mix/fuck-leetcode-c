#include <stdio.h>

void *my_memcpy(void *dest, void *src, int byte_count)
{
    char *s = src;
    char *d = dest;

    while (byte_count--)
    {
        *d++ = *s++;
    }

    return dest;
}

int main()
{
    printf("hello world!\n");
    char *a = "hello world!";
    char b[13];
    my_memcpy(b, a, 12);
    printf("%s", b);
    return 0;
}