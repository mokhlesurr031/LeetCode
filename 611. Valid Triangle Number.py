def triangleNumber(nums):
    count = 0

    ln = len(nums)
    for i in range(ln):
        for j in range(i + 1, ln):
            for k in range(j + 1, ln):
                if nums[i] + nums[j] > nums[k] and nums[i] + nums[k] > nums[j] and nums[j] + nums[k] > nums[i]:
                    count += 1

    return count

x = [4,3,6,7]

r = triangleNumber(x)

print(r)