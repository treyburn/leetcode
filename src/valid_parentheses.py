class Solution:
    _open = ['(', '{', '[']
    _close = [')', '}', ']']

    def is_valid(self, s: str) -> bool:
        _list = []
        _len = len(s)
        if not _len:
            return True
        if _len % 2:
            return False
        for _idx in range(_len):
            if s[_idx] in self._open:
                _list.append(s[_idx])
            elif not len(_list) and s[_idx] in self._close:
                return False
            elif self._open[self._close.index(s[_idx])] == _list[len(_list) - 1]:
                _list.pop()
            else:
                return False
        if not len(_list):
            return True






solution = Solution()

assert solution.is_valid("()")
assert solution.is_valid("()[]{}")
assert not solution.is_valid("(]")
assert not solution.is_valid("([)]")
assert solution.is_valid("{[]}")
assert not solution.is_valid("(")
assert not solution.is_valid("((")
assert not solution.is_valid("){")
