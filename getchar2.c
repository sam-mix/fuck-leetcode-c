#include <stdio.h>
#include <unistd.h>

int getchar(void)
{
    static char buf[BUFSIZ];
    static char *bufp = buf;
    static int n = 0;

    if (n == 0)
    {
        n = read(0, buf, sizeof buf);
        bufp = buf;
    }
    return (--n >= 0) ? (unsigned char)*bufp++ : EOF;
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
