"""https://leetcode.com/problems/two-sum/description/
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
"""
"""根据排列组合的知识，可以知道总共的组合可能是 (n*n -n)/2 种。挨个遍历的时间复杂度就是On2。所以我在leetcode上提交代码总是超时。
"""


# this code they say timeout


# import copy

# class Solution:
    # def twoSum(self, nums, target):
        # """
        # :type nums: List[int]
        # :type target: int
        # :rtype: List[int]
        # """
        # for i in range(len(nums)):
            # newnums = copy.copy(nums)
            # base = newnums.pop(i)
            # for j, item in enumerate(newnums):
                # if base + item == target:
                    # if j >= i:
                        # j += 1
                    # return [i, j]
                
 
# timeout again


# class Solution:
    # def twoSum(self, nums, target):
        # for i in range(len(nums)):
            # for j in range(i+1, len(nums)):
                # if nums[i] + nums[j] == target:
                    # return [i, j]
    

# 之前其实有遇到类似的问题，果然没有培养出思维意识，还是会忘
    
class Solution:

    def twoSum(self, nums, target):
        numdict = {}
        for index, item, in enumerate(nums):
            numdict[item] = index
        for i, n in enumerate(nums):
            if target - n in numdict:
                if numdict[target - n] != i:
                    return [i, numdict[target - n]]

        
 
                
if __name__ == "__main__":
    
    s = Solution()
    print(s.twoSum([1, 2, 3, 4, 5], 5))
                    
                
                    
                    
                    
                    
                    
                    
                    
                    
                    
                    
                    
                    
                    
                    
            
