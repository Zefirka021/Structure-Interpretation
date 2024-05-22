# zdes byla vasya

def f1(a, b):
    return a + b

def f2(op, prekol1, prekol2):
    if op == "+":
        return prekol1 + prekol2
    elif op == "-":
        return prekol1 - prekol2
    elif op == "*":
        return prekol1 * prekol2
    elif op == "/":
        return prekol1 / prekol2
    else:
        raise Exception("Invalid operator")
    
def pofig(a, b, op, shag, f1, f2):
    result = 0
    for i in range(shag):
        if i == 0:
            result = 1 / f1(a, b)
            a += 4
            b += 4
            print(result)
        else:
            result = f2(op, result, 1 / f1(a, b))
            a += 4
            b += 4
            print(result)
    return result
    
print(pofig(1, 3, '/', 5, f1, f2))