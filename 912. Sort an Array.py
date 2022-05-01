class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        if len(nums) <= 1:
            return nums

        mid = len(nums) // 2
        left = nums[:mid]
        right = nums[mid:]

        left = self.sortArray(left)
        right = self.sortArray(right)

        return self.mergeTwoLists(left, right)

    def mergeTwoLists(self, a, b):
        sorted_list = []

        len_a = len(a)
        len_b = len(b)
        i = j = 0

        while i < len_a and j < len_b:
            if a[i] <= b[j]:
                sorted_list.append(a[i])
                i += 1
            else:
                sorted_list.append(b[j])
                j += 1

        while (i < len_a):
            sorted_list.append(a[i])
            i += 1
        while (j < len_b):
            sorted_list.append(b[j])
            j += 1

        return sorted_list
