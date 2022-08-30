def abs(a):
    '''
    Function to get absolute value of number.
    
    Example:
    
    >>> abs(1)
    1
    >>> abs(-1)
    1
    >>> abs(0)
    0
    >>> abs("a")
    Traceback (most recent call last):
        ...
    TypeError: must be int or float!
    '''

    if not isinstance(a, (int, float)):
        raise TypeError("must be int or float!")
    if a >= 0:
        return a
    return -a


if __name__ == '__main__':
    import doctest

    doctest.testmod()

# 文档测试
# python3  -m doctest .\abs.py
