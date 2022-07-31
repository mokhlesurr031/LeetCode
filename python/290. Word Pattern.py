pattern = "abba"
s = "dog cat cat fish"



lst = list(s.split(' '))
d = {}

count = 0
for i in range(len(pattern)):

    if pattern[i] not in d:
        d[pattern[i]] = lst[i]
    else:
        # print('yess')
        if d[pattern[i]] == lst[i]:
            continue
        else:
            count+=1
            break
#
# if count==0:
#     # print(True)
#     return True
# else:
#     # print(False)
#     return False
# print(d)

