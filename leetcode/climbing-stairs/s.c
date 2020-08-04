#include <stdio.h>

int climb_stairs(int n, int* mem){
    if (n <= 2)
        return n;

    if (!mem[n - 1])
        mem[n-1] = climb_stairs(n-1, mem);
    if (!mem[n - 2])
        mem[n-2] = climb_stairs(n-2, mem);
    return mem[n - 1] + mem[n - 2];
}


int main()
{
    int mem[100];
    printf("%d\n", sizeof(mem));
    memset(mem, 0, sizeof(mem));
    long long c = climb_stairs(45, mem);
    printf("have %d solutions.\n", c);
}

