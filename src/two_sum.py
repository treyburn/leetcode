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
