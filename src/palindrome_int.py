class Solution:
    def is_palindrome(self, x: int) -> bool:
        """
        Converts int to string and does a string reversal comparison - if strings match then its a palindrome
        :param x: the integer to check if meets palindrome condition
        :return: bool value of whether int meets palindrome condition
        """
        if str(x) == str(x)[::-1]:
            return True
        return False


solution = Solution()

assert solution.is_palindrome(121) is True
assert not solution.is_palindrome(-121)
assert not solution.is_palindrome(10)
