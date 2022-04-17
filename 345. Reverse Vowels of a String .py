s = "Aa"

ln = len(s)
vowel = ['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U']
lst = []


for i in s:
    if i in vowel:
        lst.append(i)

st = ''
rlst = lst[::-1]
# print(rlst)

count = 0
for j in range(ln):
    if s[j] in vowel:
        # print(count)
        st+=rlst[count]
        count+=1
        # break

    else:
        st+=s[j]
        # count+=1
        # break

# print(count)
print(st)
# return st