#include <stdarg.h>
#include <stdio.h>

/* minprintf函数，带有可变参数表的printf函数 */
void minprintf(char *fmt, ...)
{
    va_list ap; /* 依次指向每个无名参数 */
    char *p, *sval;
    int ival;
    double fval;

    va_start(ap, fmt);
    for (p = fmt; *p; p++)
    {
        if (*p != '%')
        {
            putchar(*p);
            continue;
        }
        switch (*++p)
        {
        case 'd':
            ival = va_arg(ap, int);
            printf("%d", ival);
            break;
        case 'f':
            fval = va_arg(ap, double);
            printf("%f", fval);
            break;
        case 's':
            for (sval = va_arg(ap, char *); *sval; sval++)
            {
                putchar(*sval);
            }
            break;
        default:
            putchar(*p);
            break;
        }
    }
}

int main(int argc, char const *argv[])
{
    minprintf("%s\n", "Hello World!");
    minprintf("%f\n", 3.1415926f);
    minprintf("%d\n", 10);
    return 0;
}
