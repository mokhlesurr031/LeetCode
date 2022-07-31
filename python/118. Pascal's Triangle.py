def generate(numRows):
    res = []
    for line in range(1, numRows + 1):
        C = 1
        l = []
        for i in range(1, line + 1):
            l.append(C)
            C = int(C * (line - i) / i)

        # print(l)
        res.append(l)
        l=[]
    # print(res)
    return res
generate(5)