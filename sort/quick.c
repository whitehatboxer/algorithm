#include <stdio.h>

void quick_sort(int* array, int left, int right)
{
    if (left >= right)
        return;

    int i = left;
    int j = right;
    int key = array[left];

    while (i < j) {
        while (i < j && array[j] >= key) {
            j--;
        }
        // 交换位置，将 j 移到左边
        array[i] = array[j];

        while (i < j && array[i] <= key) {
            i++;
        }
        // 将 i 移到右边，此时 j 的位置已经空缺了
        array[j] = array[i];
    }

    // 将 key 还原
    array[i] = key;
    quick_sort(array, left, i - 1);
    quick_sort(array, i + 1, right);
}


int main()
{
    int array[] = {1, 32, 2, -1, 0, 12, 99, 23};
    int length = sizeof(array) / 4;

    quick_sort(array, 0, length - 1);
    for (int i = 0; i < length; i++) {
        printf("%d ", array[i]);
    }
    printf("\n");
}


