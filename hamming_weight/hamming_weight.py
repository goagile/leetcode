"""


Вес Хэмминга
============

Задача
------

Написать функцию, которая посчитает количество '1' в двоичной форме целого числа.

    >>> bin(12), bin(13)
    ('0b1100', '0b1101')

Решение с использованием спискового генератора

    >>> s = ListCompSolution()
    >>> s.hammingWeight(12)
    2
    >>> s.hammingWeight(13)
    3

Производительность

    In [16]: s
    Out[16]: <hamming_weight.ListCompSolution at 0x5be87f0>

    In [17]: %timeit s.hammingWeight(12)
    1.8 µs ± 4.83 ns per loop (mean ± std. dev. of 7 runs, 1000000 loops each)


Решение с встроенной функции count
Строковая функция str.count подсичтывает количество указанных символов в строке
Функция bin переводит целое число в строковый формат

    >>> s =BinSolution()
    >>> s.hammingWeight(12)
    2
    >>> s.hammingWeight(13)
    3

Производительность

    In [18]: s
    Out[18]: <hamming_weight.BinSolution at 0x5bbd588>

    In [19]: %timeit s.hammingWeight(12)
    565 ns ± 1.62 ns per loop (mean ± std. dev. of 7 runs, 1000000 loops each)


"""


class ListCompSolution(object):

    def hammingWeight(self, n):
        """
        :type n: int
        :rtype: int
        """
        return len([x for x in str(bin(n)) if x == '1'])


class BinSolution(object):

    def hammingWeight(self, n):
        """
        :type n: int
        :rtype: int
        """
        return bin(n).count('1')
