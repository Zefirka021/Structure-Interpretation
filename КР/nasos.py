import math

def calculate_r(a):
    return a*15.7


def calculate_c(a):
    return a*12.45


def calculate_kr(a):
    return math.sqrt(1.5*a)


def calculate_kc(a):
    return math.log10(1.4*a)

def nasos(func_h, func_k):
    def calculate(h, k):
        return 0.29*func_h(h)*func_k(k)
    return calculate

calc_r_func = nasos(calculate_r, calculate_kr)
calc_c_func = nasos(calculate_c, calculate_kc)

h_value = 6.66
k_value = 7.77
print("Для насоса R: ", calc_r_func(h_value, k_value))
print("Для насоса C: ", calc_c_func(h_value, k_value))