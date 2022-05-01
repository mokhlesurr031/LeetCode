def findTheDifference(s, t):
    count = {}

    for i in t:
        count[i] = count.get(i, 0)+1

    # print(count)

    for j in s:
        count[j]-=1

    # print(count)

    for res in count:
        if count[res]>0:
            return res



findTheDifference('abcd', 'abcdde')





