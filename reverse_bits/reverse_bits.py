"""


Реверс последовательности битов
===============================

    >>> bin(12)
    '0b1100'
    >>> int('0b1100', 2)
    12
    >>> int('0b0011', 2)
    3

    >>> s = Solution()
    >>> s.reverseBits(12)
    3

Производительность

    In [26]: s
    Out[26]: <reverse_bits.Solution at 0x5d36588>

    In [27]: %timeit s.reverseBits(12)
    2.45 µs ± 14.1 ns per loop (mean ± std. dev. of 7 runs, 100000 loops each)


    >>> s = Bit32Solution()
    >>> s.reverseBits(1)
    2147483648


"""


class Solution:

    def reverseBits(self, n):
        return int('0b' + ''.join(reversed(bin(n)[2:])), 2)


class Bit32Solution:

    def reverseBits(self, n):
        x = bin(n)[2:]
        y = len(x)
        z = x
        if y < 32:
            z = '0' * (32-y) + x
        return int('0b' + ''.join(reversed(z)), 2)
