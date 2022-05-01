count = {}

words = ['foo', 'bar', 'the', 'the']

for word in words:
    count[word] = count.get(word, 0)+1

print(count)