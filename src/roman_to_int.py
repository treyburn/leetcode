class Solution:
    value_dict = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
    sub_dict = {"V": "I", "X": "I", "L": "X", "C": "X", "D": "C", "M": "C"}

    def roman_to_int(self, s: str) -> int:
        """
        Takes in roman numerals and converts to integer. Utilizes value_dict and sub_dict to determine value of numerals
        :param s: roman numeral string
        :return: integer value of the roman numerals
        """
        _solution = 0
        for idx in range(len(s)):
            if idx + 1 < len(s):
                if s[idx + 1] in self.sub_dict and s[idx] == self.sub_dict[s[idx + 1]]:
                    _solution -= self.value_dict[s[idx]]
                else:
                    _solution += self.value_dict[s[idx]]
            else:
                _solution += self.value_dict[s[idx]]
        return _solution


solution = Solution()

assert solution.roman_to_int("III") == 3
assert solution.roman_to_int("IV") == 4
assert solution.roman_to_int("IX") == 9
assert solution.roman_to_int("LVIII") == 58
assert solution.roman_to_int("MCMXCIV") == 1994
