def countPrimes(n):
    if n<1:
        # print(0)
        return 0

    primes = [True]*(n+1)
    p = 2

    while(p*p<=n):
        if primes[p]:
            for i in range(p**2, n+1, p):
                primes[i]=False

        p+=1

    primes[0]=False
    primes[1]=False

    # print(primes)


    count=0
    for i in primes:
        if i:
            count+=1

    if primes[-1]==True:
        count-=1

    # print(count)
    return count






countPrimes(10)