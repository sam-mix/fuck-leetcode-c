#include <stdio.h>

#define MY_HELLO_WORLD printf(MY_FORMAT, MY_HELLO)
#define MY_HELLO "Hello"
#define MY_FORMAT "%s, world!"

int main() {
    MY_HELLO_WORLD;
    return 0;
}
