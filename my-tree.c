#include <stdio.h>
#include <ctype.h>
#include <string.h>

#define MAXWORD 10

struct tnode
{
    char *word;
    int count;
    struct tnode *left;
    struct tnode *right;
};

struct tnode *addtree(struct tnode *, char *);
void treeprint(struct tnode *);
int getword(char *, int);

// /* 单词出现频率的统计 */
// int main()
// {
//     struct tnode *root;
//     char word[MAXWORD];
//     root = NULL;

//     while (getword(word, MAXWORD) != EOF)
//     {
//         if (word[0] == 'q')
//         {
//             break;
//         }
//         if (isalpha(word[0]))
//         {
//             root = addtree(root, word);
//         }
//     }
//     treeprint(root);

//     return 0;
// }

int main()
{
}

struct tnode *tallc(void);
char *mystrdup(char *);

/* addtree 函数：在 p 的位置或p 的下方添加一个 w 节点 */
struct tnode *addtree(struct tnode *p, char *w)
{
    int cond;
    if (p == NULL)
    {                // 该单词是个新的单词
        p = tallc(); // 创建一个新的节点
        p->word = mystrdup(w);
        p->count = 1;
        p->left = p->right = NULL;
    }
    else if ((cond = strcmp(p->word, w)) == 0)
    {
        p->count++;
    }
    else if (cond < 0)
    {
        p->left = addtree(p->left, w);
    }
    else
    {
        p->right = addtree(p->right, w);
    }
    return p;
}

void treeprint(struct tnode *p)
{
    if (p != NULL)
    {
        treeprint(p->left);
        printf("%4d, %s\n", p->count, p->word);
    }
}

#include <stdlib.h>

struct tnode *tallc(void)
{
    return (struct tnode *)malloc(sizeof(struct tnode));
}

char *mystrdup(char *s)
{
    char *p;
    p = (char *)malloc(strlen(s) + 1);
    if (p != NULL)
    {
        strcpy(p, s);
    }
    return p;
}

int getword(char *word, int lim)
{
    int c, getch(void);
    void ungetch(int);

    char *w = word;
    while (isspace(c = getch()))
        ;
    if (c != EOF)
    {
        *w++ = c;
    }
    if (!isalpha(c))
    {
        *w = '\0';
        return c;
    }
    for (; --lim > 0; w++)
    {
        if (!isalnum(*w = getch()))
        {
            ungetch(*w);
            break;
        }
    }
    *w = '\0';
    return word[0];
}

#define BUFSIZE 100
char buf[BUFSIZE];

int bufp = 0;

int getch(void)
{
    return (bufp > 0) ? buf[bufp--] : getchar();
}

void ungetch(int c)
{
    if (bufp > BUFSIZE)
    {
        printf("ungetch: too many characters\n");
    }
    else
    {
        buf[bufp++] = c;
    }
}