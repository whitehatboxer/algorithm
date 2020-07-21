#include <stdio.h>

int get_area(int a, int b, int length)
{
    int height;
    if (a > b)
        height = b;
    else
        height = a;
    printf("left is: %d, right is: %d\n", a, b);
    printf("length is: %d\n", height);
    return length * height;
}

int max_area(int* height, int height_size)
{
    // 分别从头部和尾部开始遍历
    int left = 0;
    int right = height_size - 1;

    // 存放部分中间结果，这种情况下的中间结果并不多
    int max = get_area(height[left], height[right], right - left);
    
    int scan_left = 1;
    while (left < right) {
        if (left == right - 1) {
            int value = get_area(height[left], height[right], right - left);
            if (value > max) {
                max = value;
            }
            break;
        }

        if (scan_left) {
            if (height[left+1] > height[left]) {
                int value = get_area(height[left+1], height[right], right - left - 1);
                if (value > max)
                    max = value;
                scan_left = 0; // 切换遍历方向
            }
            left ++;
        } else {
            if (height[right-1] > height[right]) {
                int value = get_area(height[left], height[right-1], right - 1 - left);
                if (value > max)
                    max = value;
                scan_left = 1; // 切换遍历方向
            }
            right --;
        }
    }
    return max;
}

int main()
{
    // int array[] = {1,8,6,2,5,4,8,3,7};
    int array[] = {2,3,10,5,7,8,9};
    int max = max_area(array, sizeof(array)/4);
    printf("max is: %d\n", max);
}

