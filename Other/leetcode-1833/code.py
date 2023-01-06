from mergesort import merge

def maxIceCream(costs, coins) -> int:
    merge(costs)
    total = 0
    for c in costs:
        if coins == 0:
            return total
        coins -= c
        if coins >= 0:
            total += 1
    return total
