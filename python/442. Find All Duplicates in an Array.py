def findDuplicates(nums):
    nums = sorted(nums)

    twice_list = []
    # print(nums)
    for i in range(1, len(nums)):
        if nums[i - 1] == nums[i]:
            twice_list.append(nums[i])

    print(twice_list)

    return twice_list

findDuplicates([4,3,2,7,8,2,3,1])