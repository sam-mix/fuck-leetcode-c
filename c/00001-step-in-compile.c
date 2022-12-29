/*
# 预处理:
gcc -E 00001-step-in-compile.c -o 00001-step-in-compile.i
# 编译:
gcc -S 00001-step-in-compile.i -o 00001-step-in-compile.s
# 汇编:
gcc -c 00001-step-in-compile.s -o 00001-step-in-compile.o
# 链接:
gcc 00001-step-in-compile.o -o 00001-step-in-compile.exe

*/

int main(int argc, char const *argv[])
{
    return 0;
}
