from typing import List


class Solution:

    def two_sum(self, nums: List[int], target: int) -> List[int]:
        """
        two_sum takes in a list in integers and determines if two of the values in that list can sum up to the target.
        If that condition is met then it returns the index location of those values in the list.
        :param nums: a list of integers
        :param target: target integer
        :return: returns a list of the two indexes - or none
        """
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
