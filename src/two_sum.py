# Given an array of integers, return indices of the two numbers such that they add up to a specific target.
#
# You may assume that each input would have exactly one solution, and you may not use the same element twice.
#
# Example:
#
# Given nums = [2, 7, 11, 15], target = 9,
#
# Because nums[0] + nums[1] = 2 + 7 = 9,
# return [0, 1].

from typing import List


class Solution:

    def two_sum(self, nums: List[int], target: int) -> List[int]:
        idx = 0
        for val1 in nums:
            idx += 1
            val2 = target - val1
            if val2 in nums[idx:]:
                nums.pop(idx - 1)
                return [idx - 1, nums.index(val2) + 1]


solution = Solution()

assert solution.two_sum([2, 7, 11, 15], 9) == [0, 1]
assert solution.two_sum([3, 2, 4], 6) == [1, 2]
assert solution.two_sum([3, 3], 6) == [0, 1]
