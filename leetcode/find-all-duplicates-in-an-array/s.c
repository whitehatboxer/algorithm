#include <stdio.h>
#include <stdlib.h>
#include <string.h>
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *findDuplicates(int *nums, int numsSize, int *returnSize)
{
    int *res = malloc(numsSize * sizeof(int));
    int *start = res;
    memset(res, 0, numsSize * sizeof(int));
    int i;
    int v;
    *returnSize = 0;

    for (i = 0; i < numsSize; i++)
    {
        v = nums[i];
        if (v < 0)
        {
            v = -v;
        }

        if (nums[v - 1] < 0)
        {
            *res++ = v;
            (*returnSize)++;
        }
        else
        {
            nums[v - 1] = -nums[v - 1];
        }
    }

    return start;
}

int main()
{
    int nums[] = {4, 3, 2, 7, 8, 2, 3, 1};
    int returnSize;
    int *res = findDuplicates(nums, sizeof(nums) / sizeof(int), &returnSize);
    int i;
    printf("return size: %d\n", returnSize);
    for (i = 0; i < returnSize; i++)
    {
        printf("%d ", res[i]);
    }
    printf("\n");
}
