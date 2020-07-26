from typing import List

class Solution:
    
    def three_sum(self, nums: List[int], target: int) -> List[List[int]]:
        left = 0
        right = 0
        nums.sort()
        for i in range(len(nums)):
            while 
            need_dict[i] = target - nums[i]
        for k, v in need_dict.items():
            supple_res = self.two_sum(nums, k, v)
            for sr in supple_res:
                res.append(sr + [nums[k]])
        res = self.remove_duplication(res)
        return res



sl = Solution()
res = sl.three_sum([-1, 0, 1, 2, -1, -4], 0)


print(res)



