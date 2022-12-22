#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define HASHSIZE 101
static struct nlist *hashtab[HASHSIZE];

struct nlist
{
    struct nlist *next;
    char *name;
    char *defn;
};

/* hash函数: 位字符串s生成散列值*/
unsigned hash(char *s)
{
    unsigned hashval;
    for (hashval = 0; *s != '\0'; s++)
    {
        hashval = *s + 31 * hashval;
    }
    return hashval % HASHSIZE;
}

/* lookup 函数: 在hashtab中查找s */
struct nlist *lookup(char *s)
{
    struct nlist *np;
    for (np = hashtab[hash(s)]; np != NULL; np = np->next)
    {
        if (strcmp(s, np->name) == 0)
        {
            return np;
        }
    }
    return NULL;
}

struct nlist *install(char *name, char *defn)
{
    struct nlist *np;
    unsigned hashval;
    if ((np = lookup(name)) == NULL)
    {
        np = (struct nlist *)malloc(sizeof(*np));
        if (np == NULL || (np->name = strdup(name)) == NULL)
        {
            return NULL;
        }
        hashval = hash(name);
        np->next = hashtab[hashval];
        hashtab[hashval] = np;
    }
    else
    {
        free((void *)np->defn);
    }
    if ((np->defn = strdup(defn)) == NULL)
    {
        return NULL;
    }
    return np;
}

typedef struct tnode *Treeptr;

typedef struct tnode
{
    char *word;
    int count;
    Treeptr left;
    Treeptr right;
} Treenode;

Treeptr talloc(void)
{
    return (Treeptr)malloc(sizeof(Treenode));
}

int main()
{
    return 0;
}