#include <stdio.h>

int find_length(int* a, int a_size, int* b, int b_size) {
    
    int max = 0;
    int dp[b_size][a_size];
    
    for (int i = 0; i < a_size; i++) {
        for (int j = 0; j < b_size; j++) {
            dp[j][i] = 0;
        }
    }

    for (int i = 0; i < a_size; i++) {
        dp[0][i] = a[i] == b[0] ? 1 : 0;
        max = max > dp[0][i] ? max : dp[0][i];
    }
    
    for (int j = 0; j < b_size; j++) {
        dp[j][0] = a[0] == b[j] ? 1 : 0;
        max = max > dp[j][0] ? max : dp[j][0];
    }

    if (a_size == 1 || b_size == 1)
        return max;

    for (int i = 1; i < a_size; i++) {
        for (int j = 1; j < b_size; j++) {
            dp[j][i] = a[i] == b[j] ? dp[j-1][i-1] + 1 : 0;
            max = max > dp[j][i] ? max : dp[j][i];
        }
    }

    return max;            
}


int main()
{
    int a[] = {1,2,3,2,1};
    int b[] = {3,2,1,4,7};
    int res = find_length(a, 5, b, 5);
    printf("res is: %d\n", res);
}
