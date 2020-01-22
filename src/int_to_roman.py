class Solution:
    value_dict = {1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C", 400: "CD", 500: "D",
                  900: "CM", 1000: "M"}

    def int_to_roman(self, num: int) -> str:
        pass


solution = Solution()

assert solution.int_to_roman(3) == "III"
assert solution.int_to_roman(4) == "IV"
assert solution.int_to_roman(9) == "IX"
assert solution.int_to_roman(58) == "LVIII"
assert solution.int_to_roman(1994) == "MCMXCIV"
