#include <stdio.h>

int max_area(int* height, int height_size){
	
	int max = 0;
    for (int len = 1; len < height_size; len ++) {
		for (int i = 0; i < height_size; i ++) {
			int j = i + len;
			if (j >= height_size) {
				break;
			}
			
			int value;
			if (height[j] > height[i])
				value = height[i] * len;
			else
				value = height[j] * len;
			if (value > max) {
				max = value;
			}
		}
	}
	
	return max;
}


int main()
{
    int array[] = {1,8,6,2,5,4,8,3,7};
    int max = max_area(array, sizeof(array)/4);
    printf("max is: %d\n", max);
}
