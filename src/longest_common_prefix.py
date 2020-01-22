from typing import List


class Solution:
    def longest_common_prefix(self, strs: List[str]) -> str:
        prefix = ""
        if strs:
            if None in strs:
                return prefix
            shortest = min(strs, key=len)
        else:
            return prefix
        for idx in range(len(shortest)):
            for word in strs:
                if not shortest[idx] == word[idx]:
                    return prefix
            prefix += shortest[idx]
        return prefix


solution = Solution()

assert solution.longest_common_prefix(["flower", "flow", "flight"]) == "fl"
assert solution.longest_common_prefix(["dog", "racecar", "car"]) == ""
assert solution.longest_common_prefix(["dog", "racecar", None]) == ""
assert solution.longest_common_prefix([]) == ""
