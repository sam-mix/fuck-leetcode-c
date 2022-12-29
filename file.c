#include <stdio.h>
#include <stdlib.h>

/* filecopy 函数：将文件ifp复制到文件ofp */
void filecopy(FILE *, FILE *);

int main(int argc, char const *argv[])
{
    FILE *fp;
    const char *prog = argv[0];
    if (argc == 1)
    {
        // filecopy(stdin, stdout);
        printf("至少输入一个文件名！！！！");
    }
    else
    {
        while (--argc > 0)
        {
            if ((fp = fopen(*++argv, "r")) == NULL)
            {
                fprintf(stderr, "%s: cat not open %s\n", prog, *argv);
                exit(1);
            }
            else
            {
                filecopy(fp, stdout);
                fclose(fp);
            }
        }
    }
    if (ferror(stdout))
    {
        fprintf(stderr, "%s: error writting stdout\n", prog);
        exit(2);
    }
    return 0;
}

void filecopy(FILE *ifp, FILE *ofp)
{
    int c;
    while ((c = getc(ifp)) != EOF)
    {
        putc(c, ofp);
    }
    putc('\n', ofp);
}