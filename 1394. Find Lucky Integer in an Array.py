arr = [1,2,2,2,3,2,3]


def check_lucky_number(nums):
    count = {}

    for num in nums:
        count[num] = count.get(num, 0)+1

    res = 0

    for key, val in count.items():
        if key==val:
            if key > res:
                res=key


    if res==0:
        return -1
    else:
        return res

print(check_lucky_number(arr))