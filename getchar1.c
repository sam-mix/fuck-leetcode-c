#include <unistd.h>
#include <stdio.h>

int getchar(void)
{
    char c;
    return (read(0, &c, 1) == 1) ? (unsigned char)c : EOF;
}

// int main(int argc, char const *argv[])
// {
//     int x = getchar();
//     printf("%d\n", x);
//     return 0;
// }

int main(int argc, char const *argv[])
{
    /* code */
    return 0;
}
