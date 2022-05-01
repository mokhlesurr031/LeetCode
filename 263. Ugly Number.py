import math
def isUgly(n):
    if n <= 0:
        return False

    p = 2

    factors = []
    primes = []
    j=2
    f=1

    while(p<=n/2):
        if n%p==0:
            # factors.append(p)
            num = p
            while (j < int(math.sqrt(num))):
                if num % j == 0:
                    f = 0
                    break
                j += 1
            if f == 1:
                primes.append(num)

        p+=1


    n_fact = 0
    for i in range(2, int(math.sqrt(n))):
        if n%i==0:
            n_fact+=1

    if n_fact==0:
        primes.append(n)

    print(primes)


    # factors.append(n)

    # print(factors)
    # primes = []
    # for i in factors:
    #     num = i
    #     j = 2
    #     f = 1
    #     while(j<int(math.sqrt(num))):
    #         if num%j==0:
    #             f = 0
    #             break
    #         j+=1
    #     if f==1:
    #         primes.append(num)

    # print(primes)

    # if primes[-1]>5:
    #     return False
    # return True



print(isUgly(214797179))