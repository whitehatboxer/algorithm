#include <stdio.h>
#include <string.h>

int climb_stairs(int n){
    if (n <= 2)
        return n;

    int mem[100];
    memset(mem, 0, sizeof(mem));
    mem[0] = 0;
    mem[1] = 1;
    mem[2] = 2;
    for (int i = 3; i <= n; i++) {
        mem[i] = mem[i - 1] + mem[i - 2];
    }

    return mem[n];
}


int main()
{
    long c = climb_stairs(45);
    printf("have %d solutions.\n", c);
}

