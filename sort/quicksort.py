# coding: utf-8


def quicksort(array, left, right):

    key_index = left

    if left < right:

        l = left + 1
        r = right
        
        while l < r:
            while array[r] >= array[key_index] and r > l:
                r -= 1
            while array[l] <= array[key_index] and l < r:
                l += 1

            array[l], array[r] = array[r], array[l]
        
        quicksort(array, left, key_index-1)
        quicksort(array, key_index, right)


array = [2, 9, 3, 4, 6, 11, 5, 2, 12]
import pdb; pdb.set_trace()
quicksort(array, 0, 8)

print array