from typing import List

class Solution:
    
    def two_sum(self, nums: List[int], except_index: int, target: int) -> List[List[int]]:
        need_dict = {}
        res = []
        for i in range(len(nums)):
            if i == except_index:
                continue
            n = nums[i]
            need_dict[target - n] = i
        for i in range(len(nums)):
            if i == except_index:
                continue
            n = nums[i]
            value_index = need_dict.get(n)
            if value_index is not None and value_index != i:
                res.append([n, nums[value_index]])
        return res

    def three_sum(self, nums: List[int], target: int) -> List[List[int]]:
        need_dict = {}
        res = []
        for i in range(len(nums)):
            need_dict[i] = target - nums[i]
        for k, v in need_dict.items():
            supple_res = self.two_sum(nums, k, v)
            for sr in supple_res:
                res.append(sr + [nums[k]])
        res = self.remove_duplication(res)
        return res

    def remove_duplication(self, dup_list: List[List[int]]):
        res = []
        the_filter = set()
        for d in dup_list:
            d.sort()
            hashed_value = str(d)
            if hashed_value in the_filter:
                continue
            the_filter.add(hashed_value)
            res.append(d)
        return res
            



sl = Solution()
res = sl.three_sum([-1, 0, 1, 2, -1, -4], 0)


print(res)



