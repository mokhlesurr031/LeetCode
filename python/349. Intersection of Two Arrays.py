def intersection(a, b, m, n):
    i = 0
    j = 0

    a = sorted(a)
    b = sorted(b)

    # print(a)
    # print(b)
    l = []

    while (i<m and j<n):
        if a[i]<b[j]:
            i+=1
        elif (a[i]>b[j]):
            j+=1
        elif(a[i]==b[j]):
            # print(a[i])
            l.append(a[i])
            i+=1
            j+=1

    print(l)

a = [1,2,2,1]
b = [2,2]
#
# a = [4,9,5]
# b = [9,4,9,8,4]

m = len(a)
n = len(b)

intersection(a, b, m, n)