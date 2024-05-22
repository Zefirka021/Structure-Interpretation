from math import gcd

def make_rat(n, d):
    return (n, d)

def numer(x):
    g = gcd(x[0], x[1])
    return x[0] // g

def denom(x):
    g = gcd(x[0], x[1])
    return x[1] // g

def add_rat(x, y):
    return make_rat(
        (numer(x) * denom(y) + numer(y) * denom(x)),
        (denom(x) * denom(y))
    )

x = make_rat(1, 2)
y = make_rat(1, 3)
sum_result = add_rat(x, y)

print(f"Сумма чисел {numer(x)}/{denom(x)} и {numer(y)}/{denom(y)} равна {numer(sum_result)}/{denom(sum_result)}")
