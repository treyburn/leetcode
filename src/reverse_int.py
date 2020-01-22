class Solution:
    def reverse(self, x: int) -> int:
        """
        reverse takes an integer and reverses its value. ie 123 becomes 321, -123 becomes -321, 120 becomes 21.
        If the value of the reversed string is greater than 2**31, we overflow and produce a 0 as the result.
        :param x: an integer
        :return: solution integer or 0 if this would cause int overflow
        """
        str_x = str(x)[::-1]
        if '-' in str_x:
            str_x = '-' + str_x.strip('-')
        if abs(int(str_x)) - 1 >= 2 ** 31:
            return 0
        else:
            return int(str_x)


solution = Solution()

assert solution.reverse(123) == 321
assert solution.reverse(-123) == -321
assert solution.reverse(120) == 21
