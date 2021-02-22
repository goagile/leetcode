"""


Чередующиеся биты
=================

    >>> s = Solution()
    >>> s.hasAlternatingBits(5)
    True
    >>> s.hasAlternatingBits(7)
    False
    >>> s.hasAlternatingBits(11)
    False
    >>> s.hasAlternatingBits(10)
    True

    >>> s = FasterSolution()
    >>> s.hasAlternatingBits(5)
    True
    >>> s.hasAlternatingBits(7)
    False
    >>> s.hasAlternatingBits(11)
    False
    >>> s.hasAlternatingBits(10)
    True


"""


class Solution(object):

    def hasAlternatingBits(self, n):
        """
        :type n: int
        :rtype: bool
        """
        prev = None
        for x in bin(n)[2:]:
            if prev and prev == x:
                return False
            prev = x
        return True


class FasterSolution(object):
    def hasAlternatingBits(self, n):
        """
        :type n: int
        :rtype: bool
        """
        b = bin(n)[2:]
        for i in range(1, len(b)):
            if b[i] == b[i-1]:
                return False
        return True
